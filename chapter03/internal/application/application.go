package application

import (
	"context"
	"taksmanagement/internal/domain"
	"taksmanagement/internal/model"
)

type (
	CreateTask struct {
		ID          string
		Title       string
		Description string
	}

	TaskDone struct {
		ID string
	}

	TaskInProgress struct {
		ID string
	}

	App interface {
		AllTaks(ctx context.Context) ([]domain.Task,error)
		CreateTask(ctx context.Context, newTask CreateTask) error
		TaskDone(ctx context.Context, done TaskDone) error
		TaskInprogress(ctx context.Context, inProgress TaskInProgress) error
	}

	Application struct {
		tasks domain.TaskRepository
	}
)

var _ App = (*Application)(nil)

func (a Application) AllTaks(ctx context.Context) ([]domain.Task,error) {
	tasks := make([]model.Task, 0)
	allTask := make([]domain.Task, 0)

	if err := a.tasks.GetAllTask(ctx, &tasks); err != nil {
		return allTask, err
	}

	for _, task := range tasks {
		allTask = append(allTask, taskModelToDomain(&task))
	}

	return allTask, nil
}

func (a Application) CreateTask(ctx context.Context, newTask CreateTask) error {
	if newTask.Title == "" {
		return domain.ErrEmptyTitle
	}

	task := domain.NewTask(newTask.ID, newTask.Title, newTask.Description)
	taskModel := taskDomainToModel(&task)

	if err := a.tasks.Insert(ctx, &taskModel); err != nil {
		return err
	}

	return nil
}

func (a Application) TaskDone(ctx context.Context, done TaskDone) error {
	task, err := a.tasks.FindByID(ctx, done.ID) 
	if err != nil {
		return err
	}

	taskDomain := taskModelToDomain(task)
	taskDomain.Done()

	updatedTask := taskDomainToModel(&taskDomain)

	err = a.tasks.Update(ctx, &updatedTask)
	if err != nil {
		return err
	}

	return nil
}

func (a Application) TaskInprogress(ctx context.Context, inProgress TaskInProgress) error {
	task, err := a.tasks.FindByID(ctx, inProgress.ID) 
	if err != nil {
		return err
	}

	taskDomain := taskModelToDomain(task)
	taskDomain.InProgress()

	updatedTask := taskDomainToModel(&taskDomain)

	err = a.tasks.Update(ctx, &updatedTask)
	if err != nil {
		return err
	}

	return nil
}

