package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

// ============================================================
// Lesson 28 — 依赖注入思想：手写 NewXxx 组装对象
// ============================================================
//
// 本课从一个"任务管理"小例子出发，演示：
// 1. 什么是依赖注入（DI）
// 2. 为什么需要 DI（对比反例）
// 3. 手写 NewXxx 组装对象的模式
// 4. 组合根（Composition Root）的概念
// 5. 和 issue_api 中 Wire 的关系

// ============ 1. 模型定义 ============

// Task 是任务模型，纯数据结构，不依赖任何东西
type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

// ============ 2. Repo 接口（业务层依赖的"合同"）============

// TaskRepo 定义"任务仓库需要什么能力"
// 业务层只依赖这个接口，不关心底层是内存、MySQL 还是 Redis
type TaskRepo interface {
	Create(ctx context.Context, task *Task) error
	GetByID(ctx context.Context, id int) (*Task, error)
	List(ctx context.Context) ([]Task, error)
}

var ErrTaskNotFound = errors.New("task not found")

// ============ 3. Repo 实现：内存版 ============

// memoryTaskRepo 是小写开头，对外不可见，只能通过 NewMemoryTaskRepo() 创建
type memoryTaskRepo struct {
	tasks  map[int]*Task
	nextID int
}

// NewMemoryTaskRepo 返回 TaskRepo 接口（隐藏具体实现）
func NewMemoryTaskRepo() TaskRepo {
	return &memoryTaskRepo{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

func (r *memoryTaskRepo) Create(ctx context.Context, task *Task) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	task.ID = r.nextID
	r.nextID++
	task.CreatedAt = time.Now()
	r.tasks[task.ID] = task
	return nil
}

func (r *memoryTaskRepo) GetByID(ctx context.Context, id int) (*Task, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	t, ok := r.tasks[id]
	if !ok {
		return nil, ErrTaskNotFound
	}
	return t, nil
}

func (r *memoryTaskRepo) List(ctx context.Context) ([]Task, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	result := make([]Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		result = append(result, *t)
	}
	return result, nil
}

// ============ 4. 额外依赖：Logger 接口 ============

// Logger 定义日志能力
type Logger interface {
	Info(msg string)
}

// consoleLogger 控制台日志实现
type consoleLogger struct{}

func NewConsoleLogger() Logger {
	return &consoleLogger{}
}

func (l *consoleLogger) Info(msg string) {
	log.Println("[INFO]", msg)
}

// ============ 5. 业务层：TaskService（依赖 TaskRepo + Logger）============

// TaskService 是业务逻辑层
// 注意：它依赖的是接口 TaskRepo 和 Logger，不是具体实现
type TaskService struct {
	repo   TaskRepo
	logger Logger
}

// NewTaskService 通过构造函数注入所有依赖
// 这就是"依赖注入"的核心：需要的依赖从外面传进来，不在内部自己创建
func NewTaskService(repo TaskRepo, logger Logger) *TaskService {
	return &TaskService{
		repo:   repo,
		logger: logger,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, title string) (*Task, error) {
	if title == "" {
		return nil, errors.New("title is empty")
	}
	task := &Task{Title: title}
	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("创建任务: ID=%d, Title=%s", task.ID, task.Title))
	return task, nil
}

func (s *TaskService) GetTask(ctx context.Context, id int) (*Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TaskService) ListTasks(ctx context.Context) ([]Task, error) {
	return s.repo.List(ctx)
}

// ============ 6. App：组合根（Composition Root）============

// App 是"应用容器"，把所有对象组装在一起
// 这里就是"组合根"——整个程序唯一创建和组装对象的地方
type App struct {
	TaskService *TaskService
}

// NewApp 创建并组装所有对象
// 组装顺序：从底层到上层，像搭积木一样
func NewApp() *App {
	// 第 1 步：创建最底层——数据实现
	taskRepo := NewMemoryTaskRepo()

	// 第 2 步：创建基础设施——日志
	logger := NewConsoleLogger()

	// 第 3 步：把依赖注入到上层——业务层
	taskService := NewTaskService(taskRepo, logger)

	// 第 4 步：返回组装好的 App
	return &App{
		TaskService: taskService,
	}
}

// ============ 7. 对比：没有 DI 的写法（反例）============

// BadTaskService 是反例——硬编码依赖的写法
type BadTaskService struct {
	repo *memoryTaskRepo // ❌ 依赖具体类型，不是接口
}

func NewBadTaskService() *BadTaskService {
	// ❌ 在内部自己创建依赖，外部无法替换
	return &BadTaskService{
		repo: &memoryTaskRepo{
			tasks:  make(map[int]*Task),
			nextID: 1,
		},
	}
}

// ============ 8. main：程序入口 ============

func main() {
	fmt.Println("========== Lesson 28: 依赖注入思想 ==========")
	fmt.Println()

	// DI 方式：一行创建，所有依赖自动组装
	app := NewApp()

	ctx := context.Background()

	// 创建任务
	t1, _ := app.TaskService.CreateTask(ctx, "学习依赖注入")
	t2, _ := app.TaskService.CreateTask(ctx, "理解 Wire 原理")
	fmt.Printf("创建: [%d] %s\n", t1.ID, t1.Title)
	fmt.Printf("创建: [%d] %s\n", t2.ID, t2.Title)
	fmt.Println()

	// 查询单个
	task, _ := app.TaskService.GetTask(ctx, 1)
	fmt.Printf("查询: [%d] %s, Done=%v\n", task.ID, task.Title, task.Done)
	fmt.Println()

	// 列表
	tasks, _ := app.TaskService.ListTasks(ctx)
	fmt.Println("任务列表:")
	for _, t := range tasks {
		fmt.Printf("  [%d] %s (创建于 %s)\n", t.ID, t.Title, t.CreatedAt.Format("15:04:05"))
	}
	fmt.Println()

	// ========== 演示 DI 的核心好处 ==========
	fmt.Println("========== DI 的好处：替换实现不用改业务代码 ==========")
	fmt.Println()

	// 假设明天要换成 MySQL，只需要：
	//   taskRepo := NewMySQLTaskRepo(db)  // 改这一行
	//   taskService := NewTaskService(taskRepo, logger) // 这行不用改！
	fmt.Println("想换存储实现？")
	fmt.Println("  旧: taskRepo := NewMemoryTaskRepo()")
	fmt.Println("  新: taskRepo := NewMySQLTaskRepo(db)")
	fmt.Println("  TaskService 一行都不用改！")
	fmt.Println()

	// 假设想换成静默日志（测试时不输出），只需要：
	//   logger := NewSilentLogger()  // 改这一行
	fmt.Println("想换日志实现？")
	fmt.Println("  旧: logger := NewConsoleLogger()")
	fmt.Println("  新: logger := NewSilentLogger()")
	fmt.Println("  TaskService 一行都不用改！")
	fmt.Println()

	// ========== 对比反例：BadTaskService ==========
	fmt.Println("========== 反例：没有 DI 会怎样？ ==========")
	fmt.Println("BadTaskService 的问题：")
	fmt.Println("  1. 想换成 MySQL？得改 BadTaskService 内部代码")
	fmt.Println("  2. 想写单元测试？必须真的连数据库，无法用 fake")
	fmt.Println("  3. 想加日志？得改 BadTaskService 的构造函数")
	fmt.Println("  => 每改一次需求，就要改业务代码，耦合太紧！")
	fmt.Println()

	// ========== 和 issue_api 的关系 ==========
	fmt.Println("========== 和 issue_api 的关系 ==========")
	fmt.Println("本课手写的 NewApp() 就是 Wire 的「手动版」：")
	fmt.Println("  - 本课: NewApp() 里手动写 taskRepo := NewMemoryTaskRepo()")
	fmt.Println("  - Wire: 写一个 wire.go 描述依赖关系，Wire 自动生成这些代码")
	fmt.Println("  - 下一课 Lesson 29 就会学 Wire！")
	fmt.Println()

	// 避免 BadTaskService 未使用报错
	_ = NewBadTaskService
}
