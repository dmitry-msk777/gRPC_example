package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dmitry-msk777/gRPC_example/proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5300", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewCRMswapClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	S := &pb.S{
		CustomerId:    "123007", //"123007_delete_this_postfix",
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
		log.Println("something went wrong", err)
	}
	log.Println("rply2: ", rply2.Response)

	//If
	// message SF {
	// 	repeated S S = 1;
	// 	F F = 2;
	// }

	// Then

	// var Array_S []*pb.S

	// S := &pb.S{
	// 	CustomerId:    "123007",
	// 	CustomerName:  "Alex gRPC",
	// 	CustomerType:  "Google",
	// 	CustomerEmail: "gRPC@mail.ru",
	// }
	// Array_S = append(Array_S, S)

	// S = &pb.S{
	// 	CustomerId:    "123007",
	// 	CustomerName:  "Alex gRPC",
	// 	CustomerType:  "Google",
	// 	CustomerEmail: "gRPC@mail.ru",
	// }
	//Array_S = append(Array_S, S)

}
