package unitver1

//go test

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	"log"
	"time"

	pb "../proto"
)

func TestRPCClient(t *testing.T) {

	conn, err := grpc.Dial("localhost:5300", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewCRMswapClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t.Run("SayHello", func(t *testing.T) {

		S := &pb.S{
			CustomerId:    "123007_del",
			CustomerName:  "Alex gRPC",
			CustomerType:  "Google",
			CustomerEmail: "gRPC@mail.ru",
		}

		var Fmap = make(map[string]string)
		Fmap["CustomerId"] = "123007"
		Fmap["CustomerName"] = "Alex gRPC"
		Fmap["CustomerType"] = "Google"
		Fmap["CustomerEmail"] = "gRPC@mail.ru"

		F := &pb.F{
			Keyvalue: Fmap,
		}

		SF := &pb.SF{
			S: S,
			F: F,
		}

		rply2, err := c.A(ctx, SF)
		if err != nil {
			t.Error("Expected 'Hello world', got ", rply2)
			//log.Println("something went wrong", err)
		}

		t.Logf("Greeting: %s", rply2)

	})

	// // Set up a connection to the Server.
	// const address = "localhost:50051"
	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	// if err != nil {
	// 	t.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// c := pb.NewGreeterClient(conn)

	// // Test SayHello
	// t.Run("SayHello", func(t *testing.T) {
	// 	name := "world"
	// 	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	// 	if err != nil {
	// 		t.Fatalf("could not greet: %v", err)
	// 	}
	// 	t.Logf("Greeting: %s", r.Message)
	// 	if r.Message != "Hello "+name {
	// 		t.Error("Expected 'Hello world', got ", r.Message)
	// 	}

	// })
}
