package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/liuyongbing/hello-go-web/user-web/config"
)

var (
	Trans        ut.Translator
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)
