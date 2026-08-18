package data

import (
	"errors"
	"testing"
	"time"

	"learn_go/lesson26_gorm_intro/model"
	"learn_go/lesson26_gorm_intro/repo"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestTaskRepo(t *testing.T) repo.TaskRepo {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite failed: %v", err)
	}
	if err := AutoMigrate(db); err != nil {
		t.Fatalf("auto migrate failed: %v", err)
	}
	return NewGormTaskRepo(db)
}

func TestGormTaskRepoCRUD(t *testing.T) {
	taskRepo := newTestTaskRepo(t)

	now := time.Now()
	task := &model.Task{Title: "GORM CRUD", Done: false, CreatedAt: now, UpdatedAt: now}
	if err := taskRepo.Create(t.Context(), task); err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if task.ID == 0 {
		t.Fatal("Create should write auto increment id back to task")
	}

	foundTask, err := taskRepo.GetByID(t.Context(), task.ID)
	if err != nil {
		t.Fatalf("GetByID failed: %v", err)
	}
	if foundTask.Title != "GORM CRUD" || foundTask.Done {
		t.Fatalf("GetByID returned unexpected task: %+v", foundTask)
	}

	finishedTask, err := taskRepo.MarkDone(t.Context(), task.ID)
	if err != nil {
		t.Fatalf("MarkDone failed: %v", err)
	}
	if !finishedTask.Done {
		t.Fatalf("MarkDone should set Done=true, got: %+v", finishedTask)
	}

	tasks, err := taskRepo.List(t.Context())
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if len(tasks) != 1 {
		t.Fatalf("List expected 1 task, got: %d", len(tasks))
	}
}

func TestGormTaskRepoNotFound(t *testing.T) {
	taskRepo := newTestTaskRepo(t)

	_, err := taskRepo.GetByID(t.Context(), 999)
	if !errors.Is(err, repo.ErrTaskNotFound) {
		t.Fatalf("GetByID expected ErrTaskNotFound, got: %v", err)
	}
}
