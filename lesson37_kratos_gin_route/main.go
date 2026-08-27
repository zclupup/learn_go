package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ========== Lesson 37: Kratos + Gin 路由链路 ==========
// 目标：看懂 issue_api internal/server/gin.go 的路由注册逻辑
// 对比我们之前学的 Gin 入门 + Gin 分层小项目，Kratos 风格的路由写法只是"更整齐的分组"

func main() {
	fmt.Println("========== Lesson 37: Kratos + Gin 路由链路 ==========")
	fmt.Println()

	// 1. 创建 Gin 引擎（和我们之前 lesson22 写的完全一样）
	r := gin.Default()

	// 2. 用我们之前学过的路由分组，模拟 issue_api 写法
	// 真实 issue_api 里是 internal/server/gin.go 里的 NewGinServer
	registerRoutes(r)

	fmt.Println("路由表全部注册完成，打印所有注册的路由：")
	for _, route := range r.Routes() {
		fmt.Printf("  %-6s %s\n", route.Method, route.Path)
	}
	fmt.Println()

	fmt.Println("启动服务，访问 http://127.0.0.1:9090/healthz 验证")
	// 注释掉启动，避免卡终端，学习阶段看代码理解即可
	// _ = r.Run(":9090")
}

// registerRoutes 完全对照 issue_api internal/server/gin.go 的结构
// 真实项目里就是把不同业务模块的路由分组，一个个挂进去
func registerRoutes(r *gin.Engine) {
	// 2.1 全局健康检查、根路径路由，不属于任何业务模块
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// 2.2 模拟 /api/v1 全局前缀分组，和 issue_api 真实路径完全对齐
	v1 := r.Group("/api/v1")
	{
		// 2.3 模拟 issue_manage 业务分组
		issueManageGroup := v1.Group("/issue_manage")
		{
			// 2.4 prod 子分组
			prodGroup := issueManageGroup.Group("/prod")
			{
				// 就是你在 Lesson36 见过的那个著名批量导入接口
				prodGroup.POST("/batch_import_issue", func(c *gin.Context) {
					// 这里我们模拟一下，真实代码里这里把请求转发给 service.XXX 方法
					serviceHandler := &TaskService{}
					serviceHandler.HttpBatchImportIssue(c)
				})

				// 再加几个同路径下的其他接口，和真实项目一样
				prodGroup.GET("/issue_list", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{
						"err_no":  0,
						"err_msg": "",
						"results": []string{"issue1", "issue2"},
					})
				})
			}
		}
	}
}

// ========== 这就是 internal/service 里的 TaskService 样子 ==========
// 真实项目里 service 结构体里会提前注入好 biz/usecase 等依赖，Lesson36 玩具版我们已经见过
type TaskService struct{}

// HttpBatchImportIssue 就是路由最终绑定的 handler
func (s *TaskService) HttpBatchImportIssue(c *gin.Context) {
	// 第一步：绑定请求体，用 c.ShouldBindJSON 把 JSON 解析到 req 结构体
	// 这个 req 结构体就在 internal/rto 包里，我们之前学过的 struct tag + 指针传参
	var req BatchImportIssueReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err_no":  400,
			"err_msg": "参数错误: " + err.Error(),
			"results": nil,
		})
		return
	}

	// 第二步：真实项目这里调用 s.taskUseCase.BatchImportIssue(ctx, &req) 走业务层
	fmt.Printf("service 层收到请求: event_uuid=%s, vin=%s\n", req.EventUUID, req.VIN)

	// 第三步：统一返回格式，和 issue_api Return 函数的逻辑完全一致
	c.JSON(http.StatusOK, gin.H{
		"err_no":  0,
		"err_msg": "",
		"results": gin.H{"total_cnt": 1, "success_cnt": 1, "fail_cnt": 0},
	})
}

// BatchImportIssueReq 就是 internal/rto 下的请求结构体
type BatchImportIssueReq struct {
	EventUUID string `json:"event_uuid" binding:"required"`
	VIN       string `json:"vin" binding:"required"`
	Summary   string `json:"summary"`
}
