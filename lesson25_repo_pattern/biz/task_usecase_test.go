package biz

import (
	"context"
	"errors"
	"testing"

	"learn_go/lesson25_repo_pattern/model"
	"learn_go/lesson25_repo_pattern/repo"
)

type fakeTaskRepo struct {
	created bool
	tasks   []model.Task
}

func (r *fakeTaskRepo) Create(ctx context.Context, task model.Task) (model.Task, error) {
	r.created = true
	task.ID = 100
	r.tasks = append(r.tasks, task)
	return task, nil
}

func (r *fakeTaskRepo) GetByID(ctx context.Context, id int) (model.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return model.Task{}, repo.ErrTaskNotFound
}

func (r *fakeTaskRepo) List(ctx context.Context) ([]model.Task, error) {
	return r.tasks, nil
}

func (r *fakeTaskRepo) MarkDone(ctx context.Context, id int) (model.Task, error) {
	return model.Task{ID: id, Title: "done", Done: true}, nil
}

func TestCreateTask(t *testing.T) {
	fakeRepo := &fakeTaskRepo{}
	useCase := NewTaskUseCase(fakeRepo)

	task, err := useCase.CreateTask(t.Context(), model.CreateTaskRequest{Title: "  学 repo 模式  "})
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}
	if !fakeRepo.created {
		t.Fatal("expected usecase to call repo.Create")
	}
	if task.ID != 100 || task.Title != "学 repo 模式" {
		t.Fatalf("CreateTask returned unexpected task: %+v", task)
	}
}

func TestCreateTaskInvalidTitle(t *testing.T) {
	fakeRepo := &fakeTaskRepo{}
	useCase := NewTaskUseCase(fakeRepo)

	_, err := useCase.CreateTask(t.Context(), model.CreateTaskRequest{Title: "   "})
	if !errors.Is(err, ErrInvalidTaskTitle) {
		t.Fatalf("CreateTask expected ErrInvalidTaskTitle, got: %v", err)
	}
	if fakeRepo.created {
		t.Fatal("invalid title should not call repo.Create")
	}
}
