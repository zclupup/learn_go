package main

import (
	"errors"
	"fmt"
)

// ========== Lesson 39: Wire 自动依赖注入 ==========
//
// 目标：理解 Wire 是什么、怎么用，以及 issue_api 里 wire.go 和 wire_gen.go 的关系。
// 分三步走：
//   1. 手写 DI（你已经会的）
//   2. 用 Wire 替代手写（本课新增）
//   3. 对照 issue_api 真实项目

func main() {
	fmt.Println("========== Lesson 39: Wire 自动依赖注入 ==========")
	fmt.Println()

	step1_manual_di()
	step2_wire_concept()
	step3_issue_api_mapping()
}

// ================================================================
// 第 1 步：手写 DI（回顾 Lesson28）
// ================================================================
func step1_manual_di() {
	fmt.Println("━━━ 第 1 步：手写依赖注入（你已经会的）━━━")
	fmt.Println()

	// 这就是你之前一直在做的事情：在 main 里手动 New 对象，把依赖传进去
	repo := NewUserRepo()           // 造 data 层
	usecase := NewUserUseCase(repo) // 把 repo 注入 usecase
	service := NewUserService(usecase) // 把 usecase 注入 service

	result, err := service.GetUser("zhangcl")
	fmt.Printf("service.GetUser 返回: result=%+v, err=%v\n", result, err)
	fmt.Println()
	fmt.Println("手写 DI 的问题：")
	fmt.Println("  - 对象多了，main 里几十行 NewXxx 代码，容易漏参数")
	fmt.Println("  - 改了一个构造函数的签名（比如 NewUserUseCase 多了一个参数），main 也要改")
	fmt.Println("  - Wire 就是帮你自动生成这部分代码，不是替代这部分逻辑")
	fmt.Println()
}

// ================================================================
// 第 2 步：Wire 概念——用代码演示
// ================================================================
func step2_wire_concept() {
	fmt.Println("━━━ 第 2 步：Wire 的概念（用代码模拟）━━━")
	fmt.Println()

	// Wire 做三件事：
	// ① 你写 wire.go：声明"谁提供什么"——这叫 Provider
	// ② 运行 wire 命令
	// ③ Wire 生成 wire_gen.go：自动写出 NewXxx 组装代码

	fmt.Println("【wire.go 里你写的内容（概念版）】：")
	fmt.Println("  func InitApp() *UserService {")
	fmt.Println("      wire.Build(")
	fmt.Println("          NewUserRepo,       // Provider: 提供 *UserRepo")
	fmt.Println("          NewUserUseCase,    // Provider: 提供 *UserUseCase（需要 *UserRepo）")
	fmt.Println("          NewUserService,    // Provider: 提供 *UserService（需要 *UserUseCase）")
	fmt.Println("      )")
	fmt.Println("      return nil  // 这个 return 不会执行，Wire 只读 wire.Build 里的内容")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("【wire_gen.go 里 Wire 自动生成的代码】：")
	fmt.Println("  func InitApp() *UserService {")
	fmt.Println("      repo := NewUserRepo()           // ← 自动生成")
	fmt.Println("      usecase := NewUserUseCase(repo)  // ← 自动生成")
	fmt.Println("      service := NewUserService(usecase) // ← 自动生成")
	fmt.Println("      return service")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("关键理解：")
	fmt.Println("  - Wire 不改变你的代码结构，只是把 main 里的组装代码搬到 wire_gen.go")
	fmt.Println("  - Wire 通过函数签名推导依赖关系：NewUserUseCase 的参数是 *UserRepo，")
	fmt.Println("    所以 Wire 自动去找能提供 *UserRepo 的 Provider（也就是 NewUserRepo）")
	fmt.Println("  - 编译后 wire_gen.go 就是普通 Go 代码，Wire 完全退场")
	fmt.Println()
}

