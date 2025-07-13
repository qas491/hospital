# API 网关服务（api）

## 模块简介

本模块为医院管理系统的 API 网关服务，负责统一对外提供 HTTP/gRPC 接口，路由转发请求到医生端和患者端微服务，实现接口聚合、参数校验等功能。采用 Go 语言开发，便于与后端各微服务高效集成。

## 目录结构

```
api/
├── etc/                # 配置文件目录
│   └── hospital.yaml   # 服务配置
├── go.mod              # Go 依赖管理
├── hospital.api        # API 协议定义（如使用 go-zero）
├── hospital.go         # 服务入口
├── internal/           # 内部实现
│   ├── config/         # 配置加载
│   ├── handler/        # 路由与请求处理
│   │   ├── doctor/     # 医生相关接口处理
│   │   ├── patient/    # 患者相关接口处理
│   │   └── routes.go   # 路由注册
│   ├── logic/          # 业务逻辑
│   ├── svc/            # 服务上下文
│   └── types/          # 类型定义
```

## 主要功能

- **统一 API 网关**：对外暴露 HTTP/gRPC 接口，聚合后端服务能力。
- **路由转发**：根据请求类型转发到医生端、患者端等微服务。
- **参数校验**：对请求参数进行校验。
- **接口聚合**：支持多服务数据整合返回。

## 技术栈

- 语言：Go
- 框架：go-zero（如适用）
- 通信协议：HTTP、gRPC
- 配置中心：etcd（可选）

## 环境依赖

- Go 1.18 及以上
- etcd 3.x（如需注册中心）
- 需先启动 doctor_srv、patient_srv 等后端微服务

## 配置说明

- 配置文件位于 `etc/hospital.yaml`。
- 需根据实际环境修改服务端口、后端微服务地址等信息。

## 启动方式

在项目根目录下执行：

```bash
cd api
# 启动 API 网关服务
go run hospital.go
```

## 简要接口说明

API 网关对外暴露的主要接口包括（以 RESTful 风格为例）：

- `/api/patient/appointment`：预约挂号、查询、取消等
- `/api/doctor/prescription`：处方相关操作
- `/api/doctor/medicine`：药品相关操作
- `/api/doctor/ranking`：医生绩效排名

详细接口定义见 `hospital.api` 文件。

## 许可证

本模块采用 MIT 许可证，详见项目根目录 [LICENSE](../LICENSE)。 