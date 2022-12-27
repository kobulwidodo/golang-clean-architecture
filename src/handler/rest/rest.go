package rest

import (
	"context"
	"fmt"
	"go-clean/src/business/usecase"
	"go-clean/src/lib/auth"
	"go-clean/src/lib/configreader"
	"go-clean/src/lib/log"
	"go-clean/src/utils/config"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	infoRequest  string = `httpclient Sent Request: uri=%v method=%v`
	infoResponse string = `httpclient Received Response: uri=%v method=%v resp_code=%v`
)

var once = &sync.Once{}

type REST interface {
	Run()
}

type rest struct {
	http         *gin.Engine
	conf         config.GinConfig
	configreader configreader.Interface
	auth         auth.Interface
	log          log.Interface
	uc           *usecase.Usecase
}

func Init(conf config.GinConfig, confReader configreader.Interface, auth auth.Interface, log log.Interface, uc *usecase.Usecase) REST {
	r := &rest{}
	once.Do(func() {
		switch conf.Mode {
		case gin.ReleaseMode:
			gin.SetMode(gin.ReleaseMode)
		case gin.DebugMode, gin.TestMode:
			gin.SetMode(gin.TestMode)
		default:
			gin.SetMode("")
		}

		httpServ := gin.New()

		r = &rest{
			conf:         conf,
			configreader: confReader,
			log:          log,
			auth:         auth,
			http:         httpServ,
			uc:           uc,
		}

		switch r.conf.CORS.Mode {
		case "allowall":
			r.http.Use(cors.New(cors.Config{
				AllowAllOrigins: true,
				AllowHeaders:    []string{"*"},
				AllowMethods: []string{
					http.MethodHead,
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodPatch,
					http.MethodDelete,
				},
			}))
		default:
			r.http.Use(cors.New(cors.DefaultConfig()))
		}

		// Set Recovery
		r.http.Use(gin.Recovery())

		r.Register()
	})

	return r
}

func (r *rest) Run() {
	port := ":8080"
	if r.conf.Port != "" {
		port = fmt.Sprintf(":%s", r.conf.Port)
	}

	server := &http.Server{
		Addr:    port,
		Handler: r.http,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			r.log.Error(fmt.Sprintf("Serving HTTP error: %s", err.Error()))
		}
	}()
	r.log.Info(fmt.Sprintf("Listening and Serving HTTP on %s", server.Addr))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	r.log.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), r.conf.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		r.log.Fatal(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	r.log.Info("Server exiting")
}

func (r *rest) Register() {
	publicApi := r.http.Group("/public", r.BodyLogger)

	publicApi.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})
}
