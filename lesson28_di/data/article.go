package data

import (
	"context"
	"sync"

	"learn_go/lesson28_di/model"
	"learn_go/lesson28_di/repo"
)

// memoryArticleRepo 是 ArticleRepo 的内存实现。
// 小写开头，外部不能直接创建，只能通过 NewMemoryArticleRepo 拿到接口。
type memoryArticleRepo struct {
	mu       sync.Mutex
	articles []model.Article
	nextID   int
}

// NewMemoryArticleRepo 创建内存版文章仓库。
// 返回的是接口类型 repo.ArticleRepo，隐藏具体实现。
func NewMemoryArticleRepo() repo.ArticleRepo {
	return &memoryArticleRepo{
		articles: make([]model.Article, 0),
		nextID:   1,
	}
}

func (r *memoryArticleRepo) Create(ctx context.Context, article *model.Article) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	article.ID = r.nextID
	r.nextID++
	r.articles = append(r.articles, *article)
	return nil
}

func (r *memoryArticleRepo) GetByID(ctx context.Context, id int) (*model.Article, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.articles {
		if r.articles[i].ID == id {
			// 返回副本，保护内部数据
			a := r.articles[i]
			return &a, nil
		}
	}
	return nil, nil // 没找到返回 nil，由 biz 层判断
}

func (r *memoryArticleRepo) List(ctx context.Context) ([]model.Article, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// 返回副本
	result := make([]model.Article, len(r.articles))
	copy(result, r.articles)
	return result, nil
}