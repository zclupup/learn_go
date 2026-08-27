package main

import "fmt"

// Lesson 36 用一个极简程序模拟 issue_api 的整体结构。
// 真实项目很复杂，本课只先建立地图：main -> server/router -> service -> biz -> repo/data。
func main() {
	fmt.Println("========== Lesson 36: issue_api 项目结构总览 ==========")
	fmt.Println()

	printProjectMap()
	printStartupFlow()
	printRequestFlow()
	runMiniIssueAPI()
	printRealCodeChecklist()
}

func printProjectMap() {
	fmt.Println("--- 1. issue_api 目录地图 ---")
	fmt.Println("cmd/issue_api/main.go        程序入口：读配置、初始化日志、组装 App、启动")
	fmt.Println("cmd/issue_api/wire.go        依赖注入声明：告诉 Wire 怎么把对象拼起来")
	fmt.Println("internal/server/gin.go       HTTP 服务和路由：URL 在这里注册到 handler")
	fmt.Println("internal/service/            Service 层：像 Gin handler/controller，处理请求和响应")
	fmt.Println("internal/biz/                Biz 层：业务逻辑，用例 UseCase")
	fmt.Println("internal/biz/repo/           Repo 接口：Biz 需要哪些数据能力")
	fmt.Println("internal/data/               Data 层：真正访问 MySQL、COS、Redis、外部系统")
	fmt.Println("internal/model/              数据库模型：GORM struct 和表字段映射")
	fmt.Println("internal/rto/                请求/响应对象：HTTP JSON 对应的结构体")
	fmt.Println("internal/conf/ + configs/    配置结构和配置文件")
	fmt.Println("pkg/tool/                    全局对象：GlobalConf、GlobalMysql、Logger 等")
	fmt.Println()
}

func printStartupFlow() {
	fmt.Println("--- 2. 启动链路：main.go 做什么 ---")
	fmt.Println("1) flag.Parse() 读取 -conf 参数，默认 ./configs")
	fmt.Println("2) Kratos config 从配置目录加载 yaml")
	fmt.Println("3) c.Scan(&bc) 把配置写入 conf.Bootstrap 结构体")
	fmt.Println("4) 初始化 logger，并写入 tool.GlobalConf / tool.GlobalLogger")
	fmt.Println("5) validateRequiredConfig 检查关键配置是否缺失")
	fmt.Println("6) biz.Init() 做业务初始化")
	fmt.Println("7) wireApp(...) 通过 Wire 组装 server/service/biz/data")
	fmt.Println("8) app.Run() 启动 HTTP server 和 cronjob server")
	fmt.Println()
}

func printRequestFlow() {
	fmt.Println("--- 3. 请求链路：以 batch_import_issue 为例 ---")
	fmt.Println("POST /api/v1/issue_manage/prod/batch_import_issue")
	fmt.Println("  -> internal/server/gin.go 注册路由")
	fmt.Println("  -> internal/service/prod.go: TaskService.HttpBatchImportIssue")
	fmt.Println("  -> RequestUnmarshal2(c, &req) 绑定 JSON 请求体")
	fmt.Println("  -> t.taskUseCase.BatchImportIssue(c, &req) 进入业务层")
	fmt.Println("  -> internal/biz/task/issue_import.go 按 event_uuid upsert")
	fmt.Println("  -> t.issueTrackingResultRepo.Xxx(...) 调 repo 接口")
	fmt.Println("  -> internal/data/... repo 实现访问数据库")
	fmt.Println("  -> service.Return(...) 统一包装 err_no / err_msg / results")
	fmt.Println()
}

func runMiniIssueAPI() {
	fmt.Println("--- 4. 玩具版 issue_api 链路演示 ---")

	repo := NewMemoryIssueRepo()
	useCase := NewIssueUseCase(repo)
	service := NewIssueService(useCase)

	req := ImportIssueRequest{
		EventUUID: "marker_event-LDP95C966TY007696-20260824154818",
		VIN:       "LDP95C966TY007696",
		Summary:   "单边超声波水平车位未释放",
	}

	resp := service.BatchImportIssue(req)
	fmt.Printf("HTTP response: err_no=%d, err_msg=%q, results=%+v\n", resp.ErrNo, resp.ErrMsg, resp.Results)
	fmt.Println()
}

func printRealCodeChecklist() {
	fmt.Println("--- 5. 以后追 issue_api 接口时的问题清单 ---")
	fmt.Println("1) 路由在哪里注册？")
	fmt.Println("2) 调的是哪个 service handler？")
	fmt.Println("3) 请求结构体在 internal/rto 还是 dto？")
	fmt.Println("4) service 调哪个 UseCase？")
	fmt.Println("5) UseCase 调哪个 repo 接口？")
	fmt.Println("6) data 层哪个结构体实现了 repo？")
	fmt.Println("7) 最后查哪张表、调哪个外部系统？")
	fmt.Println("8) 错误如何变成 err_no / err_msg 返回？")
}

type ImportIssueRequest struct {
	EventUUID string
	VIN       string
	Summary   string
}

type ImportIssueResult struct {
	TotalCnt   int
	SuccessCnt int
	FailCnt    int
}

type Response[T any] struct {
	ErrNo   int
	ErrMsg  string
	Results T
}

type IssueRepo interface {
	UpsertByEventUUID(eventUUID string, vin string, summary string) error
}

type MemoryIssueRepo struct {
	data map[string]ImportIssueRequest
}

func NewMemoryIssueRepo() *MemoryIssueRepo {
	return &MemoryIssueRepo{data: make(map[string]ImportIssueRequest)}
}

func (r *MemoryIssueRepo) UpsertByEventUUID(eventUUID string, vin string, summary string) error {
	r.data[eventUUID] = ImportIssueRequest{
		EventUUID: eventUUID,
		VIN:       vin,
		Summary:   summary,
	}
	fmt.Println("data: 按 event_uuid 写入或更新数据")
	return nil
}

type IssueUseCase struct {
	repo IssueRepo
}

func NewIssueUseCase(repo IssueRepo) *IssueUseCase {
	return &IssueUseCase{repo: repo}
}

func (u *IssueUseCase) BatchImportIssue(req ImportIssueRequest) (ImportIssueResult, error) {
	fmt.Println("biz: 检查 event_uuid，然后调用 repo")
	if err := u.repo.UpsertByEventUUID(req.EventUUID, req.VIN, req.Summary); err != nil {
		return ImportIssueResult{TotalCnt: 1, FailCnt: 1}, err
	}
	return ImportIssueResult{TotalCnt: 1, SuccessCnt: 1}, nil
}

type IssueService struct {
	useCase *IssueUseCase
}

func NewIssueService(useCase *IssueUseCase) *IssueService {
	return &IssueService{useCase: useCase}
}

func (s *IssueService) BatchImportIssue(req ImportIssueRequest) Response[ImportIssueResult] {
	fmt.Println("service: 绑定请求参数后，调用 biz/usecase")
	result, err := s.useCase.BatchImportIssue(req)
	if err != nil {
		return Response[ImportIssueResult]{ErrNo: 500500, ErrMsg: err.Error(), Results: result}
	}
	return Response[ImportIssueResult]{ErrNo: 0, ErrMsg: "", Results: result}
}
