package initialize

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/liuyongbing/hello-go-web/user-web/global"
	"github.com/liuyongbing/hello-go-web/user-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

/*
InitSrvConn
GRPC 服务连接初始化
*/
func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	cfg := api.DefaultConfig()
	// cfg.Address = "127.0.0.1:8500"
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// filter := `Service == "user-web"`
	filter := fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name)
	data, err := client.Agent().ServicesWithFilter(filter)
	if err != nil {
		panic(err)
	}

	host := ""
	port := 0

	for _, value := range data {
		host = value.Address
		port = value.Port
		break
	}
	logStr := ""
	if host == "" {
		logStr = fmt.Sprintf("[InitSrvConn] 无法发现服务 [%s]", global.ServerConfig.UserSrvInfo.Name)
		zap.S().Fatal(logStr)
		return
	}

	// 拨号连接 user grpc 服务
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		logStr = fmt.Sprintf("[InitSrvConn] 无法连接服务 [%s]", global.ServerConfig.UserSrvInfo.Name)
		zap.S().Errorw(logStr, "msg", err.Error())
	}
	// 生成 grpc 的 client 并调用接口
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
