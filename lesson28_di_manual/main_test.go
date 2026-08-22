package main

import (
	"context"
	"errors"
	"testing"
)

func TestTaskServiceCreateTaskWithFakeRepo(t *testing.T) {
	repoImpl := &fakeTaskRepo{tasks: make(map[int]*Task)}
	logger := NewSilentLogger()
	service := NewTaskService(repoImpl, logger)

	task, err := service.CreateTask(context.Background(), "学习 fake repo 测试")
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}
	if task.ID != 1 {
		t.Fatalf("expected ID=1, got=%d", task.ID)
	}
	if !repoImpl.createCalled {
		t.Fatal("expected fake repo Create to be called")
	}
}

func TestTaskServiceCreateTaskInvalidTitle(t *testing.T) {
	repoImpl := &fakeTaskRepo{tasks: make(map[int]*Task)}
	logger := NewSilentLogger()
	service := NewTaskService(repoImpl, logger)

	_, err := service.CreateTask(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty title, got nil")
	}
	if repoImpl.createCalled {
		t.Fatal("repo Create should not be called when title is empty")
	}
}

func TestTaskServiceGetTaskNotFound(t *testing.T) {
	repoImpl := &fakeTaskRepo{tasks: make(map[int]*Task)}
	logger := NewSilentLogger()
	service := NewTaskService(repoImpl, logger)

	_, err := service.GetTask(context.Background(), 999)
	if !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("expected ErrTaskNotFound, got: %v", err)
	}
}

func TestNewAppWithDepsUsesInjectedRepo(t *testing.T) {
	repoImpl := &fakeTaskRepo{tasks: make(map[int]*Task)}
	logger := NewSilentLogger()
	app := NewAppWithDeps(repoImpl, logger)

	_, err := app.TaskService.CreateTask(context.Background(), "通过注入 repo 创建")
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}
	if !repoImpl.createCalled {
		t.Fatal("expected injected fake repo to be used")
	}
}
