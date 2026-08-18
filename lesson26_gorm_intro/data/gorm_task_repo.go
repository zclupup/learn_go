package data

import (
	"context"
	"errors"

	"learn_go/lesson26_gorm_intro/model"
	"learn_go/lesson26_gorm_intro/repo"

	"gorm.io/gorm"
)

type gormTaskRepo struct {
	db *gorm.DB
}

func NewGormTaskRepo(db *gorm.DB) repo.TaskRepo {
	return &gormTaskRepo{db: db}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Task{})
}

func (r *gormTaskRepo) Create(ctx context.Context, task *model.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *gormTaskRepo) GetByID(ctx context.Context, id int) (*model.Task, error) {
	var task model.Task
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repo.ErrTaskNotFound
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *gormTaskRepo) List(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.WithContext(ctx).Order("id ASC").Find(&tasks).Error
	return tasks, err
}

func (r *gormTaskRepo) MarkDone(ctx context.Context, id int) (*model.Task, error) {
	task, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).Model(task).Update("done", true).Error
	if err != nil {
		return nil, err
	}
	task.Done = true
	return task, nil
}
