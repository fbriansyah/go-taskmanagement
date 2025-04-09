package taskmanagement

import (
	"context"
	"taksmanagement/internal/application"
	"taksmanagement/internal/grpc"
	"taksmanagement/internal/logging"
	memorydb "taksmanagement/internal/memoryDB"
	"taksmanagement/internal/rest"
	"taksmanagement/internal/system"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, sys system.System) error {
	tasks := memorydb.New()

	var app application.App
	app = application.New(tasks)
	app = logging.LogApplicationAccess(app, sys.Logger())

	if err := grpc.RegisterServer(app, sys.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, sys.Mux(), sys.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(sys.Mux()); err != nil {
		return err
	}
	return nil
}