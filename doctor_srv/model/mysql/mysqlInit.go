package mysql

import (
	"fmt"
	"log"

	"github.com/qas491/hospital/doctor_srv/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// MysqlInit 初始化MySQL数据库连接
// 从配置中读取数据库连接信息并建立连接
func MysqlInit() {
	con := configs.WiseConfig.Mysql
	var err error

	// 检查配置是否完整
	if con.Host == "" || con.Database == "" {
		log.Println("数据库配置不完整，请检查配置文件")
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		con.User, con.Password, con.Host, con.Port, con.Database)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("数据库连接失败: %v", err)
		return
	}

	// 测试数据库连接
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("获取数据库实例失败: %v", err)
		return
	}

	// 测试连接
	if err := sqlDB.Ping(); err != nil {
		log.Printf("数据库连接测试失败: %v", err)
		return
	}

	fmt.Println("MySQL连接成功")
}

// GetDB 获取数据库连接，如果连接为nil则返回错误
func GetDB() (*gorm.DB, error) {
	if DB == nil {
		return nil, fmt.Errorf("数据库连接未初始化")
	}
	return DB, nil
}
