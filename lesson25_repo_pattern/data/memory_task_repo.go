package data

import (
	"context"
	"sync"
	"time"

	"learn_go/lesson25_repo_pattern/model"
	"learn_go/lesson25_repo_pattern/repo"
)

type memoryTaskRepo struct {
	mu     sync.Mutex
	tasks  []model.Task
	nextID int
}

func NewMemoryTaskRepo() repo.TaskRepo {
	return &memoryTaskRepo{
		tasks: []model.Task{
			{ID: 1, Title: "阅读 issue_api/internal/server/gin.go", Done: false, CreatedAt: time.Now()},
			{ID: 2, Title: "找到 service -> biz -> repo 链路", Done: false, CreatedAt: time.Now()},
		},
		nextID: 3,
	}
}

func (r *memoryTaskRepo) Create(ctx context.Context, task model.Task) (model.Task, error) {
	if err := ctx.Err(); err != nil {
		return model.Task{}, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = r.nextID
	task.CreatedAt = time.Now()
	r.nextID++
	r.tasks = append(r.tasks, task)
	return task, nil
}

func (r *memoryTaskRepo) GetByID(ctx context.Context, id int) (model.Task, error) {
	if err := ctx.Err(); err != nil {
		return model.Task{}, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return model.Task{}, repo.ErrTaskNotFound
}

func (r *memoryTaskRepo) List(ctx context.Context) ([]model.Task, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]model.Task, len(r.tasks))
	copy(result, r.tasks)
	return result, nil
}

func (r *memoryTaskRepo) MarkDone(ctx context.Context, id int) (model.Task, error) {
	if err := ctx.Err(); err != nil {
		return model.Task{}, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.tasks {
		if r.tasks[i].ID == id {
			r.tasks[i].Done = true
			return r.tasks[i], nil
		}
	}
	return model.Task{}, repo.ErrTaskNotFound
}
