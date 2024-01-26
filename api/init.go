package api

import (
	"fmt"
	"jade_backend/api/route"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Engine *gin.Engine

type WebServerCfg struct {
	Port         int `mapstructure:"Port"`
	WriteTimeout int `mapstructure:"WriteTimeout"`
	ReadTimeout  int `mapstructure:"ReadTimeout"`
}

func StartServer() error {
	var cfg WebServerCfg
	if err := viper.Sub("WebServer").UnmarshalExact(&cfg); err != nil {
		return err
	}

	Engine = gin.Default()

	// 使用cors中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000","http://localhost:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true

	Engine.Use(cors.New(config))

	route.SetupJadeController(Engine.Group("/api"))

	//	route.SetupRouter(e.Group("/api"))

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        Engine,
		ReadTimeout:    time.Second * time.Duration(cfg.ReadTimeout),
		WriteTimeout:   time.Second * time.Duration(cfg.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}
