package main

import (
	"context"
	"log"

	"github.com/leeprince/protobuf/grpc/greeter"
	"github.com/asim/go-micro/v3"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *greeter.GreeterReq, rsp *greeter.GreeterRsp) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter.server"),
		// micro.Address(":8081"),
	)
	
	service.Init()

	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
