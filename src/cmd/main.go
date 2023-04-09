package main

import (
	"go-clean/src/business/domain"
	"go-clean/src/business/usecase"
	"go-clean/src/handler/rest"
	"go-clean/src/lib/auth"
	"go-clean/src/lib/configreader"
	"go-clean/src/lib/sql"
	"go-clean/src/utils/config"

	_ "go-clean/docs/swagger"
)

// @contact.name   Rakhmad Giffari Nurfadhilah
// @contact.url    https://fadhilmail.tech/
// @contact.email  rakhmadgiffari14@gmail.com

// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization

const (
	configFile string = "./etc/cfg/config.json"
)

func main() {
	cfg := config.Init()
	configReader := configreader.Init(configreader.Options{
		ConfigFile: configFile,
	})
	configReader.ReadConfig(&cfg)

	auth := auth.Init()

	db := sql.Init(cfg.SQL)

	d := domain.Init(db)

	uc := usecase.Init(auth, d)

	r := rest.Init(cfg.Gin, configReader, uc, auth)

	r.Run()
}
