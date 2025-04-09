package main

import (
	"fmt"
	"net/http"
	"os"
	taskmanagement "taksmanagement"
	"taksmanagement/internal/config"
	"taksmanagement/internal/logger"
	"taksmanagement/internal/system"
	"taksmanagement/internal/waiter"
	"taksmanagement/internal/web"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	var cfg config.AppConfig
	// parse config/env/...
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}

	m := app{cfg: cfg}
		m.logger = logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})
	m.rpc = initRpc(cfg.Rpc)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	m.modules = []system.Module{
		&taskmanagement.Module{},
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	// Mount general web resources
	m.mux.Mount("/", http.FileServer(http.FS(web.WebUI)))

	fmt.Println("started tasks management application")
	defer fmt.Println("stopped tasks management application")

	m.waiter.Add(
		m.waitForWeb,
		m.waitForRPC,
	)

	return m.waiter.Wait()
}

func initRpc(_ config.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}