// ================================================================
// 第 3 步：对照 issue_api 真实项目
// ================================================================
func step3_issue_api_mapping() {
	fmt.Println("━━━ 第 3 步：对照 issue_api 真实项目 ━━━")
	fmt.Println()

	fmt.Println("【issue_api 的 wire.go（cmd/issue_api/wire.go）】：")
	fmt.Println("  func wireApp(...) (*kratos.App, func(), error) {")
	fmt.Println("      panic(wire.Build(")
	fmt.Println("          server.ProviderSet,   // ← 不是单个函数，是一组函数！")
	fmt.Println("          service.ProviderSet,")
	fmt.Println("          biz.ProviderSet,")
	fmt.Println("          data.ProviderSet,")
	fmt.Println("          account.ProviderSet,")
	fmt.Println("          newApp,")
	fmt.Println("      ))")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("【ProviderSet 是什么？】")
	fmt.Println("  ProviderSet 就是把一组相关的 Provider 函数打包在一起：")
	fmt.Println()
	fmt.Println("  // internal/data/data.go:")
	fmt.Println("  var ProviderSet = wire.NewSet(")
	fmt.Println("      NewData, NewMysqlDefault, NewCosRepo, NewCosServiceDefault,")
	fmt.Println("      NewSceneRepo, NewTaskLabelRepo, NewIssueTrackingResultRepo,")
	fmt.Println("      NewIssueTrackingTaskRepo, NewPackRepo, ...  // 十几个 NewXxx")
	fmt.Println("  )")
	fmt.Println()
	fmt.Println("  // internal/service/service.go:")
	fmt.Println("  var ProviderSet = wire.NewSet(")
	fmt.Println("      NewPackService, NewUserService, NewSceneService,")
	fmt.Println("      NewCommonService, NewTaskService, NewJiraSyncService, NewAlarmService,")
	fmt.Println("  )")
	fmt.Println()
	fmt.Println("  这样 wire.go 里只需要写 5 行，而不是 30 个函数名。")
	fmt.Println()

	fmt.Println("【wire_gen.go 里生成的真实代码（简化版）】：")
	fmt.Println("  func wireApp(...) (*kratos.App, func(), error) {")
	fmt.Println("      db := data.NewMysqlDefault(confData)          // 造数据库连接")
	fmt.Println("      dataData, cleanup, _ := data.NewData(logger, cosService, db)")
	fmt.Println("      packRepo := data.NewPackRepo(dataData, logger)")
	fmt.Println("      packUseCase := pack.NewPackUseCase(packRepo, ...)")
	fmt.Println("      packService := service.NewPackService(logger, packUseCase)")
	fmt.Println("      // ... 几十行类似的组装")
	fmt.Println("      httpServer := server.NewHttpServer(confServer, packService, ...)")
	fmt.Println("      app := newApp(logger, httpServer, cronjobServer)")
	fmt.Println("      return app, cleanup, nil")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("【对照你之前的玩具版代码】：")
	fmt.Println("  你手写:    repo := NewMemoryIssueRepo()")
	fmt.Println("            useCase := NewIssueUseCase(repo)")
	fmt.Println("            service := NewIssueService(useCase)")
	fmt.Println()
	fmt.Println("  Wire生成:  packRepo := data.NewPackRepo(dataData, logger)")
	fmt.Println("            packUseCase := pack.NewPackUseCase(packRepo, ...)")
	fmt.Println("            packService := service.NewPackService(logger, packUseCase)")
	fmt.Println()
	fmt.Println("  完全一样的模式！只是对象多、名字长，本质没变。")
}

// ================================================================
// 下面是用到的类型定义（和之前课程一样）
// ================================================================

type User struct {
	Name string
	Age  int
}

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) FindByName(name string) (*User, error) {
	if name == "zhangcl" {
		return &User{Name: "zhangcl", Age: 23}, nil
	}
	return nil, errors.New("用户不存在")
}

type UserUseCase struct {
	repo *UserRepo
}

func NewUserUseCase(repo *UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) GetUser(name string) (*User, error) {
	return u.repo.FindByName(name)
}

type UserService struct {
	useCase *UserUseCase
}

func NewUserService(useCase *UserUseCase) *UserService {
	return &UserService{useCase: useCase}
}

func (s *UserService) GetUser(name string) (*User, error) {
	return s.useCase.GetUser(name)
}