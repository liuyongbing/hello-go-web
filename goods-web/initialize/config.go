package initialize

import (
	"encoding/json"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/liuyongbing/hello-go-web/goods-web/config"
	"github.com/liuyongbing/hello-go-web/goods-web/global"
	"github.com/liuyongbing/hello-go-web/goods-web/utils"
)

/*
GetEnvInfo
获取环境变量
*/
func GetEnvInfo(env string) string {
	viper.AutomaticEnv()

	return viper.GetString(env)
}

/*
LoadConfig
从配置中心加载配置信息
*/
func LoadConfigFromNacos(cfg config.NacosConfig) {
	// Nacos 服务器配置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: cfg.Host,
			Port:   cfg.Port,
		},
	}

	// Nacos 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         cfg.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建 Nacos 客户端连接
	nacosClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}

	// 读取 Nacos 配置信息
	configInfo, err := nacosClient.GetConfig(vo.ConfigParam{
		DataId: cfg.DataId,
		Group:  cfg.Group,
	})
	if err != nil {
		panic(err)
	}
	zap.S().Infof("Nacos 原始配置内容：%s", configInfo)

	serverConfig := global.ServerConfig
	if err := json.Unmarshal([]byte(configInfo), &serverConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("Nacos 绑定配置内容：%v", serverConfig)
}

/*
InitConfig
初始化配置信息
*/
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

	// 读取本地配置
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	baseCfg := config.BaseConfig{}
	if err := v.Unmarshal(&baseCfg); err != nil {
		panic(err)
	}
	zap.S().Infof("本地配置：%v", baseCfg)

	nacosCfg := baseCfg.NacosInfo
	LoadConfigFromNacos(nacosCfg)

}

/*
InitConfig
初始化配置信息
*/
func InitConfig2() {
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

	// 根据环境，动态生成端口
	if configFileMode != "dev" {
		freePort, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = freePort
			zap.S().Infof("重置端口[%d]", global.ServerConfig.Port)
		}
	}

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
