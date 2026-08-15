package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// ======== Lesson 22：Gin 入门 ========
//
// Gin 是 Go 里很常用的 Web 框架，可以理解成标准库 net/http 的增强版。
// 它帮我们把路由、参数读取、JSON 响应、请求体绑定等常见工作写得更简单。
//
// 和 Lesson 21 标准库写法对比：
// - http.HandleFunc(...)       → r.GET(...) / r.POST(...)
// - r.URL.Query().Get("name") → c.Query("name")
// - strings.TrimPrefix(...)    → c.Param("id")
// - json.NewDecoder(...).Decode(&req) → c.ShouldBindJSON(&req)
// - writeJSON(w, status, data) → c.JSON(status, data)

// User 是接口返回/接收的用户数据。
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// CreateUserRequest 是创建用户时的请求体。
// binding:"required" 是 Gin/validator 会读取的 tag，表示必填。
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

// practice
type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

type APIResponse struct {
	Message string `json:"message"`
}

var users = []User{
	{ID: 1, Name: "张三", Age: 20},
	{ID: 2, Name: "李四", Age: 25},
}

var nextUserID = 3
var usersMu = sync.Mutex{}

func main() {
	// gin.Default() 会创建一个 Gin 引擎，并默认带上日志和 recover 中间件。
	r := gin.Default()

	// GET /ping：最简单的健康检查接口。
	r.GET("/ping", pingHandler)

	// GET /hello?name=张三：读取 query 参数。
	r.GET("/hello", helloHandler)

	// /users 这一组接口用于演示列表和创建。
	r.GET("/users", listUsers)
	r.POST("/users", createUser)

	// GET /users/:id：读取路径参数。
	r.GET("/users/:id", getUserDetail)

	r.PUT("/users/:id", updateUser)

	r.DELETE("/users/:id", deleteUser)

	// 路由分组：以后真实项目常用 /api/v1 这类前缀。
	api := r.Group("/api")
	api.GET("/status", statusHandler)
	api.GET("/users", listUsers)

	// 启动服务。Gin 的 Run 本质上也是启动 HTTP 服务。
	r.Run(":8928")
}

func updateUser(c *gin.Context) {
	id, ok := getParamID(c)
	if !ok {
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{Message: "请求体不是合法 JSON，或缺少必填字段"})
		return
	}

	usersMu.Lock()
	defer usersMu.Unlock()

	// 更新用户
	for i, user := range users {
		if user.ID == id {
			users[i].Name = req.Name
			users[i].Age = req.Age
			c.JSON(http.StatusOK, users[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, APIResponse{Message: "用户不存在"})
}

func getParamID(c *gin.Context) (int, bool) {
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{Message: "用户id必须是数字"})
		return 0, false
	}
	return id, true
}

func deleteUser(c *gin.Context) {
	id, ok := getParamID(c)
	if !ok {
		return
	}
	// 删除用户
	usersMu.Lock()
	defer usersMu.Unlock()

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, APIResponse{Message: "用户已删除"})
			return
		}
	}
	c.JSON(http.StatusNotFound, APIResponse{Message: "用户不存在"})
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func helloHandler(c *gin.Context) {
	name := c.Query("name")
	name = strings.TrimSpace(name)
	if name == "" {
		name = "同学"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "你好，" + name,
	})
}

func listUsers(c *gin.Context) {
	usersMu.Lock()
	defer usersMu.Unlock()
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{Message: "请求体不是合法 JSON，或缺少必填字段"})
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		c.JSON(http.StatusBadRequest, APIResponse{Message: "name 不能为空"})
		return
	}

	user := User{
		ID:   nextUserID,
		Name: req.Name,
		Age:  req.Age,
	}

	usersMu.Lock()
	defer usersMu.Unlock()
	nextUserID++
	users = append(users, user)

	c.JSON(http.StatusCreated, user)
}

func getUserDetail(c *gin.Context) {
	// c.Param("id") 读取 /users/:id 里的 id，返回 string。
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{Message: "用户 id 必须是数字"})
		return
	}

	usersMu.Lock()
	defer usersMu.Unlock()

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	log.Printf("用户 id 不存在: %d", id)
	c.JSON(http.StatusNotFound, APIResponse{Message: "用户不存在"})
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"lesson": "Gin 入门",
	})
}
