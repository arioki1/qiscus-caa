package server

import (
	"context"
	"fmt"
	"github.com/arioki1/qiscus-caa/config"
	"github.com/arioki1/qiscus-caa/helpers"
	"github.com/arioki1/qiscus-caa/registry"
	"github.com/arioki1/qiscus-caa/src/api/v1/delivery"
	"github.com/hibiken/asynq"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ApiStart(cfg config.Config, router *gin.Engine) {
	helpers.PrintInfoStringLog("starting server..")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.GetPort()),
		Handler: router,
	}
	// Create a new Redis connection for the client.
	redisConnection, err := asynq.ParseRedisURI(cfg.GetRedisURL())
	if err != nil {
		helpers.PrintFatalStringLog(fmt.Sprintf("ParseRedisURI Error: %+v", err))
		return
	}

	// Create a new Asynq client.
	client := asynq.NewClient(redisConnection)
	defer client.Close()

	go func() {
		initServer(cfg, router, client)
	}()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("server forced to shutdown %s", err.Error())
	}

	log.Info().Msg("server exiting")
}

func initServer(cfg config.Config, router *gin.Engine, clientAsynq *asynq.Client) {
	rep := registry.NewRepositoryRegistry(cfg)
	uc := registry.NewUseCaseRegistry(rep, cfg)

	//Qiscus Webhook
	qwDelivery := delivery.NewQiscusWebhookDelivery(uc.QiscusWebhook(), clientAsynq)
	qwGroup := router.Group("/api/v1/webhook")
	qwDelivery.Mount(qwGroup)

}
