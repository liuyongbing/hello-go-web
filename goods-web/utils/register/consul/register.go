package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

/*
Register
服务注册
*/
func (r *Registry) Register(addr string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

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
	return nil
}

func (r *Registry) DeRegister(serviceId string) error {
	return nil
}
