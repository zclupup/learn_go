package main

import "fmt"

// ========== Lesson 40: issue_api main 函数启动逻辑 ==========
//
// 目标：逐行读懂 issue_api cmd/issue_api/main.go 的启动流程。
// 一共 8 步，每步对应 main.go 里的几行代码。

func main() {
	fmt.Println("========== Lesson 40: issue_api main 函数启动逻辑 ==========")
	fmt.Println()
	fmt.Println("模拟启动: go run ./cmd/issue_api -conf ./configs")
	fmt.Println()

	// 模拟 main 函数参数
	flagconf := "./configs"

	step1_flagParse(flagconf)
	step2_loadConfig(flagconf)
	step3_initLogger()
	step4_validateConfig()
	step5_otelInit()
	step6_bizInit()
	step7_wireApp()
	step8_appRun()
}

// ================================================================
// 第 1 步：flag.Parse() — 读取命令行参数
// ================================================================
func step1_flagParse(flagconf string) {
	fmt.Println("━━━━━ 第 1 步：flag.Parse() 读取命令行参数 ━━━━━")
	fmt.Println()

	fmt.Println("main.go 代码：")
	fmt.Println("  func init() {")
	fmt.Println("      flag.StringVar(&flagconf, \"conf\", \"./configs\", \"config path\")")
	fmt.Println("  }")
	fmt.Println("  func main() {")
	fmt.Println("      flag.Parse()")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("执行效果：")
	fmt.Printf("  读取 -conf 参数，得到配置目录路径: %s\n", flagconf)
	fmt.Println("  如果运行: go run ./cmd/issue_api -conf ./configs")
	fmt.Println("  则 flagconf = ./configs")
	fmt.Println()

	fmt.Println("flag 是 Go 标准库（Lesson27 学过），用来解析命令行参数。")
	fmt.Println("init() 在 main() 之前自动执行，注册 -conf 参数，默认值 ./configs。")
	fmt.Println()
}

// ================================================================
// 第 2 步：加载配置 — 读取 yaml → 解析到 struct
// ================================================================
func step2_loadConfig(flagconf string) {
	fmt.Println("━━━━━ 第 2 步：加载配置 ━━━━━")
	fmt.Println()

	fmt.Println("main.go 代码：")
	fmt.Println("  c := config.New(config.WithSource(file.NewSource(flagconf)))")
	fmt.Println("  c.Load()")
	fmt.Println()
	fmt.Println("  var bc conf.Bootstrap")
	fmt.Println("  c.Scan(&bc)")
	fmt.Println()

	fmt.Println("执行效果（和 Lesson27 学的 yaml 解析完全一样）：")
	fmt.Println("  1. 读取 ./configs/ 目录下的所有 yaml 文件")
	fmt.Println("  2. 把 yaml 内容解析到 conf.Bootstrap 结构体")
	fmt.Println("  3. bc 就是项目全局配置，包含 server/data/project/file 等子配置")
	fmt.Println()

	fmt.Println("Bootstrap 结构体大致长这样：")
	fmt.Println("  type Bootstrap struct {")
	fmt.Println("      Server  *conf.Server   // HTTP 端口、超时设置")
	fmt.Println("      Data    *conf.Data     // MySQL、Redis、COS 连接信息")
	fmt.Println("      Project *conf.Project  // 项目级参数：模式、飞书 webhook 等")
	fmt.Println("      File    *conf.File     // 文件路径、磁盘挂载点")
	fmt.Println("      Log     *conf.Log      // 日志级别、输出路径")
	fmt.Println("  }")
	fmt.Println()
}

// ================================================================
// 第 3 步：初始化日志
// ================================================================
func step3_initLogger() {
	fmt.Println("━━━━━ 第 3 步：初始化日志 ━━━━━")
	fmt.Println()

	fmt.Println("main.go 代码：")
	fmt.Println("  logger := zaplog.Logger(bc.Log)")
	fmt.Println("  tool.GlobalConf = &bc")
	fmt.Println("  tool.GlobalLogger = logger")
	fmt.Println("  tool.GlobalLoggerHelper = log.NewHelper(logger)")
	fmt.Println()

	fmt.Println("执行效果：")
	fmt.Println("  1. 根据配置的日志级别、输出路径，创建 logger 实例")
	fmt.Println("  2. 把配置和 logger 存入全局变量 tool.GlobalConf / tool.GlobalLogger")
	fmt.Println("  3. 项目里任何地方都能通过 tool.Log() 拿到 logger 写日志")
	fmt.Println()

	fmt.Println("⭐ 和 Lesson31 学的 slog 一样：")
	fmt.Println("  先创建 logger → 设为全局默认 → 项目各处直接写日志")
	fmt.Println()
}

