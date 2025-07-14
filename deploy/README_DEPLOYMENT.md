# 医院管理系统部署指南

## 概述

本文档提供了医院管理系统基于go-zero框架的完整部署方案，包括Docker容器化部署、服务配置、监控和维护等内容。

## 系统架构

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Nginx (80)    │    │   API Gateway   │    │   Microservices │
│   (反向代理)     │───▶│   (8888)        │───▶│   (8080/8081)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │                        │
                                ▼                        ▼
                       ┌─────────────────┐    ┌─────────────────┐
                       │   etcd (2379)   │    │   MySQL (3306)  │
                       │   (服务发现)     │    │   (数据存储)     │
                       └─────────────────┘    └─────────────────┘
                                                       │
                                                       ▼
                                              ┌─────────────────┐
                                              │   Redis (6379)  │
                                              │   (缓存)        │
                                              └─────────────────┘
```

## 部署方式

### 方式1: Docker Compose 部署（推荐）

#### 前置要求

1. **Docker环境**
   ```bash
   # Ubuntu/Debian
   sudo apt-get update
   sudo apt-get install docker.io docker-compose
   
   # CentOS/RHEL
   sudo yum install docker docker-compose
   
   # macOS
   brew install docker docker-compose
   ```

2. **端口检查**
   - 80: Nginx HTTP
   - 443: Nginx HTTPS
   - 2379: etcd
   - 2380: etcd peer
   - 3306: MySQL
   - 6379: Redis
   - 8080: 医生服务
   - 8081: 患者服务
   - 8888: API网关

#### 快速部署

```bash
# 1. 克隆项目
git clone <repository-url>
cd hospital

# 2. 给部署脚本执行权限
chmod +x deploy/deploy.sh

# 3. 执行部署
./deploy/deploy.sh deploy
```

#### 部署脚本使用

```bash
# 完整部署
./deploy/deploy.sh deploy

# 启动服务
./deploy/deploy.sh start

# 停止服务
./deploy/deploy.sh stop

# 重启服务
./deploy/deploy.sh restart

# 查看状态
./deploy/deploy.sh status

# 查看日志
./deploy/deploy.sh logs
./deploy/deploy.sh logs api-gateway

# 清理资源
./deploy/deploy.sh cleanup

# 查看帮助
./deploy/deploy.sh help
```

### 方式2: 手动部署

#### 1. 构建镜像

```bash
# 构建API网关
docker build -f deploy/Dockerfile.api -t hospital-api-gateway:latest .

# 构建医生服务
docker build -f deploy/Dockerfile.doctor -t hospital-doctor-service:latest .

# 构建患者服务
docker build -f deploy/Dockerfile.patient -t hospital-patient-service:latest .
```

#### 2. 启动基础服务

```bash
cd deploy

# 启动etcd
docker-compose up -d etcd

# 启动MySQL
docker-compose up -d mysql

# 启动Redis
docker-compose up -d redis
```

#### 3. 启动微服务

```bash
# 启动医生服务
docker-compose up -d doctor-service

# 启动患者服务
docker-compose up -d patient-service
```

#### 4. 启动API网关

```bash
# 启动API网关
docker-compose up -d api-gateway
```

#### 5. 启动Nginx

```bash
# 启动Nginx
docker-compose up -d nginx
```

## 服务配置

### 环境变量

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| ETCD_ENDPOINTS | etcd服务地址 | etcd:2379 |
| MYSQL_HOST | MySQL主机地址 | mysql |
| MYSQL_PORT | MySQL端口 | 3306 |
| MYSQL_USER | MySQL用户名 | hospital |
| MYSQL_PASSWORD | MySQL密码 | hospital123 |
| MYSQL_DATABASE | MySQL数据库名 | hospital |
| REDIS_HOST | Redis主机地址 | redis |
| REDIS_PORT | Redis端口 | 6379 |

### 配置文件

#### API网关配置 (`api/etc/hospital.yaml`)

```yaml
Name: hospital
Host: 0.0.0.0
Port: 8888
Mode: dev

Etcd:
  Hosts:
    - etcd:2379
  Key: hospital.rpc

DoctorRpc:
  Etcd:
    Hosts:
      - etcd:2379
    Key: doctor.rpc
  NonBlock: true

PatientRpc:
  Etcd:
    Hosts:
      - etcd:2379
    Key: patient.rpc
  NonBlock: true

DataSource: hospital:hospital123@tcp(mysql:3306)/hospital?charset=utf8mb4&parseTime=True&loc=Local
```

#### 医生服务配置 (`doctor_srv/etc/doctor.yaml`)

```yaml
Name: doctor
ListenOn: 0.0.0.0:8080
Mode: dev

Etcd:
  Hosts:
    - etcd:2379
  Key: doctor.rpc

DataSource: hospital:hospital123@tcp(mysql:3306)/hospital?charset=utf8mb4&parseTime=True&loc=Local

Redis:
  Host: redis:6379
  Type: node
```

#### 患者服务配置 (`patient_srv/etc/patient.yaml`)

```yaml
Name: patient
ListenOn: 0.0.0.0:8081
Mode: dev

Etcd:
  Hosts:
    - etcd:2379
  Key: patient.rpc

DataSource: hospital:hospital123@tcp(mysql:3306)/hospital?charset=utf8mb4&parseTime=True&loc=Local

