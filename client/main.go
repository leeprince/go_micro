package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/leeprince/protobuf/grpc/greeter"
	"log"
)

func main() {
	// 获取服务实例
	service := micro.NewService(
		micro.Name("greeter.client"),
	)
	
	// 初始化服务
	service.Init()
	
	// 创建新的客户端
	greeterSrv := greeter.NewGreeterService("greeter.server", service.Client())
	
	// 调用 grpc 服务端的方法
	res, err := greeterSrv.Hello(context.Background(), &greeter.GreeterReq{
		Name: "prince",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("返回结果：", res)
	
}
