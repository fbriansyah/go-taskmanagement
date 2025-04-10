package domain

type TaskStatus string

const (
	TaskIsInProgress = TaskStatus("INPROGRESS")
	TaskIsDone       = TaskStatus("DONE")
)

func (t TaskStatus) String() string {
	return string(t)
}

type Task struct {
	ID          string
	Title       string
	Description string
	Status      TaskStatus
}

func NewTask(id, title, description string) Task {
	return Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      TaskIsInProgress,
	}
}

func (t *Task) Done() *Task {
	t.Status = TaskIsDone
	return t
}

func (t *Task) InProgress() *Task {
	t.Status = TaskIsInProgress
	return t
}