Redis:
  Host: redis:6379
  Type: node
```

## 数据库初始化

### MySQL数据库结构

系统包含以下主要数据表：

- `departments`: 科室信息
- `doctors`: 医生信息
- `patients`: 患者信息
- `time_slots`: 时间段信息
- `appointments`: 预约信息
- `medicines`: 药品信息
- `prescriptions`: 处方信息
- `prescription_items`: 处方项目
- `care_history`: 病例信息

### 初始化数据

部署时会自动执行 `deploy/mysql/init.sql` 脚本，包含：

- 创建数据库和表结构
- 插入初始科室数据
- 插入初始医生数据
- 插入初始患者数据
- 插入初始药品数据

## 服务监控

### 健康检查

```bash
# API服务健康检查
curl http://localhost:8888/departments

# Nginx健康检查
curl http://localhost/health

# etcd健康检查
curl http://localhost:2379/health
```

### 日志查看

```bash
# 查看所有服务日志
./deploy/deploy.sh logs

# 查看特定服务日志
./deploy/deploy.sh logs api-gateway
./deploy/deploy.sh logs doctor-service
./deploy/deploy.sh logs patient-service
./deploy/deploy.sh logs mysql
./deploy/deploy.sh logs redis
./deploy/deploy.sh logs etcd
./deploy/deploy.sh logs nginx
```

### 容器状态

```bash
# 查看所有容器状态
./deploy/deploy.sh status

# 或者直接使用docker-compose
cd deploy
docker-compose ps
```

## 性能优化

### 1. 数据库优化

```sql
-- 创建索引
CREATE INDEX idx_doctors_department ON doctors(department_id);
CREATE INDEX idx_appointments_patient ON appointments(patient_id);
CREATE INDEX idx_appointments_doctor ON appointments(doctor_id);
CREATE INDEX idx_time_slots_doctor_date ON time_slots(doctor_id, date);
```

### 2. Redis缓存策略

- 缓存科室列表
- 缓存医生信息
- 缓存药品信息
- 缓存用户会话

### 3. Nginx优化

- 启用Gzip压缩
- 配置静态文件缓存
- 设置合理的超时时间
- 配置负载均衡

## 安全配置

### 1. SSL证书

```bash
# 生成自签名证书
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout deploy/nginx/ssl/server.key \
    -out deploy/nginx/ssl/server.crt \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=Hospital/OU=IT/CN=your-domain.com"
```

### 2. 防火墙配置

```bash
# 只开放必要端口
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw allow 22/tcp
sudo ufw enable
```

### 3. 数据库安全

- 修改默认密码
- 限制数据库访问IP
- 定期备份数据
- 启用SSL连接

## 备份和恢复

### 数据库备份

```bash
# 备份MySQL数据
docker exec hospital-mysql mysqldump -u hospital -phospital123 hospital > backup.sql

# 恢复MySQL数据
docker exec -i hospital-mysql mysql -u hospital -phospital123 hospital < backup.sql
```

### 配置文件备份

```bash
# 备份配置文件
tar -czf config_backup.tar.gz deploy/nginx/ deploy/*.yaml

# 恢复配置文件
tar -xzf config_backup.tar.gz
```

## 故障排除

### 常见问题

#### 1. 服务无法启动

```bash
# 检查端口占用
netstat -tuln | grep :8888

# 检查容器日志
docker logs hospital-api-gateway
```

#### 2. 数据库连接失败

```bash
# 检查MySQL容器状态
docker ps | grep mysql

# 检查数据库连接
docker exec hospital-mysql mysql -u hospital -phospital123 -e "SHOW DATABASES;"
```

#### 3. etcd连接失败

```bash
# 检查etcd状态
docker logs hospital-etcd

# 检查etcd健康状态
curl http://localhost:2379/health
```

#### 4. 微服务注册失败

```bash
# 检查服务注册
curl http://localhost:2379/v2/keys/hospital.rpc

# 重启服务
./deploy/deploy.sh restart
```

### 日志分析

```bash
# 查看错误日志
docker logs hospital-api-gateway 2>&1 | grep ERROR

# 查看访问日志
docker logs hospital-nginx | grep "GET /"
```

## 扩展部署

### 1. 生产环境部署

```bash
# 修改配置文件为生产模式
sed -i 's/Mode: dev/Mode: prod/g' api/etc/hospital.yaml
sed -i 's/Mode: dev/Mode: prod/g' doctor_srv/etc/doctor.yaml
sed -i 's/Mode: dev/Mode: prod/g' patient_srv/etc/patient.yaml
```

### 2. 负载均衡

```nginx
upstream api_backend {
    server api-gateway-1:8888;
    server api-gateway-2:8888;
    server api-gateway-3:8888;
}
```

### 3. 监控集成

- 集成Prometheus监控
- 配置Grafana仪表板
- 设置告警规则

## 维护命令

```bash
# 更新代码后重新部署
git pull
./deploy/deploy.sh deploy

# 查看资源使用情况
docker stats

# 清理未使用的镜像
docker image prune -f

# 清理未使用的容器
docker container prune -f

# 清理未使用的网络
docker network prune -f

# 清理未使用的卷
docker volume prune -f
```

## 联系信息

如有部署问题，请联系开发团队。 