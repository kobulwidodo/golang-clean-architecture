package main

import (
	"go-clean/src/business/domain"
	"go-clean/src/business/usecase"
	"go-clean/src/handler/rest"
	"go-clean/src/lib/auth"
	"go-clean/src/lib/configreader"
	"go-clean/src/lib/sql"
	"go-clean/src/utils/config"
	"log"

	_ "go-clean/docs/swagger"

	"github.com/spf13/cobra"
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

	rootCmd := &cobra.Command{Use: "app"}

	restCmd := &cobra.Command{
		Use:   "rest",
		Short: "Run the REST API Server",
		Run: func(cmd *cobra.Command, args []string) {
			r := rest.Init(cfg.Gin, uc, auth)
			r.Run()
		},
	}

	rootCmd.AddCommand(restCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error executing command : %v", err)
	}
}
