package initinalize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/liuyongbing/hello-go-web/user-web/global"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()

	return viper.GetString(env)
}

func InitConfig() {
	// 根据环境变量加载配置文件
	debug := GetEnvInfo("PATH")
	pathStr := ""
	configFileMode := "prd"
	if pathStr != debug {
		configFileMode = "dev"
	}
	configFileName := fmt.Sprintf("config/config-%s.yaml", configFileMode)

	v := viper.New()
	v.SetConfigFile(configFileName)

	// 读取配置文件内容
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 将配置文件内容映射到 配置 struct
	serverConfig := global.ServerConfig
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("加载配置信息：%v", serverConfig)

	// 动态监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化：%v", e.Name)
		v.ReadInConfig()
		v.Unmarshal(&serverConfig)
		zap.S().Infof("配置信息：%v", serverConfig)
	})

	// time.Sleep(time.Second * 100)

}
