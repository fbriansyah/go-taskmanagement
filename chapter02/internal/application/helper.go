package application

import (
	"taksmanagement/internal/domain"
	"taksmanagement/internal/model"
)

func taskModelToDomain(task *model.Task) domain.Task {
	if task == nil {
		return domain.Task{}
	}
	
	return domain.Task{
		ID: task.ID,
		Title: task.Title,
		Description: task.Description,
		Status: domain.TaskStatus(task.Status),
	} 
}

func taskDomainToModel(task *domain.Task) model.Task {
	if task == nil {
		return model.Task{}
	}

	return model.Task{
		ID: task.ID,
		Title: task.Title,
		Description: task.Description,
		Status: task.Status.String(),
	}
}