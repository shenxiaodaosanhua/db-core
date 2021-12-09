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
		"id": 1,
	})
	params := &pbfiles.SimpleParams{
		Params: paramStruct,
	}

	req := &pbfiles.QueryRequest{Name: "userlist", Params: params}
	rsp := &pbfiles.QueryResponse{}
	err = client.Invoke(context.Background(),
		"/DBService/Query", req, rsp)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range rsp.Result {
		fmt.Println(item.AsMap())
	}
}
