package example1

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/DamplerHa/grpc-go-tutorial/example1/proto"
)

type SearchService struct {
	pb.UnimplementedSearchServiceServer
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Println("handle Request")
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func TestSimpleServer(t *testing.T) {
	// 1. 生成gprc服务端
	server := grpc.NewServer()

	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", "127.0.0.1:"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}

type StreamService struct {
	pb.UnimplementedStreamServiceServer
}

// List 服务端流式rpc
func (s *StreamService) List(point *pb.StreamPoint, server pb.StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := server.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  point.Name,
				Value: point.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StreamService) Record(server pb.StreamService_RecordServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *StreamService) Route(server pb.StreamService_RouteServer) error {
	//TODO implement me
	panic("implement me")
}

func TestStreamService(t *testing.T) {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err:%v", err)
	}

	server.Serve(lis)
}
