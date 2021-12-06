package main

import (
	"db-core/core"
	"db-core/pbfiles"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	core.InitConfig()
	core.InitDB()

	s := grpc.NewServer()
	pbfiles.RegisterDBServiceServer(s, &core.DbService{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
