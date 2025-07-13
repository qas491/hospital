# 医院管理系统API测试指南

## 概述

本文档提供了医院管理系统API的完整测试方案，包括多种测试方法和工具。

## 测试前准备

### 1. 启动服务

在测试API之前，需要确保以下服务正在运行：

```bash
# 1. 启动etcd服务（如果还没有运行）
# etcd服务应该在2379端口运行

# 2. 启动doctor微服务
cd doctor_srv
go run doctor.go

# 3. 启动patient微服务
cd patient_srv
go run patient.go

# 4. 启动API网关
cd api
go run hospital.go
```

## 测试方法

### 方法1: 使用Python测试脚本

```bash
# 安装依赖
pip install requests

# 运行测试
python test_api.py
```

**特点：**
- 完整的测试覆盖
- 详细的测试报告
- 易于扩展和修改

### 方法2: 使用Bash脚本

```bash
# 给脚本执行权限
chmod +x test_api.sh

# 运行测试
./test_api.sh
```

**特点：**
- 轻量级
- 不依赖Python环境
- 适合CI/CD集成

### 方法3: 使用Postman

1. 导入 `hospital_api_tests.postman_collection.json` 到Postman
2. 设置环境变量：
   - `base_url`: `http://localhost:8888`
3. 运行测试集合

**特点：**
- 图形化界面
- 易于调试
- 支持环境变量

### 方法4: 使用curl命令

```bash
# 测试获取科室列表
curl http://localhost:8888/departments

# 测试预约挂号
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{
       "patient_id": 1,
       "doctor_id": 1,
       "department_id": 1,
       "timeslot_id": 1
     }' \
     http://localhost:8888/appointments
```

## API接口列表

### 患者端接口

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 获取科室列表 | GET | `/departments` | 获取所有科室信息 |
| 获取医生列表 | GET | `/departments/{id}/doctors` | 获取指定科室的医生 |
| 获取时间段列表 | GET | `/doctors/{id}/timeslots` | 获取医生可预约时间段 |
| 预约挂号 | POST | `/appointments` | 创建预约 |
| 获取预约详情 | GET | `/appointments/{id}` | 获取预约详细信息 |
| 取消预约 | POST | `/appointments/{id}/cancel` | 取消预约 |

### 医生端接口

| 接口 | 方法 | 路径 | 描述 |
|------|------|------|------|
| 创建处方 | POST | `/prescriptions` | 创建新处方 |
| 获取处方列表 | GET | `/prescriptions` | 获取处方列表 |
| 获取处方详情 | GET | `/prescriptions/{id}` | 获取处方详细信息 |
| 审核处方 | POST | `/prescriptions/{id}/review` | 审核处方 |
| 获取药品列表 | GET | `/medicines` | 获取药品列表 |
| 获取药品详情 | GET | `/medicines/{id}` | 获取药品详细信息 |
| 选择药品 | POST | `/medicines/select` | 选择药品 |
| 创建病例 | POST | `/care-history` | 创建病例记录 |
| 获取病例列表 | GET | `/care-history` | 获取病例列表 |
| 获取病例详情 | GET | `/care-history/{id}` | 获取病例详细信息 |
| 获取周排行榜 | GET | `/rankings/weekly` | 获取医生周排行榜 |
| 获取医生业绩 | GET | `/doctors/{id}/performance` | 获取医生业绩统计 |

## 测试数据示例

### 预约挂号请求
```json
{
  "patient_id": 1,
  "doctor_id": 1,
  "department_id": 1,
  "timeslot_id": 1
}
```

### 创建处方请求
```json
{
  "co_id": "PRES_1234567890",
  "co_type": "prescription",
  "user_id": 1,
  "patient_id": "P001",
  "patient_name": "张三",
  "ch_id": "CH001",
  "all_amount": 150.50,
  "create_by": "doctor001",
  "items": [
    {
      "item_id": "ITEM001",
      "item_ref_id": "MED001",
      "item_name": "阿莫西林胶囊",
      "item_type": "medicine",
      "num": 2.0,
      "price": 25.00,
      "amount": 50.00,
      "remark": "每日3次，每次1粒",
      "status": "active"
    }
  ]
}
```

### 创建病例请求
```json
{
  "ch_id": "CH_1234567890",
  "user_id": 1,
  "user_name": "李医生",
  "patient_id": "P001",
  "patient_name": "张三",
  "dept_id": 1,
  "dept_name": "内科",
  "receive_type": "门诊",
  "is_contagious": "否",
  "case_date": "2024-01-15",
  "reg_id": "REG001",
  "case_title": "感冒",
  "case_result": "上呼吸道感染",
  "doctor_tips": "多休息，多喝水",
  "remark": "患者症状轻微"
}
```

## 常见问题

### 1. 连接被拒绝
- 检查API服务是否启动
- 确认端口8888是否被占用
- 检查防火墙设置

### 2. 微服务连接失败
- 确认etcd服务正在运行
- 检查微服务是否已注册到etcd
- 验证服务配置是否正确

### 3. 数据库连接失败
- 检查数据库服务是否运行
- 验证数据库连接配置
- 确认数据库表结构正确

## 性能测试

可以使用Apache Bench (ab) 进行性能测试：

```bash
# 测试获取科室列表接口
ab -n 1000 -c 10 http://localhost:8888/departments
```

## 监控和日志

- API服务日志：查看控制台输出
- 微服务日志：查看各服务的控制台输出
- 数据库日志：查看MySQL日志
- etcd日志：查看etcd服务日志

## 扩展测试

可以根据需要添加更多测试用例：

1. 边界值测试
2. 错误处理测试
3. 并发测试
4. 安全性测试
5. 集成测试 