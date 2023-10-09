package example1

import (
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/DamplerHa/grpc-go-tutorial/example1/proto"
)

func TestSimpleClient(t *testing.T) {
	// 获取一个底层连接
	conn, err := grpc.Dial("127.0.0.1:"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	// 获取对应服务的客户端
	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
