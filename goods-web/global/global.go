package global

import (
	ut "github.com/go-playground/universal-translator"

	"github.com/liuyongbing/hello-go-web/goods-web/config"
	"github.com/liuyongbing/hello-go-web/goods-web/proto"
)

var (
	// 多语言翻译器
	Trans ut.Translator

	// 服务器配置信息
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	// GRPC Client
	GoodsSrvClient proto.GoodsClient
)
