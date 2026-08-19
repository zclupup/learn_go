package repo

import (
	"context"

	"learn_go/lesson28_di/model"
)

// ArticleRepo 定义上层需要的数据能力。
// biz 层只依赖这个接口，不关心底层是内存、MySQL 还是外部 API。
type ArticleRepo interface {
	Create(ctx context.Context, article *model.Article) error
	GetByID(ctx context.Context, id int) (*model.Article, error)
	List(ctx context.Context) ([]model.Article, error)
}