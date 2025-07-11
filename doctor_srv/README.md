# 医生服务 (Doctor Service)

基于go-zero框架开发的医院医生服务，提供开处方、审核处方、选择药品等功能。

## 功能特性

### 1. 开处方功能
- **高并发库存管理**: 使用悲观锁和事务确保库存不会超卖
- **参数验证**: 完整的请求参数验证
- **事务处理**: 确保数据一致性
- **操作日志**: 记录所有操作日志

### 2. 审核处方功能
- **状态管理**: 支持通过、拒绝、待审核等状态
- **库存恢复**: 审核拒绝时自动恢复库存
- **通知机制**: 审核结果通知患者
- **日志记录**: 完整的审核日志

### 3. 选择药品功能
- **多条件查询**: 支持按ID、名称、类型等条件查询
- **库存检查**: 实时检查药品库存状态
- **状态验证**: 验证药品是否可用
- **危险库存预警**: 库存不足时发出警告

### 4. 病例管理
- **创建病例**: 完整的病例创建流程
- **数据验证**: 验证患者和医生信息
- **通知机制**: 病例创建后通知相关人员

## 高并发库存管理解决方案

### 问题分析
在高并发场景下，多个医生同时开处方可能导致库存超卖问题。

### 解决方案
1. **悲观锁**: 使用 `SELECT FOR UPDATE` 锁定库存记录
2. **事务处理**: 确保库存检查和扣减在同一个事务中
3. **原子操作**: 使用 `UPDATE` 语句的原子性确保库存扣减的准确性
4. **库存预检查**: 在扣减前先检查库存是否充足

### 核心代码示例
```go
// 检查并锁定库存
func (l *CreatePrescriptionLogic) checkAndLockInventory(tx *gorm.DB, items []*doctor.PrescriptionItem) error {
    for _, item := range items {
        // 使用SELECT FOR UPDATE锁定库存记录
        var medicine mysql.StockMedicines
        if err := tx.Set("gorm:query_option", "FOR UPDATE").
            Where("medicines_id = ? AND del_flag = '0'", item.ItemRefId).
            First(&medicine).Error; err != nil {
            return fmt.Errorf("药品不存在: %s", item.ItemRefId)
        }

        // 检查库存是否充足
        if medicine.MedicinesStockNum < item.Num {
            return fmt.Errorf("药品 %s 库存不足", medicine.MedicinesName)
        }
    }
    return nil
}

// 扣减库存
func (l *CreatePrescriptionLogic) deductInventory(tx *gorm.DB, medicinesId string, num float64) error {
    result := tx.Model(&mysql.StockMedicines{}).
        Where("medicines_id = ? AND medicines_stock_num >= ? AND del_flag = '0'", medicinesId, num).
        Update("medicines_stock_num", gorm.Expr("medicines_stock_num - ?", num))

    if result.RowsAffected == 0 {
        return fmt.Errorf("库存不足或药品不存在")
    }
    return nil
}
```

## 项目结构

```
doctor_srv/
├── configs/                 # 配置文件
│   └── config.yaml
├── internal/
│   ├── config/             # 配置结构
│   │   └── config.go
│   ├── logic/              # 业务逻辑层
│   │   ├── createprescriptionlogic.go
│   │   ├── reviewprescriptionlogic.go
│   │   ├── selectmedicineslogic.go
│   │   ├── getmedicineslistlogic.go
│   │   ├── getprescriptionlistlogic.go
│   │   └── createcarehistorylogic.go
│   └── svc/                # 服务上下文
│       └── servicecontext.go
├── model/                  # 数据模型
│   └── mysql/
│       └── sxt_his.go
├── doctor.proto           # gRPC协议定义
├── go.mod                 # Go模块文件
└── README.md             # 项目说明
```

## 安装和运行

### 1. 安装依赖
```bash
go mod tidy
```

### 2. 配置数据库
修改 `configs/config.yaml` 中的数据库连接信息：
```yaml
Database:
  DataSource: username:password@tcp(localhost:3306)/hospital?charset=utf8mb4&parseTime=True&loc=Local
```

### 3. 生成gRPC代码
```bash
protoc --go_out=. --go-grpc_out=. doctor.proto
```

### 4. 运行服务
```bash
go run main.go
```

## API接口

### 开处方
- **接口**: `CreatePrescription`
- **功能**: 医生开处方，包含库存检查和扣减
- **特点**: 高并发安全，防止超卖

### 审核处方
- **接口**: `ReviewPrescription`
- **功能**: 审核处方，支持通过/拒绝
- **特点**: 审核拒绝时自动恢复库存

### 选择药品
- **接口**: `SelectMedicines`
- **功能**: 查询和选择药品
- **特点**: 实时库存检查，多条件查询

### 获取药品列表
- **接口**: `GetMedicinesList`
- **功能**: 分页查询药品列表
- **特点**: 支持多种筛选条件

### 创建病例
- **接口**: `CreateCareHistory`
- **功能**: 创建患者病例
- **特点**: 完整的数据验证和通知机制

## 性能优化

1. **数据库索引**: 在关键字段上建立索引
2. **连接池**: 配置合适的数据库连接池大小
3. **缓存**: 对常用数据进行缓存
4. **异步处理**: 日志记录和通知发送使用异步处理

## 监控和日志

- **操作日志**: 记录所有关键操作
- **错误日志**: 详细的错误信息记录
- **性能监控**: 接口响应时间监控
- **库存监控**: 实时库存状态监控

## 注意事项

1. **数据库事务**: 确保所有涉及多表操作的功能都使用事务
2. **库存安全**: 严格遵循库存检查和扣减的流程
3. **并发控制**: 使用适当的锁机制防止并发问题
4. **数据验证**: 对所有输入参数进行严格验证
5. **错误处理**: 完善的错误处理和回滚机制 