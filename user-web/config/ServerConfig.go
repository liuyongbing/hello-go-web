package config

// User API 服务配置
type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	Host        string        `mapstructure:"host"`
	Port        int           `mapstructure:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"`
}
