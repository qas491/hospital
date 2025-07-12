# 患者端服务（patient_srv）

## 模块简介

本模块为医院管理系统的患者端微服务，负责处理患者相关的业务逻辑，包括预约挂号、查询医生与科室、查看和取消预约等。服务采用 Go 语言开发，支持 gRPC 通信协议，作为后端微服务的一部分与其他模块协同工作。

## 目录结构

```
patient_srv/
├── configs/         # 配置文件目录
│   ├── config.go    # 配置结构体定义
│   ├── dev.yaml     # 开发环境配置示例
├── etc/             # 运行时配置文件
│   └── patient.yaml # 服务配置
├── internal/        # 内部实现
│   ├── config/      # 配置加载
│   ├── logic/       # 业务逻辑实现
│   ├── server/      # gRPC 服务注册
│   └── svc/         # 服务上下文
├── medicalservice/  # 患者服务实现
├── model/           # 数据库模型
│   ├── mysql/       # MySQL 相关
│   └── redis/       # Redis 相关
├── patient/         # gRPC 相关代码
│   ├── patient_grpc.pb.go
│   └── patient.pb.go
├── patient.go       # 服务入口
├── patient.proto    # gRPC 协议定义
```

## 主要功能

- **预约挂号**：患者可选择科室、医生和时间段进行预约。
- **查询医生与科室**：支持按科室查询医生列表。
- **查看预约**：患者可查询自己的预约信息。
- **取消预约**：支持患者主动取消未完成的预约。

## 技术栈

- 语言：Go
- 通信协议：gRPC
- 配置中心：etcd（可选）
- 数据库：MySQL、Redis

## 环境依赖

- Go 1.18 及以上
- MySQL 5.7 及以上
- Redis 5 及以上
- etcd 3.x（如需注册中心）

## 配置说明

- 配置文件位于 `configs/` 和 `etc/` 目录下。
- 需根据实际环境修改数据库、Redis、etcd 等连接信息。
- 参考 `dev.yaml` 或 `patient.yaml` 进行配置。

## 启动方式

在项目根目录下执行：

```bash
cd patient_srv
# 启动患者端服务
go run patient.go
```

## 简要接口说明

服务采用 gRPC 协议，主要接口包括：

- `ListDepartments`：获取科室列表
- `ListDoctors`：获取指定科室的医生列表
- `ListTimeSlots`：获取医生可预约时间段
- `MakeAppointment`：创建预约
- `GetAppointment`：查询预约详情
- `CancelAppointment`：取消预约

详细接口定义见 `patient.proto` 文件。

## 许可证

本模块采用 MIT 许可证，详见项目根目录 [LICENSE](../LICENSE)。 