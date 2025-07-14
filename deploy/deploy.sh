#!/bin/bash

# 医院管理系统部署脚本

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# 检查Docker是否安装
check_docker() {
    if ! command -v docker &> /dev/null; then
        log_error "Docker未安装，请先安装Docker"
        log_info "安装命令:"
        log_info "  Ubuntu/Debian: sudo apt-get install docker.io"
        log_info "  CentOS/RHEL: sudo yum install docker"
        log_info "  macOS: brew install docker"
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        log_error "Docker Compose未安装，请先安装Docker Compose"
        log_info "安装命令:"
        log_info "  sudo curl -L \"https://github.com/docker/compose/releases/download/v2.20.0/docker-compose-$(uname -s)-$(uname -m)\" -o /usr/local/bin/docker-compose"
        log_info "  sudo chmod +x /usr/local/bin/docker-compose"
        exit 1
    fi
    
    log_success "Docker环境检查通过"
}

# 检查端口是否被占用
check_ports() {
    local ports=("80" "443" "2379" "2380" "3306" "6379" "8080" "8081" "8888")
    
    for port in "${ports[@]}"; do
        if netstat -tuln | grep -q ":$port "; then
            log_warning "端口 $port 已被占用，请确保没有其他服务在使用"
        fi
    done
}

# 创建必要的目录
create_directories() {
    log_info "创建必要的目录..."
    
    mkdir -p deploy/nginx/ssl
    mkdir -p deploy/logs
    mkdir -p deploy/data/mysql
    mkdir -p deploy/data/redis
    mkdir -p deploy/data/etcd
    
    log_success "目录创建完成"
}

# 生成SSL证书（自签名）
generate_ssl_cert() {
    log_info "生成SSL证书..."
    
    if [ ! -f "deploy/nginx/ssl/server.crt" ]; then
        openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
            -keyout deploy/nginx/ssl/server.key \
            -out deploy/nginx/ssl/server.crt \
            -subj "/C=CN/ST=Beijing/L=Beijing/O=Hospital/OU=IT/CN=localhost"
        log_success "SSL证书生成完成"
    else
        log_info "SSL证书已存在，跳过生成"
    fi
}

# 构建镜像
build_images() {
    log_info "构建Docker镜像..."
    
    # 构建API网关镜像
    log_info "构建API网关镜像..."
    docker build -f deploy/Dockerfile.api -t hospital-api-gateway:latest .
    
    # 构建医生服务镜像
    log_info "构建医生服务镜像..."
    docker build -f deploy/Dockerfile.doctor -t hospital-doctor-service:latest .
    
    # 构建患者服务镜像
    log_info "构建患者服务镜像..."
    docker build -f deploy/Dockerfile.patient -t hospital-patient-service:latest .
    
    log_success "镜像构建完成"
}

# 启动服务
start_services() {
    log_info "启动服务..."
    
    cd deploy
    
    # 启动基础服务
    log_info "启动基础服务 (etcd, MySQL, Redis)..."
    docker-compose up -d etcd mysql redis
    
    # 等待基础服务启动
    log_info "等待基础服务启动..."
    sleep 30
    
    # 启动微服务
    log_info "启动微服务..."
    docker-compose up -d doctor-service patient-service
    
    # 等待微服务启动
    log_info "等待微服务启动..."
    sleep 20
    
    # 启动API网关
    log_info "启动API网关..."
    docker-compose up -d api-gateway
    
    # 等待API网关启动
    log_info "等待API网关启动..."
    sleep 10
    
    # 启动Nginx
    log_info "启动Nginx..."
    docker-compose up -d nginx
    
    cd ..
    
    log_success "所有服务启动完成"
}

# 检查服务状态
check_services() {
    log_info "检查服务状态..."
    
    cd deploy
    
    # 检查容器状态
    docker-compose ps
    
    # 检查服务健康状态
    log_info "检查API服务..."
    if curl -s http://localhost:8888/departments >/dev/null 2>&1; then
        log_success "API服务运行正常"
    else
        log_error "API服务未正常运行"
    fi
    
    log_info "检查Nginx服务..."
    if curl -s http://localhost/health >/dev/null 2>&1; then
        log_success "Nginx服务运行正常"
    else
        log_error "Nginx服务未正常运行"
    fi
    
    cd ..
}

# 停止服务
stop_services() {
    log_info "停止服务..."
    
    cd deploy
    docker-compose down
    cd ..
    
    log_success "服务已停止"
}

# 清理资源
cleanup() {
    log_info "清理资源..."
    
    cd deploy
    docker-compose down -v --remove-orphans
    cd ..
    
    # 删除镜像
    docker rmi hospital-api-gateway:latest hospital-doctor-service:latest hospital-patient-service:latest 2>/dev/null || true
    
    log_success "资源清理完成"
}

# 查看日志
view_logs() {
    local service=$1
    
    if [ -z "$service" ]; then
        log_info "查看所有服务日志..."
        cd deploy
        docker-compose logs -f
        cd ..
    else
        log_info "查看 $service 服务日志..."
        cd deploy
        docker-compose logs -f $service
        cd ..
    fi
}

# 重启服务
restart_services() {
    log_info "重启服务..."
    
    cd deploy
    docker-compose restart
    cd ..
    
    log_success "服务重启完成"
}

# 显示帮助信息
show_help() {
    echo "医院管理系统部署脚本"
    echo ""
    echo "用法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  deploy      部署所有服务"
    echo "  start       启动服务"
    echo "  stop        停止服务"
    echo "  restart     重启服务"
    echo "  status      查看服务状态"
    echo "  logs [服务] 查看日志"
    echo "  cleanup     清理资源"
    echo "  help        显示帮助信息"
    echo ""
    echo "示例:"
    echo "  $0 deploy    # 完整部署"
    echo "  $0 logs      # 查看所有日志"
    echo "  $0 logs api-gateway  # 查看API网关日志"
}

# 主函数
main() {
    case "$1" in
        "deploy")
            log_info "开始部署医院管理系统..."
            check_docker
            check_ports
            create_directories
            generate_ssl_cert
            build_images
            start_services
            check_services
            log_success "部署完成！"
            log_info "访问地址: http://localhost"
            log_info "API地址: http://localhost:8888"
            ;;
        "start")
            start_services
            check_services
            ;;
        "stop")
            stop_services
            ;;
        "restart")
            restart_services
            ;;
        "status")
            cd deploy
            docker-compose ps
            cd ..
            ;;
        "logs")
            view_logs $2
            ;;
        "cleanup")
            cleanup
            ;;
        "help"|"-h"|"--help")
            show_help
            ;;
        *)
            log_error "未知选项: $1"
            show_help
            exit 1
            ;;
    esac
}

# 脚本入口
if [ "$0" = "$BASH_SOURCE" ]; then
    main "$@"
fi 