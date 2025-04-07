package memorydb

import (
	"sync"
	"taksmanagement/internal/domain"
	"taksmanagement/internal/model"
)

var (
	once     sync.Once
	taskData map[string]model.Task
)

type MemoryDB struct {
	mu sync.RWMutex
}


func New() *MemoryDB {
	once.Do(func() {
		taskData = make(map[string]model.Task)
	})
	return &MemoryDB{}
}

var _ domain.TaskRepository = (*MemoryDB)(nil)
