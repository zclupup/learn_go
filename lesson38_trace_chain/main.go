package main

import (
	"errors"
	"fmt"
)

// ========== Lesson 38: 逐层追踪完整请求链路 ==========
//
// 本课目标：用 issue_api 真实接口 batch_import_issue，把 7 层链路每一层都拆开讲清楚。
// 读完本课，你拿到 issue_api 任意一个接口，都能自己从路由追到数据库。

func main() {
	fmt.Println("========== Lesson 38: 逐层追踪完整请求链路 ==========")
	fmt.Println()
	fmt.Println("追踪接口: POST /api/v1/issue_manage/prod/batch_import_issue")
	fmt.Println()

	// 整个链路分 7 步，每一步我们都在下面用注释标注清楚
	step1_route()
	step2_gin_register()
	step3_service_bind_request()
	step4_service_call_biz()
	step5_biz_call_repo()
	step6_data_access_db()
	step7_response_flow_back()
}

// ================================================================
// 第 1 步：URL 请求到达服务端
// ================================================================
func step1_route() {
	fmt.Println("━━━━━ 第 1 步：HTTP 请求到达 ━━━━━")
	fmt.Println("客户端发送: POST /api/v1/issue_manage/prod/batch_import_issue")
	fmt.Println("请求体 JSON:")
	fmt.Println(`  {`)
	fmt.Println(`    "event_uuid": "marker_event-LDP95C966TY007696-20260824154818",`)
	fmt.Println(`    "vin":       "LDP95C966TY007696",`)
	fmt.Println(`    "summary":   "单边超声波水平车位未释放"`)
	fmt.Println(`  }`)
	fmt.Println()
	fmt.Println("Gin 引擎收到请求后，会在路由表里找匹配的 pattern。")
	fmt.Println("路由表长这样（和 Lesson37 打印的完全一样）：")
	fmt.Println("  GET    /healthz")
	fmt.Println("  GET    /api/v1/issue_manage/prod/issue_list")
	fmt.Println("  POST   /api/v1/issue_manage/prod/batch_import_issue   ← 匹配这条！")
	fmt.Println()
	fmt.Println("匹配到路由后，Gin 执行这条路由后面绑定的 handler 函数。")
	fmt.Println()
}

// ================================================================
// 第 2 步：gin.go 路由注册
// ================================================================
func step2_gin_register() {
	fmt.Println("━━━━━ 第 2 步：路由注册（internal/server/gin.go）━━━━━")
	fmt.Println()
	fmt.Println("真实 issue_api 里 gin.go 的代码结构（简化版）：")
	fmt.Println()
	fmt.Println("  func NewGinServer(ts *service.TaskService) *gin.Engine {")
	fmt.Println("      r := gin.Default()")
	fmt.Println("")
	fmt.Println("      v1 := r.Group(\"/api/v1\")              // 第一层分组")
	fmt.Println("      issueManage := v1.Group(\"/issue_manage\") // 第二层分组")
	fmt.Println("      prod := issueManage.Group(\"/prod\")       // 第三层分组")
	fmt.Println("      {")
	fmt.Println("          prod.POST(\"/batch_import_issue\",       // 最终路径")
	fmt.Println("              ts.HttpBatchImportIssue)            // ← 绑定到 service 方法")
	fmt.Println("      }")
	fmt.Println("      return r")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("关键理解：")
	fmt.Println("  - ts 是 *service.TaskService，已经在 main 里通过 Wire 组装好")
	fmt.Println("  - ts.HttpBatchImportIssue 是 TaskService 的一个方法")
	fmt.Println("  - 方法签名必须是 func(c *gin.Context)，这是 Gin 的硬性要求")
	fmt.Println("  - 把 URL 和 handler 方法「绑定」的动作，就在这里完成")
	fmt.Println()
}

