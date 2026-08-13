package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ======== Lesson 20：标准库 net/http 入门 ========
//
// Go 标准库自带 net/http，不用第三方框架也能启动 HTTP 服务。
// 这一课先学最小后端接口：
// - http.HandleFunc(path, handler) 注册路由
// - http.ListenAndServe(addr, nil) 启动服务
// - http.ResponseWriter 写响应
// - *http.Request 读取请求信息
// - 注册 "/" 会兜底匹配没有更具体 handler 的路径；想返回 404 要自己判断 r.URL.Path。
//
// 常用调试命令：
// - curl -s http://localhost:8928/products
// - curl -s -X POST http://localhost:8928/products
// - ss -ltnp | grep :8928        → 查看端口是否被占用
// - Ctrl+C                       → 停止当前前台运行的服务
//
// Python 对比：
// - http.HandleFunc 类似 Flask/FastAPI 里注册路由
// - handler 类似视图函数
// - r.URL.Query().Get("name") 类似 request.args.get("name")
// - json.NewEncoder(w).Encode(data) 类似 return jsonify(data)

type APIResponse struct {
	Message string `json:"message"`
	Path    string `json:"path,omitempty"`
	Time    string `json:"time,omitempty"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/bad-json", badJSONHandler)

	addr := ":8928"
	fmt.Println("HTTP 服务启动：http://localhost" + addr)
	fmt.Println("访问示例：http://localhost:8928/hello?name=张三")

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		writeJSON(w, http.StatusNotFound, APIResponse{Message: "接口不存在"})
		return
	}

	// Fprintln 写到指定目标 w；在 HTTP 里写到 w 就是返回给客户端。
	// Println 则是写到服务端终端。
	fmt.Fprintln(w, "欢迎学习 Go 标准库 net/http")
}

func badJSONHandler(w http.ResponseWriter, r *http.Request) {
	// chan int 不能被 JSON 编码，用来演示 json.NewEncoder(w).Encode(data) 返回错误。
	// 这个错误会打印到服务端终端，而不是 curl 响应里。
	writeJSON(w, http.StatusOK, make(chan int))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Go 学习者"
	}

	response := APIResponse{
		Message: "你好，" + name,
		Path:    r.URL.Path,
	}

	writeJSON(w, http.StatusOK, response)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	response := APIResponse{
		Message: "当前服务器时间",
		// Go 用固定参考时间 2006-01-02 15:04:05 表示 年-月-日 时:分:秒。
		// 这些数字有特殊含义，分隔符和排列方式可以改，但不能写成 YYYY-MM-DD。
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}

	writeJSON(w, http.StatusOK, response)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, APIResponse{Message: "只支持 GET 请求"})
		// 已经写了错误响应，必须 return，避免继续执行后面的正常响应逻辑。
		return
	}

	users := []User{
		{ID: 1, Name: "张三"},
		{ID: 2, Name: "李四"},
		{ID: 3, Name: "王五"},
	}

	writeJSON(w, http.StatusOK, users)
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	// NewEncoder(w).Encode(data) 会把 data 编成 JSON，并直接写入 w。
	// 和 json.Marshal 不同，它不会把 []byte 返回给你。
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("JSON 响应失败:", err)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, APIResponse{Message: "只支持 GET 请求"})
		// 提前结束 handler，避免同一次请求又写 405 又继续写 200 商品列表。
		return
	}
	products := []Product{
		{ID: 1, Name: "笔记本电脑", Price: 5999.99},
		{ID: 2, Name: "智能手机", Price: 3999.99},
		{ID: 3, Name: "无线耳机", Price: 499.99},
	}
	writeJSON(w, http.StatusOK, products)
}
