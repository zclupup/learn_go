package biz

import (
	"context"
	"errors"
	"strings"
	"time"

	"learn_go/lesson26_gorm_intro/model"
	"learn_go/lesson26_gorm_intro/repo"
)

var ErrInvalidTaskTitle = errors.New("任务标题不能为空")

type TaskUseCase struct {
	taskRepo repo.TaskRepo
}

func NewTaskUseCase(taskRepo repo.TaskRepo) *TaskUseCase {
	return &TaskUseCase{taskRepo: taskRepo}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context, title string) (*model.Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, ErrInvalidTaskTitle
	}

	now := time.Now()
	task := &model.Task{
		Title:     title,
		Done:      false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := uc.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (uc *TaskUseCase) ListTasks(ctx context.Context) ([]model.Task, error) {
	return uc.taskRepo.List(ctx)
}

func (uc *TaskUseCase) FinishTask(ctx context.Context, id int) (*model.Task, error) {
	return uc.taskRepo.MarkDone(ctx, id)
}
