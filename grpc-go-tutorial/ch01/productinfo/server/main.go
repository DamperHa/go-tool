package main

import (
	"context"
	pb "go-tool/grpc-go-tutorial/ch01/productinfo/server/ecommerce"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gofrs/uuid"
)

const (
	port = ":50001"
)

type server struct {
	pb.UnimplementedProductInfoServer
	productMap map[string]*pb.Product
}

// AddProduct 添加产品，并返回ProductID
func (s *server) AddProduct(ctx context.Context, product *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID, err:%v", err)
	}

	product.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[product.Id] = product
	log.Printf("Product %v: %v - Added.", product.Id, product.Name)
	return &pb.ProductID{
		Value: product.Id,
	}, status.New(codes.OK, "").Err()
}

func (s *server) GetProduct(ctx context.Context, id *pb.ProductID) (*pb.Product, error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	product, exist := s.productMap[id.Value]
	if !exist {
		return nil, status.Errorf(codes.NotFound, "Product does not exist.", id.Value)
	}

	return product, status.New(codes.OK, "").Err()
}

func main() {
	// 1. 获取基础通信的套接字
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
