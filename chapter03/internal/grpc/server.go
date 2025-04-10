package grpc

import (
	"context"
	"taksmanagement/internal/application"
	"taksmanagement/taskpb"

	"github.com/google/uuid"
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

func (s server) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	id := uuid.New().String()

	task, err := s.app.CreateTask(ctx, application.CreateTask{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:          id,
			Title:       req.Title,
			Description: req.Description,
			Status:      task.Status.String(),
		},
	}, nil
}

// TaskDone implements taskpb.TaskServiceServer.
func (s server) TaskDone(ctx context.Context, req *taskpb.TaskDoneRequest) (*taskpb.TaskDoneResponse, error) {
	task, err := s.app.TaskDone(ctx, application.TaskDone{
		ID: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &taskpb.TaskDoneResponse{
		Task: &taskpb.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status.String(),
		},
	}, nil
}

// TaskInProgress implements taskpb.TaskServiceServer.
func (s server) TaskInProgress(ctx context.Context, req *taskpb.TaskInProgressRequest) (*taskpb.TaskInProgressResponse, error) {
	task, err := s.app.TaskInprogress(ctx, application.TaskInProgress{
		ID: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &taskpb.TaskInProgressResponse{
		Task: &taskpb.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status.String(),
		},
	}, nil
}
