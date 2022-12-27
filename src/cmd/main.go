package main

import (
	"go-clean/src/business/domain"
	"go-clean/src/business/usecase"
	"go-clean/src/handler/rest"
	"go-clean/src/lib/auth"
	"go-clean/src/lib/configreader"
	"go-clean/src/lib/log"
	sql "go-clean/src/lib/postgresql"
	"go-clean/src/utils/config"
)

const (
	configFile string = "./etc/cfg/config.json"
)

func main() {
	cfg := config.Init()
	configReader := configreader.Init(configreader.Options{
		ConfigFile: configFile,
	})
	configReader.ReadConfig(&cfg)

	log := log.Init()

	auth := auth.Init(cfg.Auth, log)

	db := sql.Init(cfg.SQL, log)

	d := domain.Init(log, db, auth)

	uc := usecase.Init(log, auth, d)

	r := rest.Init(cfg.Gin, configReader, auth, log, uc)

	r.Run()
}
