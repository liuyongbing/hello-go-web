package global

import (
	ut "github.com/go-playground/universal-translator"

	"github.com/liuyongbing/hello-go-web/user-web/config"
	"github.com/liuyongbing/hello-go-web/user-web/proto"
)

var (
	// 多语言翻译器
	Trans ut.Translator

	// 服务器配置信息
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	// User GRPC 客户端连接
	UserSrvClient proto.UserClient
)
