package repo

import (
	"context"
	"errors"

	"learn_go/lesson25_repo_pattern/model"
)

var ErrTaskNotFound = errors.New("任务不存在")

// TaskRepo 定义 biz 层需要的数据访问能力。
// biz 只依赖这个接口，不关心数据来自内存、MySQL、Redis 还是外部接口。
type TaskRepo interface {
	Create(ctx context.Context, task model.Task) (model.Task, error)
	GetByID(ctx context.Context, id int) (model.Task, error)
	List(ctx context.Context) ([]model.Task, error)
	MarkDone(ctx context.Context, id int) (model.Task, error)
}
