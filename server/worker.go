package server

import (
	"fmt"
	"github.com/arioki1/qiscus-caa/config"
	"github.com/arioki1/qiscus-caa/helpers"
	"github.com/arioki1/qiscus-caa/src/worker"
	"github.com/hibiken/asynq"
	"os"
	"os/signal"
	"syscall"
)

func WorkerStart(cfg config.Config) {
	helpers.PrintInfoStringLog("starting worker..")
	go func() {
		initWorker(cfg)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	helpers.PrintInfoStringLog("shutting down worker...")
	helpers.PrintInfoStringLog("worker exiting")
}

func initWorker(cfg config.Config) {
	workerService := worker.NewWorkerServices(cfg)

	redisConnection, err := asynq.ParseRedisURI(cfg.GetRedisURL())
	if err != nil {
		helpers.PrintFatalStringLog(fmt.Sprintf("ParseRedisURI Error: %+v", err))
		return
	}
	workerSync := asynq.NewServer(redisConnection, asynq.Config{
		//Specify how many concurrent workers to use.
		Concurrency: 1,
	})

	mux := asynq.NewServeMux()

	mux.HandleFunc(cfg.GetAppName()+":assign_agent", workerService.HandleTaskAssignAgent)

	if err := workerSync.Run(mux); err != nil {
		helpers.PrintFatalStringLog(fmt.Sprintf("%+v", err))
		return
	}
}
