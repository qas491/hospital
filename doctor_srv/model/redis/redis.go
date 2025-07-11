package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/qas491/hospital/doctor_srv/configs"
)

var RDB *redis.Client

func ExampleClient() {
	con := configs.WiseConfig.Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     con.Addr,
		Password: con.Password, // no password set
		DB:       con.Db,       // use default DB
	})

	fmt.Println("redis connect success")
}
