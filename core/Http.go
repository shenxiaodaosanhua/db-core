package core

import (
	"fmt"
	"log"
	"net/http"
)

func InitHttp() {
	go func() {
		http.HandleFunc("/reload", func(writer http.ResponseWriter, request *http.Request) {
			InitConfig()
		})
		fmt.Println(SysConfig.ServerConfig)
		port := SysConfig.ServerConfig.HttpPort
		log.Println(fmt.Sprintf("启动内部服务，端口:%d", port))
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			log.Fatalln(err)
		}
	}()
}
