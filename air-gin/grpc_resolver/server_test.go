package grpc_resolver

import (
	"air-gin/proto/gen"
	"context"
	"fmt"
	"net"
	"testing"

	"google.golang.org/grpc"
)

type Server struct {
	// 组合没有实现的UserService结构体
	gen.UnimplementedUserServiceServer
}

// 实现服务方法
func (s Server) GetById(ctx context.Context, req *gen.GetByIdReq) (*gen.GetByIdResp, error) {
	fmt.Println(req)
	return &gen.GetByIdResp{
		User: &gen.User{
			Name: "hello, world",
		},
	}, nil
}

func TestServer(t *testing.T) {
	us := &Server{}

	server := grpc.NewServer()

	gen.RegisterUserServiceServer(server, us)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		t.Fatal(err)
	}

	err = server.Serve(l)
	if err != nil {
		t.Fatal(err)
	}
}
