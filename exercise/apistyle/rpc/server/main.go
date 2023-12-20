package main

import (
	"context"
	pb "github.com/wangjian8774/my-iam/exercise/apistyle/rpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// 指定监听客户端请求的端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	//创建gRPC server实例
	s := grpc.NewServer()
	// 将服务注册到gRPC框架中
	pb.RegisterGreeterServer(s, &server{})
	// 启动gRPC服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
