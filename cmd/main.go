package main

import (
	"github.com/arioki1/qiscus-caa/config"
	"github.com/arioki1/qiscus-caa/router"
	"github.com/arioki1/qiscus-caa/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("failed to run server: %s", err.Error())
	}

	if !cfg.GetDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	go func() {
		server.WorkerStart(cfg)
	}()

	routes := router.NewRouter(cfg)
	server.ApiStart(cfg, routes)

}
