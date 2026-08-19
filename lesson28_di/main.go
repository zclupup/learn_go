package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"learn_go/lesson28_di/biz"
	"learn_go/lesson28_di/config"
	"learn_go/lesson28_di/data"
)

// ======== Lesson 28：依赖注入思想 —— 手写 NewXxx 组装对象 ========
//
// 前面我们学了分层（Lesson 23）、repo 模式（Lesson 25）、配置文件（Lesson 27）。
// 这些课里，main 已经在做"组装"这件事了，但每课只有一个两层结构。
//
// 真实项目里，对象嵌套可能很深：
//   config -> repo -> usecase -> handler -> server
//
// 每个对象都依赖别的对象。问题是：谁来创建这些依赖？怎么把它们串起来？
//
// 答案就是「依赖注入」（Dependency Injection，DI）：
//   每个对象通过构造函数声明"我需要什么"，
//   由外部（通常是 main）把依赖传进去。
//
// 本课用一个小型"文章管理"系统演示：
//   - config 层：读取 yaml 配置
//   - repo 层：定义数据接口
//   - data 层：实现接口（内存版）
//   - biz 层：业务逻辑，依赖 repo + 配置参数
//   - main：把所有东西组装在一起
//
// 对照 issue_api：
//   cmd/issue_api/main.go 里调用 wire 生成的 InitApp 函数做组装；
//   wire_gen.go 是自动生成的"组装代码"，本质和本课 main 做的事一样。

// ================================================================
// 第一部分：无依赖注入的写法（反例）—— 对象自己创建依赖
// ================================================================

// BadArticleUseCase 是一个反面教材：它在自己内部创建依赖。
// 问题：如果换一个 repo 实现（比如 MySQL），就要改 BadArticleUseCase 的代码。
type BadArticleUseCase struct {
	repo        *MemoryArticleRepoBad // 直接依赖具体类型！
	projectName string
}

// MemoryArticleRepoBad 为了演示反例，需要把 data 层的类型导出。
// 注意：这里特意把类型名大写导出，让 BadArticleUseCase 能直接依赖它。
// 实际项目中 data 层类型通常小写，就是为了防止这种"直接依赖具体实现"的写法。
type MemoryArticleRepoBad struct {
	articles []ArticleBad
	nextID   int
}

type ArticleBad struct {
	ID    int
	Title string
}

func NewMemoryArticleRepoBad() *MemoryArticleRepoBad {
	return &MemoryArticleRepoBad{articles: make([]ArticleBad, 0), nextID: 1}
}

func (r *MemoryArticleRepoBad) Create(title string) *ArticleBad {
	a := ArticleBad{ID: r.nextID, Title: title}
	r.nextID++
	r.articles = append(r.articles, a)
	return &a
}

// NewBadArticleUseCase 的依赖是"自己创建"的，而不是"从外面传进来"的。
// 这就是反面教材。
func NewBadArticleUseCase() *BadArticleUseCase {
	return &BadArticleUseCase{
		repo:        NewMemoryArticleRepoBad(), // 内部创建！换实现就得改这里
		projectName: "hardcoded",               // 硬编码！换环境就得改代码
	}
}

// ================================================================
// 第二部分：依赖注入的写法（正例）—— 在 main 中组装
// ================================================================

// consoleLogger 是 biz.Logger 的一个具体实现，用 fmt.Println 打印到控制台。
type consoleLogger struct{}

func (l *consoleLogger) Log(message string) {
	fmt.Println("[LOG]", message)
}

func main() {
	// ---------- step 1: 解析命令行参数 ----------
	confPath := flag.String("conf", "lesson28_di/config.yaml", "配置文件路径")
	flag.Parse()

	// ---------- step 2: 加载配置 ----------
	// 配置是最底层的依赖，其他对象不依赖它，但它的值会被注入到各层。
	cfg, err := config.Load(*confPath)
	if err != nil {
		log.Fatal("加载配置失败:", err)
	}
	fmt.Printf("配置加载成功：project.name = %s\n\n", cfg.Project.Name)

	// ---------- step 3: 创建底层依赖（数据仓库 + 日志组件） ----------
	// 先创建最底层的依赖：数据仓库、日志组件。
	articleRepo := data.NewMemoryArticleRepo()
	logger := &consoleLogger{} // 实现 biz.Logger 接口

	// ---------- step 4: 创建 biz 层（业务逻辑） ----------
	// 把 repo、配置参数、logger 三个依赖一起注入到 usecase 里。
	articleUseCase := biz.NewArticleUseCase(articleRepo, cfg.Project.Name, logger)

	// ---------- step 5: 执行业务操作 ----------
	// 后续如果加 HTTP handler 层，也是在这里创建：
	//   articleHandler := handler.NewArticleHandler(articleUseCase)
	// 然后注册路由、启动服务。
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 发布文章
	article1, err := articleUseCase.PublishArticle(ctx, "Go 依赖注入入门", "依赖注入让代码更灵活...", "zhangcl")
	if err != nil {
		log.Fatal("发布文章失败:", err)
	}
	fmt.Printf("发布成功：ID=%d Title=%s\n", article1.ID, article1.Title)

	article2, err := articleUseCase.PublishArticle(ctx, "理解 Wire 的前置知识", "Wire 是 Google 的依赖注入工具...", "zhangcl")
	if err != nil {
		log.Fatal("发布文章失败:", err)
	}
	fmt.Printf("发布成功：ID=%d Title=%s\n", article2.ID, article2.Title)

	// 获取单篇
	found, err := articleUseCase.GetArticle(ctx, 1)
	if err != nil {
		log.Fatal("获取文章失败:", err)
	}
	fmt.Printf("\n获取文章 1：%s（作者：%s）\n", found.Title, found.Author)

	// 列表
	allArticles, err := articleUseCase.ListArticles(ctx)
	if err != nil {
		log.Fatal("获取列表失败:", err)
	}
	fmt.Println("\n文章列表（标题带项目名前缀）：")
	for _, a := range allArticles {
		fmt.Printf("- %s\n", a.Title)
	}

	// ================================================================
	// 总结：依赖注入的核心思想
	// ================================================================
	//
	// 1. 每个对象通过构造函数声明"我需要什么"
	//    - NewArticleUseCase(repo, projectName)：我需要一个仓库和一个项目名
	//    - NewMemoryArticleRepo()：我不需要依赖（最底层）
	//
	// 2. main 是唯一的"组装中心"（Composition Root）
	//    - 所有 NewXxx 都在 main 里调用
	//    - 创建顺序 = 依赖顺序：先创建被依赖的，再创建依赖它的
	//
	// 3. 好处：
	//    - 换实现：想把内存换成 MySQL？只改 main 里 NewMemoryArticleRepo -> NewMySQLArticleRepo
	//    - 好测试：测试时传假的 repo，不用连数据库
	//    - 依赖清晰：看构造函数签名就知道每个对象依赖什么
	//
	// 4. 对比反例：
	//    BadArticleUseCase 在自己内部 new 依赖，换实现要改业务代码，
	//    配置硬编码在代码里，改环境要改代码。这就是"没有依赖注入"的问题。
	//
	// 下一课（Lesson 29）学 Wire：
	//    Wire 就是帮我们自动生成 main 里这些 NewXxx 组装代码的工具。
	//    本课手动写的这些组装代码，就是 Wire 的输出结果。
}