// ================================================================
// 第 3 步：Service 层绑定请求参数
// ================================================================
func step3_service_bind_request() {
	fmt.Println("━━━━━ 第 3 步：Service 层绑定请求（internal/service/prod.go）━━━━━")
	fmt.Println()

	// 模拟 Gin 的 Context
	mockCtx := newMockGinContext()
	mockCtx.requestBody = []byte(`{
		"event_uuid": "marker_event-LDP95C966TY007696-20260824154818",
		"vin":       "LDP95C966TY007696",
		"summary":   "单边超声波水平车位未释放"
	}`)

	// 创建 service，注入 useCase（第 5 步会用到）
	repo := &mockIssueRepo{}                       // data 层实现
	useCase := &IssueUseCase{repo: repo}           // biz 层
	service := &TaskService{taskUseCase: useCase}  // service 层注入 biz

	// 这就是 HttpBatchImportIssue 的完整实现
	service.HttpBatchImportIssue(mockCtx)

	fmt.Println("响应结果:", mockCtx.responseJSON)
	fmt.Println()
	fmt.Println("关键理解：")
	fmt.Println("  - HttpBatchImportIssue 接收 *gin.Context，这就是 Gin 传过来的请求上下文")
	fmt.Println("  - c.ShouldBindJSON(&req) 把请求体 JSON 解析到 req 结构体（需要传指针）")
	fmt.Println("  - req 结构体定义在 internal/rto/ 里，每个字段有 json tag 和 binding tag")
	fmt.Println("  - 绑定成功后，req 的每个字段都被填入了请求体里的值")
	fmt.Println()
}

// ================================================================
// 第 4 步：Service 调用 Biz/Usecase
// ================================================================
func step4_service_call_biz() {
	fmt.Println("━━━━━ 第 4 步：Service 调用 Biz（internal/biz/task/issue_import.go）━━━━━")
	fmt.Println()
	fmt.Println("HttpBatchImportIssue 方法的最后几步：")
	fmt.Println()
	fmt.Println("  func (s *TaskService) HttpBatchImportIssue(c *gin.Context) {")
	fmt.Println("      var req rto.BatchImportIssueReq")
	fmt.Println("      c.ShouldBindJSON(&req)          // 第 3 步：绑定参数")
	fmt.Println("")
	fmt.Println("      // ↓ 第 4 步：把请求和 context 传给 biz 层")
	fmt.Println("      result, err := s.taskUseCase.BatchImportIssue(c, &req)")
	fmt.Println("")
	fmt.Println("      if err != nil {")
	fmt.Println("          Return(c, 500500, err.Error(), result)   // 统一错误返回")
	fmt.Println("          return")
	fmt.Println("      }")
	fmt.Println("      Return(c, 0, \"\", result)   // 统一成功返回")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("关键理解：")
	fmt.Println("  - service 层不写任何业务逻辑，只做 3 件事：绑定参数 → 调用 biz → 返回响应")
	fmt.Println("  - c 是 *gin.Context，它同时实现了 context.Context 接口，所以能传给 biz")
	fmt.Println("  - Return 函数统一包装成 {err_no, err_msg, results} 格式")
	fmt.Println()
}

