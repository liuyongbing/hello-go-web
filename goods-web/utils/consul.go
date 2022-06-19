package utils

import (
	"fmt"

	"github.com/hashicorp/consul/api"

	"github.com/liuyongbing/hello-go-web/goods-web/global"
)

/*
Register
服务注册
*/
func Register(addr string, port int, name string, tags []string, id string) {
	consulInfo := global.ServerConfig.ConsulInfo

	cfg := api.DefaultConfig()
	// cfg.Address = "127.0.0.1:8500"
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 服务健康检查对象
	check := api.AgentServiceCheck{
		// GRPC:                           fmt.Sprintf("%s:%d", addr, port),
		HTTP:                           fmt.Sprintf("http://%s:%d/health", addr, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 服务注册对象
	registration := api.AgentServiceRegistration{
		Name:    name,
		ID:      id,
		Tags:    tags,
		Port:    port,
		Address: addr,
		Check:   &check,
	}

	// 注册服务
	err = client.Agent().ServiceRegister(&registration)
	if err != nil {
		panic(err)
	}
}
