package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yuyuancha/go-practice/grpc/grpc-example-01/hello"
	"google.golang.org/grpc"
)

func main() {
	var gRPCServer = grpc.NewServer()

	hello.RegisterHelloServiceServer(gRPCServer, new(hello.HelloService))

	var listen, err = net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("gRPC 服務正在運作...")

	log.Fatal(gRPCServer.Serve(listen))
}
