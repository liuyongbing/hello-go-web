package config

// User API 服务配置
type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	Host        string        `mapstructure:"host"`
	Port        int           `mapstructure:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
}

// Grpc: 用户微服务配置
type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}
