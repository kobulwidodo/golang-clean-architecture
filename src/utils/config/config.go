package config

import (
	"go-clean/src/lib/auth"
	sql "go-clean/src/lib/postgresql"
	"time"
)

type Application struct {
	Meta ApplicationMeta
	Gin  GinConfig
	Auth auth.Config
	SQL  sql.Config
}

type ApplicationMeta struct {
	Title       string
	Description string
	Host        string
	BasePath    string
	Version     string
}

type GinConfig struct {
	Port            string
	Mode            string
	LogRequest      bool
	LogResponse     bool
	Timeout         time.Duration
	ShutdownTimeout time.Duration
	CORS            CORSConfig
	Meta            ApplicationMeta
	Swagger         SwaggerConfig
}

type CORSConfig struct {
	Mode string
}

type SwaggerConfig struct {
	Enabled   bool
	Path      string
	BasicAuth BasicAuthConf
}

type BasicAuthConf struct {
	Username string
	Password string
}

func Init() Application {
	return Application{}
}
