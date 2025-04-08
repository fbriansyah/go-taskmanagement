package domain

import (
	"context"
	"taksmanagement/internal/model"
)

type TaskRepository interface {
	FindByID(ctx context.Context, id string) (*model.Task, error)
	GetAllTask(ctx context.Context, tasks *[]model.Task) error
	Insert(ctx context.Context, task *model.Task) error
	Update(ctx context.Context, task *model.Task) error
}