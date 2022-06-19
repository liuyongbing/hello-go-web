package config

// 配置中心 Nacos 配置项
type NacosConfig struct {
	Host        string `mapstructure:"host"`
	Port        uint64 `mapstructure:"port"`
	NamespaceId string `mapstructure:"namespace_id"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DataId      string `mapstructure:"data_id"`
	Group       string `mapstructure:"group"`
}
