package main

import (
	"context"
	"db-core/pbfiles"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

func GenParam(m map[string]interface{}) *pbfiles.SimpleParams {
	paramStruct, err := structpb.NewStruct(m)
	CheckError(err)

	params := &pbfiles.SimpleParams{
		Params: paramStruct,
	}
	return params
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	client, err := grpc.DialContext(context.Background(),
		"localhost:8080",
		grpc.WithInsecure(),
	)
	CheckError(err)

	c := pbfiles.NewDBServiceClient(client)
	txClient, err := c.Tx(context.Background())
	CheckError(err)

	params := GenParam(map[string]interface{}{
		"username": "jerry",
		"password": "123123",
		"mobile":   "18011801980",
	})

	err = txClient.Send(&pbfiles.TxRequest{
		Name:   "add_user",
		Params: params,
		Type:   "exec",
	})
	CheckError(err)

	response, err := txClient.Recv()
	CheckError(err)
	fmt.Println("返回值：", response)
	result := response.Result.AsMap()
	userId := result["exec"].([]interface{})[1].(map[string]interface{})["user_id"]
	fmt.Println("用户ID：", userId)

	addAmount := GenParam(map[string]interface{}{
		"user_id": userId,
		"amount":  "10000",
	})
	//log.Fatalln("111")
	err = txClient.Send(&pbfiles.TxRequest{
		Name:   "add_user_amounts",
		Params: addAmount,
		Type:   "exec",
	})
	CheckError(err)

	response, err = txClient.Recv()
	CheckError(err)
	fmt.Println("返回值：", response)
	fmt.Println(response.Result.AsMap())

	err = txClient.CloseSend()
	fmt.Println("操作完成关闭", err)
}
