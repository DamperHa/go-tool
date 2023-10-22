package hello

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "go-tool/grpc-go-tutorial/hello/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
	"time"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	logrus.Info("request in")

	time.Sleep(5 * time.Second)

	//select {
	//case <-ctx.Done():
	//	fmt.Println("time out Done")
	//}

	logrus.Info("requst out")

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("RPC has reached deadline exceeded state: %s", ctx.Err())
		return nil, ctx.Err()
	}

	return &pb.HelloReply{Message: "Hello, " + request.Name}, nil
}

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
