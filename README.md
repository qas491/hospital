# 医院管理系统后端

## 项目简介

本项目为医院管理系统的后端服务，采用 Go 语言开发，包含医生端、患者端和 API 网关等模块。系统支持医院常见的业务流程，包括预约挂号、处方管理、用药管理、医生绩效统计等，旨在提升医院信息化水平和服务效率。

## 目录结构

```
hospital/
├── api/                # API 网关服务，定义接口协议与路由
│   ├── etc/            # 配置文件
│   ├── internal/       # 业务逻辑与处理器
│   └── hospital.api    # API 定义文件
├── doctor_srv/         # 医生端微服务
│   ├── configs/        # 配置文件
│   ├── doctor/         # gRPC 相关代码
│   ├── doctorservice/  # 医生服务实现
│   ├── internal/       # 业务逻辑
│   └── model/          # 数据库模型
├── patient_srv/        # 患者端微服务
│   ├── configs/        # 配置文件
│   ├── internal/       # 业务逻辑
│   ├── medicalservice/ # 患者服务实现
│   └── model/          # 数据库模型
├── default.etcd/       # etcd 相关数据
├── go.work             # Go 多模块工作区配置
└── LICENSE             # 许可证
```

## 主要功能

- **预约挂号**：患者可在线预约医生，选择科室和时间段。
- **处方管理**：医生可为患者开具处方，支持处方审核与历史查询。
- **用药管理**：支持药品信息查询与选择。
- **医生绩效统计**：统计医生的工作量和绩效排名。
- **多端支持**：医生端、患者端服务分离，便于扩展和维护。

## 技术栈

- 语言：Go
- 通信协议：gRPC、HTTP
- 配置中心：etcd
- 数据库：MySQL、Redis
- 架构：微服务

## 快速开始

### 1. 环境依赖

- Go 1.18 及以上
- MySQL 5.7 及以上
- Redis 5 及以上
- etcd 3.x

### 2. 配置说明

- 各服务配置文件位于 `api/etc/`、`doctor_srv/configs/`、`patient_srv/configs/` 目录下。
- 数据库、Redis、etcd 连接信息需根据实际环境修改。
- 参考各模块下的 `config.yaml` 或 `dev.yaml` 文件进行配置。

### 3. 启动服务

依次启动各个服务（建议使用不同终端窗口）：

```bash
# 启动医生端服务
go run ./doctor_srv/doctor.go

# 启动患者端服务
go run ./patient_srv/patient.go

# 启动 API 网关服务
go run ./api/hospital.go
```

### 4. 访问接口

- 通过 API 网关暴露的 HTTP/gRPC 接口访问系统功能。
- 可结合 Postman 或 Swagger 进行接口调试。

## 许可证

本项目采用 MIT 许可证，详见 [LICENSE](./LICENSE)。 