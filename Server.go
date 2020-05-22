package main

import (
	"fmt"

	pb "./proto"

	//pb "github.com/dmitry-msk777/CRM_Test/proto"
	"context"
	"net"
	"reflect"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Customer_struct struct {
	Customer_id    string
	Customer_name  string
	Customer_type  string
	Customer_email string
}

type server struct{}

func (s *server) A(ctx context.Context, in *pb.SF) (*pb.Response, error) {

	structValue := *in.S
	fields := reflect.TypeOf(structValue)
	values := reflect.ValueOf(structValue)

	num := fields.NumField()

	var mapCheck = make(map[string]string)

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		mapCheck[field.Name] = value.String()
		//fmt.Print("Type:", field.Type, ",", field.Name, "=", value, "\n")
	}

	CheckSuccess := true

	for key, value := range in.F.Keyvalue {
		//fmt.Println("map:", key, value)

		val1, flag := mapCheck[key]
		if flag == false {
			CheckSuccess = false
			break
		}

		if val1 != value {
			CheckSuccess = false
			break
		}

		//Descriptor_byte, Descriptor_int := in.S.Descriptor()
		// in.S.ProtoMessage()
		// in.S.GetCustomerEmail()
		// in.S.Reset()
		// in.S.String()

		// fmt.Println("Descriptor:", Descriptor_byte, Descriptor_int)

		// V_Number, V_Type, V_int := protowire.ConsumeField(Descriptor_byte)
		// fmt.Println("ConsumeField:", V_Number, V_Type, V_int)

	}

	response := &pb.Response{
		Response: CheckSuccess,
	}

	return response, nil
}

// func (s *server) POST_List(ctx context.Context, in *pb.RequestPOST) (*pb.ResponsePOST, error) {

// 	// Customer_struct_out := Customer_struct{
// 	// 	Customer_id:    in.CustomerId,
// 	// 	Customer_name:  in.CustomerName,
// 	// 	Customer_type:  in.CustomerType,
// 	// 	Customer_email: in.CustomerEmail,
// 	// }

// 	//err := EngineCRMv.AddChangeOneRow(EngineCRMv.DataBaseType, Customer_struct_out)

// 	//if err != nil {
// 		//ErrorLogger.Println(err.Error())
// 	//	return nil, err
// 	//}

// 	return &pb.ResponsePOST{CustomerId: "True"}, nil
// }

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterCRMswapServer(grpcServer, &server{})
	grpcServer.Serve(listener)

	fmt.Println("123")

}
