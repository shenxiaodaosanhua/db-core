package main

import (
	"context"
	"db-core/pbfiles"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

func main() {
	client, err := grpc.DialContext(context.Background(),
		"localhost:8080",
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal(err)
	}

	paramStruct, err := structpb.NewStruct(map[string]interface{}{
		"username": "jerry4",
		"password": "123123",
		"mobile":   "18011801983",
	})
	params := &pbfiles.SimpleParams{
		Params: paramStruct,
	}

	req := &pbfiles.ExecRequest{Name: "add_user", Params: params}
	rsp := &pbfiles.ExecResponse{}
	err = client.Invoke(context.Background(),
		"/DBService/Exec", req, rsp)
	if err != nil {
		log.Fatal(err)
	}
	//for _, item := range rsp.Result {
	//	fmt.Println(item.AsMap())
	//}
	fmt.Println(rsp.Select.AsMap())
}
