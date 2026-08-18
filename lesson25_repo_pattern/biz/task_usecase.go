package biz

import (
	"context"
	"errors"
	"strings"

	"learn_go/lesson25_repo_pattern/model"
	"learn_go/lesson25_repo_pattern/repo"
)

var ErrInvalidTaskTitle = errors.New("任务标题不能为空")

type TaskUseCase struct {
	taskRepo repo.TaskRepo
}

func NewTaskUseCase(taskRepo repo.TaskRepo) *TaskUseCase {
	return &TaskUseCase{taskRepo: taskRepo}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context, req model.CreateTaskRequest) (model.Task, error) {
	title := strings.TrimSpace(req.Title)
	if title == "" {
		return model.Task{}, ErrInvalidTaskTitle
	}

	task := model.Task{Title: title, Done: false}
	return uc.taskRepo.Create(ctx, task)
}

func (uc *TaskUseCase) ListTasks(ctx context.Context) ([]model.Task, error) {
	return uc.taskRepo.List(ctx)
}

func (uc *TaskUseCase) FinishTask(ctx context.Context, id int) (model.Task, error) {
	return uc.taskRepo.MarkDone(ctx, id)
}