// ================================================================
// 第 4 步：校验关键配置
// ================================================================
func step4_validateConfig() {
	fmt.Println("━━━━━ 第 4 步：校验关键配置 ━━━━━")
	fmt.Println()

	fmt.Println("main.go 代码：")
	fmt.Println("  func validateRequiredConfig(bc *conf.Bootstrap) {")
	fmt.Println("      sc := bc.Project.GetSimCollection()")
	fmt.Println("      if sc == nil { panic(...) }")
	fmt.Println("      // 检查 sim_collection 下的 6 个 ID 是否为空")
	fmt.Println("      for name, val := range required {")
	fmt.Println("          if val == \"\" { panic(...) }")
	fmt.Println("      }")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("执行效果：")
	fmt.Println("  检查配置文件中 project.sim_collection 的 6 个 ID 是否都填了")
	fmt.Println("  任何一个缺失 → 直接 panic，程序启动失败")
	fmt.Println("  这是「快速失败」原则：启动时发现配置不全，立即报错，不等到运行时才暴露")
	fmt.Println()
}

// ================================================================
// 第 5 步：初始化 OTel（OpenTelemetry 链路追踪）
// ================================================================
func step5_otelInit() {
	fmt.Println("━━━━━ 第 5 步：初始化 OTel 链路追踪 ━━━━━")
	fmt.Println()

	fmt.Println("main.go 代码：")
	fmt.Println("  otelShutdown := server.InitOtel(&bc)")
	fmt.Println("  defer func() {")
	fmt.Println("      if otelShutdown != nil {")
	fmt.Println("          ctx, cancel := context.WithTimeout(context.Background(), 5s)")
	fmt.Println("          defer cancel()")
	fmt.Println("          otelShutdown(ctx)  // 程序退出时关闭 OTel")
	fmt.Println("      }")
	fmt.Println("  }()")
	fmt.Println()

	fmt.Println("OTel 是 OpenTelemetry 的缩写，用来做链路追踪。")
	fmt.Println("简单说就是：记录每个请求从「进服务」到「出服务」的全过程，")
	fmt.Println("包括调了哪些数据库、耗时多少，方便排查性能问题。")
	fmt.Println()

	fmt.Println("⭐ 和 Lesson15 的 context 超时控制一样：")
	fmt.Println("  defer 里的 5 秒超时确保程序退出时能优雅关闭 OTel，不会永远卡住")
	fmt.Println()
}

// ================================================================
// 第 6 步：biz.Init() — 业务初始化
// ================================================================
func step6_bizInit() {
	fmt.Println("━━━━━ 第 6 步：biz.Init() 业务初始化 ━━━━━")
	fmt.Println()

	fmt.Println("main.go 代码：")
	fmt.Println("  // 里面不能有数据库相关操作，跟数据库有关操作，移动到crontab")
	fmt.Println("  biz.Init()")
	fmt.Println()

	fmt.Println("biz.Init() 做了这些事（internal/biz/init.go）：")
	fmt.Println("  1. 如果是集群模式，初始化 Redis 连接")
	fmt.Println("  2. 初始化雪花 ID 生成器（用于生成唯一 ID）")
	fmt.Println("  3. 初始化飞书、Jira 等外部服务配置")
	fmt.Println("  4. 根据 CPU 核数，计算并发解包数量（DoUnpackNum）")
	fmt.Println("  5. 创建解包队列 channel（Lesson14 学过的 channel！）")
	fmt.Println("  6. 写 pid 文件到磁盘")
	fmt.Println()

	fmt.Println("⭐ 关键注释：")
	fmt.Println("  \"里面不能有数据库相关操作\" — 因为此时数据库还没连上")
	fmt.Println("  数据库连接是在下一步 wireApp 里通过 data.NewData() 创建的")
	fmt.Println()
}