// ================================================================
// 第 5 步：Biz 层调用 Repo 接口
// ================================================================
func step5_biz_call_repo() {
	fmt.Println("━━━━━ 第 5 步：Biz 调用 Repo 接口（internal/biz/task/issue_import.go）━━━━━")
	fmt.Println()

	// 模拟 biz 层代码
	repo := &mockIssueRepo{}
	useCase := &IssueUseCase{repo: repo}

	req := BatchImportReq{
		EventUUID: "marker_event-LDP95C966TY007696-20260824154818",
		VIN:       "LDP95C966TY007696",
		Summary:   "单边超声波水平车位未释放",
	}
	result, err := useCase.BatchImportIssue(req)

	fmt.Printf("biz 返回: result=%+v, err=%v\n", result, err)
	fmt.Println()
	fmt.Println("Biz 层代码：")
	fmt.Println("  func (u *IssueUseCase) BatchImportIssue(req BatchImportReq) (Result, error) {")
	fmt.Println("      // 1. 业务校验")
	fmt.Println("      if req.EventUUID == \"\" {")
	fmt.Println("          return Result{}, ErrInvalidParam")
	fmt.Println("      }")
	fmt.Println("      // 2. 调用 repo 接口（不关心底层是 MySQL 还是内存）")
	fmt.Println("      err := u.repo.UpsertByEventUUID(req.EventUUID, req.VIN, req.Summary)")
	fmt.Println("      if err != nil {")
	fmt.Println("          return Result{TotalCnt: 1, FailCnt: 1}, err")
	fmt.Println("      }")
	fmt.Println("      return Result{TotalCnt: 1, SuccessCnt: 1}, nil")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("关键理解：")
	fmt.Println("  - Biz 层只依赖 IssueRepo 接口，不依赖具体实现")
	fmt.Println("  - u.repo 是什么？是 Lesson25/26 学过的 repo 模式：接口在 biz 定义，实现在 data")
	fmt.Println("  - 换数据库实现时，只改 data 层，biz 层一行不用动")
	fmt.Println()
}

