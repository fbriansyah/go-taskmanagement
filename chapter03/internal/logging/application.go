package logging

import (
	"context"
	"taksmanagement/internal/application"
	"taksmanagement/internal/domain"

	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger zerolog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}
var _ application.App = (*Application)(nil)

func (a Application) AllTaks(ctx context.Context) (tasks []domain.Task, err error) {
	a.logger.Info().Msg("--> Tasks.Alltasks")
	defer func() { a.logger.Info().Err(err).Msg("<-- Tasks.Alltasks") }()
	return a.App.AllTaks(ctx)
}

func (a Application) CreateTask(ctx context.Context, newTask application.CreateTask) (err error) {
	a.logger.Info().Msg("--> Tasks.CreateTask")
	defer func() { a.logger.Info().Err(err).Msg("<-- Tasks.CreateTask") }()
	return a.App.CreateTask(ctx, newTask)
}

func (a Application) TaskDone(ctx context.Context, done application.TaskDone) (err error) {
	a.logger.Info().Msg("--> Tasks.TaskDone")
	defer func() { a.logger.Info().Err(err).Msg("<-- Tasks.TaskDone") }()
	return a.App.TaskDone(ctx, done)
}

func (a Application) TaskInprogress(ctx context.Context, inProgress application.TaskInProgress) (err error) {
	a.logger.Info().Msg("--> Tasks.TaskInprogress")
	defer func() { a.logger.Info().Err(err).Msg("<-- Tasks.TaskInprogress") }()
	return a.App.TaskInprogress(ctx, inProgress)
}
