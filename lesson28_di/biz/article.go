package biz

import (
	"context"
	"errors"
	"fmt"
	"time"

	"learn_go/lesson28_di/model"
	"learn_go/lesson28_di/repo"
)

// Logger 日志接口，定义日志能力。
// 这是业务层声明的抽象，ArticleUseCase 只依赖这个接口，不关心具体实现。
type Logger interface {
	Log(message string)
}

var (
	ErrTitleEmpty      = errors.New("文章标题不能为空")
	ErrArticleNotFound = errors.New("文章不存在")
)

// ArticleUseCase 文章业务逻辑。
// 它依赖三个东西：ArticleRepo（数据能力）、projectName（配置参数）、Logger（日志能力）。
// 这三个依赖都通过构造函数 NewArticleUseCase 从外部传入——
// 这就是依赖注入。
type ArticleUseCase struct {
	repo        repo.ArticleRepo
	projectName string // 从配置来的，不是 repo 的职责
	logger      Logger // 新增日志依赖
}

// NewArticleUseCase 是 ArticleUseCase 的构造函数。
// 参数就是它需要的依赖：一个数据仓库 + 一个项目名称 + 一个 Logger。
// 调用方（main）负责把这三个东西传进来。
func NewArticleUseCase(repo repo.ArticleRepo, projectName string, logger Logger) *ArticleUseCase {
	return &ArticleUseCase{
		repo:        repo,
		projectName: projectName,
		logger:      logger,
	}
}

// PublishArticle 发布文章。
func (uc *ArticleUseCase) PublishArticle(ctx context.Context, title, content, author string) (*model.Article, error) {
	if title == "" {
		return nil, ErrTitleEmpty
	}

	uc.logger.Log("发布文章: " + title) // 调用注入进来的日志组件

	article := &model.Article{
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
	}

	if err := uc.repo.Create(ctx, article); err != nil {
		return nil, fmt.Errorf("保存文章失败: %w", err)
	}

	return article, nil
}

// GetArticle 获取单篇文章。
func (uc *ArticleUseCase) GetArticle(ctx context.Context, id int) (*model.Article, error) {
	article, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, ErrArticleNotFound
	}
	return article, nil
}

// ListArticles 列出所有文章，带上项目名作为前缀。
func (uc *ArticleUseCase) ListArticles(ctx context.Context) ([]model.Article, error) {
	articles, err := uc.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	// 用配置里的项目名修改标题前缀（演示配置参数的使用）
	for i := range articles {
		articles[i].Title = fmt.Sprintf("[%s] %s", uc.projectName, articles[i].Title)
	}

	return articles, nil
}
