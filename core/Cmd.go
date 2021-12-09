package core

import (
	"db-core/pbfiles"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "db-core",
	Short: "DB服务",
	Long:  `DB服务`,
	Run: func(cmd *cobra.Command, args []string) {
		InitConfig()
		InitDB()
		InitHttp()

		s := grpc.NewServer()
		pbfiles.RegisterDBServiceServer(s, &DbService{})

		grpcPort := SysConfig.ServerConfig.RpcPort
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(fmt.Sprintf("db-core服务启动，端口:%d", grpcPort))

		if err = s.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of db-core",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
}
var reloadConfig = &cobra.Command{
	Use:   "reload",
	Short: "reload sysConfig(app.yaml)",
	Run: func(cmd *cobra.Command, args []string) {
		InitConfig()
		httpPort := SysConfig.ServerConfig.HttpPort
		rsp, err := http.Get(fmt.Sprintf("http://localhost:%d/reload", httpPort))
		if err != nil {
			log.Println(err)
		}
		defer func() {
			err := rsp.Body.Close()
			if err != nil {
				log.Println(err)
			}
		}()
		if rsp.StatusCode == 200 {
			log.Println("配置文件重载成功")
		} else {
			log.Println("配置文件重载失败")
		}
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(reloadConfig)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
