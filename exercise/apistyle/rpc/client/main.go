package main

import (
	"context"
	pb "github.com/wangjian8774/my-iam/exercise/apistyle/rpc/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// 创建一个gRPC连接，用于和服务端通信
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 创建客户端stub
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 1.调用方便：RPC屏蔽了底层的网络通信细节，像调用本地方法一样调用远端方法
	// 2.避免了打包和解包：RPC的调用参数和返回结果都是Go的结构体，简化调用步骤
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
}
