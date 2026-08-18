package repo

import (
	"context"
	"errors"

	"learn_go/lesson26_gorm_intro/model"
)

var ErrTaskNotFound = errors.New("任务不存在")

type TaskRepo interface {
	Create(ctx context.Context, task *model.Task) error
	GetByID(ctx context.Context, id int) (*model.Task, error)
	List(ctx context.Context) ([]model.Task, error)
	MarkDone(ctx context.Context, id int) (*model.Task, error)
}
