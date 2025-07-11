package main

import (
	"flag"
	"fmt"
	"github.com/qas491/hospital/doctor_srv/configs"
	"github.com/qas491/hospital/doctor_srv/model/mysql"
	"github.com/qas491/hospital/doctor_srv/model/redis"

	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/doctor_srv/internal/config"
	"github.com/qas491/hospital/doctor_srv/internal/server"
	"github.com/qas491/hospital/doctor_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/doctor.yaml", "the config file")

func main() {
	flag.Parse()
	configs.Init()
	mysql.MysqlInit()
	redis.ExampleClient()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		doctor.RegisterDoctorServiceServer(grpcServer, server.NewDoctorServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
