package grpc

import (
	"context"
	"taksmanagement/internal/application"
	"taksmanagement/taskpb"

	"google.golang.org/grpc"
)

type server struct {
	app application.App
	taskpb.UnimplementedTaskServiceServer
}

var _ taskpb.TaskServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	taskpb.RegisterTaskServiceServer(registrar, server{app: app})
	return nil
}

func (s server) GetAllTask(ctx context.Context, _ *taskpb.GetAllTaskRequest) (*taskpb.GetAllTaskResponse, error) {
	result := &taskpb.GetAllTaskResponse{}

	tasks, err := s.app.AllTaks(ctx)
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		result.Tasks = append(result.Tasks, &taskpb.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status.String(),
		})
	}

	return result, nil
}