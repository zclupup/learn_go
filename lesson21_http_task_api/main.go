package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ======== Lesson 21：标准库 HTTP 小项目 ========
//
// 这一课用 net/http 做一个内存版任务 API，模拟后端项目里的基础接口：
// - GET  /tasks      查询任务列表
// - POST /tasks      创建任务
// - GET  /tasks/{id} 查询单个任务
// - PUT  /tasks/{id} 修改任务状态
// - DELETE /tasks/{id} 删除任务
//
// 新知识点：
// - 用 switch r.Method 分发不同 HTTP 方法
// - 用 json.NewDecoder(r.Body).Decode(&req) 读取 JSON 请求体
// - 用 strconv.Atoi 把路径里的 id 字符串转成 int
// - 用 sync.Mutex 保护内存里的共享切片
// - 用不同 HTTP 状态码表达成功、参数错误、找不到、方法不支持

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at"`
}

type CreateTaskRequest struct {
	Title string `json:"title"`
}

type APIResponse struct {
	Message string `json:"message"`
}

type UpdateTaskRequest struct {
	Done bool `json:"done"`
}

var (
	tasks = []Task{
		{ID: 1, Title: "复习 JSON", Done: true, CreatedAt: "2026-08-13 09:00:00"},
		{ID: 2, Title: "练习 net/http", Done: false, CreatedAt: "2026-08-13 10:00:00"},
	}
	nextTaskID = 3
	tasksMu    sync.Mutex
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskDetailHandler)

	addr := ":8928"
	fmt.Println("任务 API 启动：http://localhost" + addr)
	fmt.Println("查询任务：curl -s http://localhost:8928/tasks")
	fmt.Println(`创建任务：curl -s -X POST http://localhost:8928/tasks -H "Content-Type: application/json" -d '{"title":"学习测试"}'`)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		writeJSON(w, http.StatusNotFound, APIResponse{Message: "接口不存在"})
		return
	}

	writeJSON(w, http.StatusOK, APIResponse{Message: "欢迎使用任务 API"})
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, APIResponse{Message: "只支持 GET 或 POST 请求"})
	}
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	tasksMu.Lock()
	defer tasksMu.Unlock()

	writeJSON(w, http.StatusOK, tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, APIResponse{Message: "请求体不是合法 JSON"})
		return
	}

	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		writeJSON(w, http.StatusBadRequest, APIResponse{Message: "title 不能为空"})
		return
	}

	tasksMu.Lock()
	defer tasksMu.Unlock()

	task := Task{
		ID:        nextTaskID,
		Title:     req.Title,
		Done:      false,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	nextTaskID++
	tasks = append(tasks, task)

	writeJSON(w, http.StatusCreated, task)
}

func parseTaskID(w http.ResponseWriter, r *http.Request) (int, bool) {
	idText := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idText)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, APIResponse{Message: "任务 id 必须是数字"})
		return 0, false
	}
	return id, true
}

func taskDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseTaskID(w, r)
	if !ok {
		return
	}
	log.Println("请求方法:", r.Method, "任务 ID:", id)
	switch r.Method {
	case http.MethodGet:
		getTaskDetail(w, r, id)
	case http.MethodPut:
		updateTask(w, r, id)
	case http.MethodDelete:
		deleteTask(w, r, id)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, APIResponse{Message: "只支持 GET、PUT 或 DELETE 请求"})
	}
}

func getTaskDetail(w http.ResponseWriter, r *http.Request, id int) {

	tasksMu.Lock()
	defer tasksMu.Unlock()

	for _, task := range tasks {
		if task.ID == id {
			writeJSON(w, http.StatusOK, task)
			return
		}
	}
	writeJSON(w, http.StatusNotFound, APIResponse{Message: "任务不存在"})
}

func updateTask(w http.ResponseWriter, r *http.Request, id int) {
	// 这里的逻辑和 getTaskDetail 类似，只是多了读取请求体和更新字段的步骤。
	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, APIResponse{Message: "请求体不是合法 JSON"})
		return
	}

	tasksMu.Lock()
	defer tasksMu.Unlock()

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = req.Done
			writeJSON(w, http.StatusOK, tasks[i])
			return
		}
	}
	writeJSON(w, http.StatusNotFound, APIResponse{Message: "任务不存在"})
}

func deleteTask(w http.ResponseWriter, r *http.Request, id int) {
	tasksMu.Lock()
	defer tasksMu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			writeJSON(w, http.StatusOK, APIResponse{Message: "任务已删除"})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, APIResponse{Message: "任务不存在"})
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("JSON 响应失败:", err)
		http.Error(w, "JSON 响应失败", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(jsonData); err != nil {
		log.Println("写入响应失败:", err)
	}

}