// ================================================================
// 第 6 步：Data 层实现 Repo 接口，访问数据库
// ================================================================
func step6_data_access_db() {
	fmt.Println("━━━━━ 第 6 步：Data 层实现接口，访问数据库（internal/data/）━━━━━")
	fmt.Println()

	repo := &mockIssueRepo{}
	err := repo.UpsertByEventUUID("marker_event-LDP95C966TY007696-20260824154818", "LDP95C966TY007696", "单边超声波水平车位未释放")
	fmt.Printf("data 层执行结果: err=%v\n", err)
	fmt.Println()
	fmt.Println("真实 issue_api 里 data 层的代码（简化）：")
	fmt.Println("  type issueTrackingResultRepo struct {")
	fmt.Println("      db *gorm.DB    // GORM 数据库连接，在 NewXxx 的时候注入")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  func (r *issueTrackingResultRepo) UpsertByEventUUID(eventUUID, vin, summary string) error {")
	fmt.Println("      // 先用 event_uuid 查是否存在")
	fmt.Println("      var existing model.IssueTrackingResult")
	fmt.Println("      err := r.db.Where(\"event_uuid = ?\", eventUUID).First(&existing).Error")
	fmt.Println("      if err == gorm.ErrRecordNotFound {")
	fmt.Println("          // 不存在 → INSERT")
	fmt.Println("          return r.db.Create(&model.IssueTrackingResult{...}).Error")
	fmt.Println("      }")
	fmt.Println("      // 存在 → UPDATE")
	fmt.Println("      return r.db.Model(&existing).Updates(map[string]interface{}{...}).Error")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("关键理解：")
	fmt.Println("  - data 层结构体里持有 *gorm.DB（数据库连接），在 NewXxx 时注入")
	fmt.Println("  - 每个方法实现 repo 接口里的一个方法签名")
	fmt.Println("  - GORM 调用和 Lesson26 学的完全一样：Where → First/Find → Create/Updates")
	fmt.Println("  - 真实查的是 issue_tracking_result 表，字段映射在 internal/model/ 里定义")
	fmt.Println()
}

// ================================================================
// 第 7 步：响应原路返回
// ================================================================
func step7_response_flow_back() {
	fmt.Println("━━━━━ 第 7 步：响应原路返回 ━━━━━")
	fmt.Println()
	fmt.Println("响应返回路径（和数据流方向相反）：")
	fmt.Println()
	fmt.Println("  data 层返回 (result, error)")
	fmt.Println("       ↓")
	fmt.Println("  biz 层拿到 result，包装或透传")
	fmt.Println("       ↓")
	fmt.Println("  service 层调用 Return(c, err_no, err_msg, results)")
	fmt.Println("       ↓")
	fmt.Println("  Return 函数内部调用 c.JSON(http.StatusOK, Response{...})")
	fmt.Println("       ↓")
	fmt.Println("  Gin 把 JSON 写入 HTTP Response，返回给客户端")
	fmt.Println()
	fmt.Println("最终客户端收到的 JSON：")
	fmt.Println(`  {`)
	fmt.Println(`    "err_no":  0,`)
	fmt.Println(`    "err_msg": "",`)
	fmt.Println(`    "results": {`)
	fmt.Println(`      "total_cnt":   1,`)
	fmt.Println(`      "success_cnt": 1,`)
	fmt.Println(`      "fail_cnt":    0`)
	fmt.Println(`    }`)
	fmt.Println(`  }`)
	fmt.Println()
	fmt.Println("注意：")
	fmt.Println("  - 每一层只把错误往上抛，不自己处理（除非这层能真正处理）")
	fmt.Println("  - service 层是唯一负责把 error 翻译成 HTTP 状态码和 err_no 的地方")
	fmt.Println("  - biz 和 data 层不依赖 Gin，不依赖 HTTP，所以可以复用给 gRPC、命令行等")
	fmt.Println()
}

// ================================================================
// 全链路总结
// ================================================================
func init() {
	// 在所有 step 函数执行完后，main 里打印总结
}

// ================================================================
// 下面是本课用到的类型定义（和 Lesson36 一样，但加了更详细的注释）
// ================================================================

// --- 请求/响应结构体（对应 internal/rto/） ---

// BatchImportReq 是 service 层从 HTTP 请求体绑定的结构体
// json tag 负责字段名映射，binding tag 负责参数校验
type BatchImportReq struct {
	EventUUID string `json:"event_uuid" binding:"required"`
	VIN       string `json:"vin" binding:"required"`
	Summary   string `json:"summary"`
}

// ImportResult 是 biz 层返回给 service 的结果结构体
type ImportResult struct {
	TotalCnt   int `json:"total_cnt"`
	SuccessCnt int `json:"success_cnt"`
	FailCnt    int `json:"fail_cnt"`
}

// --- 统一响应格式（对应 pkg/tool/ 或 internal/rto/ 里的 Return） ---

// Response 泛型统一响应体 —— 对应 Lesson29 学的泛型
type Response[T any] struct {
	ErrNo   int    `json:"err_no"`
	ErrMsg  string `json:"err_msg"`
	Results T      `json:"results"`
}

// Return 模拟 issue_api 的统一返回函数
func Return[T any](errNo int, errMsg string, results T) Response[T] {
	return Response[T]{ErrNo: errNo, ErrMsg: errMsg, Results: results}
}

// --- Repo 接口（对应 internal/biz/repo/） ---

// IssueRepo 是 biz 层定义的数据能力接口
// biz 层说："我需要能按 event_uuid 写入或更新数据"
// 谁实现了这两个方法，谁就能当 IssueRepo 用
type IssueRepo interface {
	UpsertByEventUUID(eventUUID, vin, summary string) error
}

// --- Biz 层（对应 internal/biz/task/） ---

// IssueUseCase 是业务用例，持有 repo 接口（不持有具体实现）
type IssueUseCase struct {
	repo IssueRepo // 接口类型！不是具体类型
}

// BatchImportIssue 是 biz 层的核心业务方法
func (u *IssueUseCase) BatchImportIssue(req BatchImportReq) (ImportResult, error) {
	fmt.Println("    [biz 层] 收到请求，开始业务处理")

	// 业务校验：参数不能为空
	if req.EventUUID == "" {
		fmt.Println("    [biz 层] event_uuid 为空，返回参数错误")
		return ImportResult{}, errors.New("event_uuid 不能为空")
	}

	fmt.Printf("    [biz 层] 调用 repo.UpsertByEventUUID(event_uuid=%s)\n", req.EventUUID)

	// 调用 repo 接口 —— 不关心底层是 MySQL 还是内存
	if err := u.repo.UpsertByEventUUID(req.EventUUID, req.VIN, req.Summary); err != nil {
		fmt.Printf("    [biz 层] repo 返回错误: %v\n", err)
		return ImportResult{TotalCnt: 1, FailCnt: 1}, err
	}

	fmt.Println("    [biz 层] 处理成功，返回结果")
	return ImportResult{TotalCnt: 1, SuccessCnt: 1}, nil
}

// --- Data 层（对应 internal/data/） ---

// mockIssueRepo 模拟 data 层的 MySQL 实现
// 真实项目里这个结构体持有 *gorm.DB
type mockIssueRepo struct {
	// 真实项目里这里是: db *gorm.DB
	storage map[string]BatchImportReq // 用内存模拟数据库
}

func (r *mockIssueRepo) UpsertByEventUUID(eventUUID, vin, summary string) error {
	fmt.Printf("    [data 层] 收到查询请求，event_uuid=%s\n", eventUUID)
	fmt.Println("    [data 层] 执行: SELECT * FROM issue_tracking_result WHERE event_uuid = ?")
	fmt.Println("    [data 层] 未找到记录，执行 INSERT")
	fmt.Println("    [data 层] 数据库写入成功")
	return nil
}

// --- Service 层（对应 internal/service/） ---

// TaskService 是 service 层，持有 biz useCase
// 真实项目里还可能持有多个 useCase 和其他依赖
type TaskService struct {
	taskUseCase *IssueUseCase // 注入的 biz 层
}

// HttpBatchImportIssue 是路由绑定的 Gin handler 方法
// 签名必须是 func(c *gin.Context)，这是 Gin 的硬性要求
func (s *TaskService) HttpBatchImportIssue(c *mockGinContext) {
	fmt.Println("  [service 层] 收到 HTTP 请求")

	// 第 3 步：绑定请求体 JSON 到 req 结构体
	var req BatchImportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("  [service 层] 参数绑定失败: %v\n", err)
		c.JSON(400, Return[ImportResult](400, "参数错误: "+err.Error(), ImportResult{}))
		return
	}
	fmt.Printf("  [service 层] 请求参数绑定成功: event_uuid=%s, vin=%s\n", req.EventUUID, req.VIN)

	// 第 4 步：调用 biz 层
	fmt.Println("  [service 层] 调用 useCase.BatchImportIssue")
	result, err := s.taskUseCase.BatchImportIssue(req)

	// 第 7 步：根据 biz 返回结果，组装 HTTP 响应
	if err != nil {
		fmt.Printf("  [service 层] biz 返回错误，组装 err_no=500500 响应\n")
		c.JSON(200, Return(500500, err.Error(), result))
		return
	}

	fmt.Println("  [service 层] 成功，组装 err_no=0 响应")
	c.JSON(200, Return(0, "", result))
}

// --- 模拟 Gin Context（只为实现本课演示，不需要理解） ---

type mockGinContext struct {
	requestBody  []byte
	responseJSON string
}

func newMockGinContext() *mockGinContext {
	return &mockGinContext{}
}

func (c *mockGinContext) ShouldBindJSON(obj interface{}) error {
	// 模拟 Gin 的 ShouldBindJSON：把 requestBody 填到 obj 里
	// 真实 Gin 会做 JSON 解析 + 结构体 tag 校验，这里简化
	req := obj.(*BatchImportReq)
	req.EventUUID = "marker_event-LDP95C966TY007696-20260824154818"
	req.VIN = "LDP95C966TY007696"
	req.Summary = "单边超声波水平车位未释放"
	return nil
}

func (c *mockGinContext) JSON(code int, obj interface{}) {
	c.responseJSON = fmt.Sprintf("HTTP %d: %+v", code, obj)
}