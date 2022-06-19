package initialize

import (
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"github.com/liuyongbing/hello-go-web/goods-web/global"
	"github.com/liuyongbing/hello-go-web/goods-web/proto"
)

/*
InitSrvConn
GRPC 服务连接初始化
*/
func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	clientConn, err := grpc.Dial(
		// consul://[user:password@]127.0.0.127:8555/my-service?[healthy=]&[wait=]&[near=]&[insecure=]&[limit=]&[tag=]&[token=]
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		logStr := fmt.Sprintf("[InitSrvConn] 无法连接服务 [%s]", global.ServerConfig.GoodsSrvInfo.Name)
		zap.S().Fatal(logStr, "msg", err.Error())
	}
	// 生成 grpc 的 client 并调用接口
	global.GoodsSrvClient = proto.NewGoodsClient(clientConn)
}