// ================================================================
// 第 7 步：wireApp() — 组装所有对象
// ================================================================
func step7_wireApp() {
	fmt.Println("━━━━━ 第 7 步：wireApp() 组装所有对象 ━━━━━")
	fmt.Println()

	fmt.Println("main.go 代码：")
	fmt.Println("  app, cleanup, err := wireApp(bc.Server, bc.Data, bc.File, bc.Project, logger)")
	fmt.Println("  if err != nil { panic(err) }")
	fmt.Println("  defer cleanup()")
	fmt.Println()

	fmt.Println("wireApp 是 Wire 生成的函数（Lesson39 学过的），它做了这些事：")
	fmt.Println("  1. 创建数据库连接（MySQL）")
	fmt.Println("  2. 创建所有 data 层的 repo 实现")
	fmt.Println("  3. 创建所有 biz 层的 usecase")
	fmt.Println("  4. 创建所有 service 层的 handler")
	fmt.Println("  5. 创建 HTTP server 和 Cronjob server")
	fmt.Println("  6. 组装成 kratos.App 对象返回")
	fmt.Println()

	fmt.Println("wire_gen.go 里生成的代码大致长这样（简化版）：")
	fmt.Println("  func wireApp(...) (*kratos.App, func(), error) {")
	fmt.Println("      db := data.NewMysqlDefault(confData)")
	fmt.Println("      dataData, cleanup, _ := data.NewData(logger, cosService, db)")
	fmt.Println("      packRepo := data.NewPackRepo(dataData, logger)")
	fmt.Println("      // ... 几十行类似的 NewXxx")
	fmt.Println("      httpServer := server.NewHttpServer(confServer, packService, ...)")
	fmt.Println("      cronjobServer := server.NewCronjobServer(...)")
	fmt.Println("      app := newApp(logger, httpServer, cronjobServer)")
	fmt.Println("      return app, func() { cleanup() }, nil")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("⭐ 注意 wireApp 的返回值：")
	fmt.Println("  - *kratos.App：组装好的应用对象，包含 HTTP server 和 Cronjob server")
	fmt.Println("  - func()：清理函数，程序退出时关闭数据库连接等资源")
	fmt.Println("  - error：组装过程中可能出错（比如数据库连不上）")
	fmt.Println()
}

// ================================================================
// 第 8 步：app.Run() — 启动服务
// ================================================================
func step8_appRun() {
	fmt.Println("━━━━━ 第 8 步：app.Run() 启动服务 ━━━━━")
	fmt.Println()

	fmt.Println("main.go 代码：")
	fmt.Println("  if err := app.Run(); err != nil {")
	fmt.Println("      panic(err)")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("app.Run() 做了这些事：")
	fmt.Println("  1. 启动 HTTP 服务，监听端口（如 8080）")
	fmt.Println("  2. 启动 Cronjob 定时任务服务")
	fmt.Println("  3. 阻塞等待退出信号（Ctrl+C 或 SIGTERM）")
	fmt.Println("  4. 收到退出信号后，优雅关闭所有服务")
	fmt.Println()

	fmt.Println("⭐ 至此，整个启动链路完成，项目开始提供服务。")
	fmt.Println()
	fmt.Println("收到退出信号时，defer 链会按逆序执行：")
	fmt.Println("  1. defer cleanup()  → 关闭数据库连接")
	fmt.Println("  2. defer otelShutdown(ctx) → 关闭链路追踪")
	fmt.Println("  3. defer c.Close() → 关闭配置读取器")
	fmt.Println("  4. 所有 defer 执行完，进程退出")
	fmt.Println()
}

func init() {
	// init 函数在 main 之前自动执行
	// 真实项目里 init() 注册 flag 参数，这里模拟展示
}

// ================================================================
// 全链路总结
// ================================================================

func init() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  issue_api 启动流程总览（8 步）")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  1. flag.Parse()        → 读取 -conf 参数，确定配置目录")
	fmt.Println("  2. config.Load()       → 读取 yaml，解析到 Bootstrap 结构体")
	fmt.Println("  3. 初始化日志           → 创建 logger，存入全局变量")
	fmt.Println("  4. validateRequiredConfig → 校验关键配置是否缺失")
	fmt.Println("  5. InitOtel()          → 初始化链路追踪")
	fmt.Println("  6. biz.Init()          → 业务初始化（Redis、雪花ID、队列）")
	fmt.Println("  7. wireApp()           → Wire 生成组装代码，创建所有对象")
	fmt.Println("  8. app.Run()           → 启动 HTTP + Cronjob，等待退出信号")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("⭐ 每步对应的之前学过的知识点：")
	fmt.Println("  第1步 → Lesson27 flag 用法")
	fmt.Println("  第2步 → Lesson27 yaml 配置解析")
	fmt.Println("  第3步 → Lesson31 slog 日志")
	fmt.Println("  第4步 → Lesson09 error/panic")
	fmt.Println("  第5步 → Lesson15 context 超时控制")
	fmt.Println("  第6步 → Lesson14 channel 队列")
	fmt.Println("  第7步 → Lesson28/39 依赖注入 + Wire")
	fmt.Println("  第8步 → Lesson20/22 Gin 启动服务")
	fmt.Println()
	fmt.Println("所有知识点你都已经学过了！")
}