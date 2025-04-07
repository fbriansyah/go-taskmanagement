package memorydb

import (
	"context"
	"taksmanagement/internal/domain"
	"taksmanagement/internal/model"
)

func (m *MemoryDB) FindByID(_ context.Context, id string) (*model.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if task, ok := taskData[id]; ok {
		return &task, nil
	}
	return nil, domain.ErrNoRecord
}

func (m *MemoryDB) GetAllTask(_ context.Context, tasks *[]model.Task) error {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if len(taskData) == 0 {
		return nil
	}

	if tasks != nil {
		for _, task := range taskData {
			*tasks = append(*tasks, task)
		}
	}

	return nil
}

func (m *MemoryDB) Insert(_ context.Context, task *model.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := taskData[task.ID]; ok {
		return domain.ErrDuplicateRecord
	}

	taskData[task.ID] = *task
	return nil
}

func (m *MemoryDB) Update(_ context.Context, task *model.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := taskData[task.ID]; !ok {
		return domain.ErrNoRecord
	}

	taskData[task.ID] = *task
	return nil
}