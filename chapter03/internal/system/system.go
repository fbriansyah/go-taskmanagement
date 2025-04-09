package system

import (
	"context"
	"taksmanagement/internal/config"
	"taksmanagement/internal/waiter"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type System interface {
	Config() config.AppConfig
	Logger() zerolog.Logger
	Mux() *chi.Mux
	RPC() *grpc.Server
	Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, System) error
}