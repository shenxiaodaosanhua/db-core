package main

import (
	"context"
	"db-core/pbfiles"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

func MakeParam(m map[string]interface{}) *pbfiles.SimpleParams {
	paramStruct, _ := structpb.NewStruct(m)
	return &pbfiles.SimpleParams{
		Params: paramStruct,
	}
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

	txClient, err := c.Tx(context.Background()) //执行事务
	CheckError(err)

	addUserParam := MakeParam(map[string]interface{}{
		"username": "jerry",
		"password": "123",
		"mobile":   "18011801980",
	})
	err = txClient.Send(&pbfiles.TxRequest{Name: "add_user", Params: addUserParam, Type: "exec"})
	CheckError(err)

	addUserRsp, err := txClient.Recv()
	CheckError(err)

	ret := addUserRsp.Result.AsMap()
	uid := ret["exec"].([]interface{})[1].(map[string]interface{})["user_id"]
	fmt.Println("用户ID是", uid)

	//log.Fatal("abc")
	addScoreParam := MakeParam(map[string]interface{}{
		"user_id": uid,
		"amount":  3, //送三个积分
	})
	err = txClient.Send(&pbfiles.TxRequest{Name: "add_user_amounts", Params: addScoreParam, Type: "exec"})
	CheckError(err)

	addScoreRsp, err := txClient.Recv()
	CheckError(err)

	fmt.Println(addScoreRsp.Result.AsMap())
	err = txClient.CloseSend()
	fmt.Println("结束")

}
