package config

// 服务启动基础配置
type BaseConfig struct {
	// 配置中心 Nacos 配置信息
	NacosInfo NacosConfig `mapstructure:"nacos"`
}
