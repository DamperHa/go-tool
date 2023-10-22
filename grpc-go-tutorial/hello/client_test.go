package hello

import (
	"context"
	pb "go-tool/grpc-go-tutorial/hello/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func TestClient(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
