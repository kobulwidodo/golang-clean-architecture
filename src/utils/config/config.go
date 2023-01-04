package config

import (
	"time"

	"gorm.io/driver/mysql"
)

type Application struct {
	Meta ApplicationMeta
	Gin  GinConfig
	SQL  mysql.Config
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
	Timeout         time.Duration
	ShutdownTimeout time.Duration
	CORS            CORSConfig
	Meta            ApplicationMeta
}

type CORSConfig struct {
	Mode string
}

func Init() Application {
	return Application{}
}
