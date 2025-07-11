package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/qas491/hospital/doctor_srv/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Database.DataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
			DB:       0,
		}),
	}
}
