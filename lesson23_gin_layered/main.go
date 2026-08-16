package main

import (
	"log"

	"learn_go/lesson23_gin_layered/handler"
	"learn_go/lesson23_gin_layered/router"
	"learn_go/lesson23_gin_layered/service"
)

// ======== Lesson 23：Gin 分层小项目 ========
//
// 这一课把 Lesson 22 的单文件 Gin 示例拆成更接近真实项目的分层结构：
// - model：数据结构，请求/响应结构体
// - service：业务逻辑，管理数据和并发锁
// - handler：HTTP 层，读取参数和请求体，返回 JSON
// - router：路由注册，集中管理 URL 和 handler 的对应关系
// - main：组装依赖并启动服务
//
// 请求流程：
// client -> router -> handler -> service -> model/data

func main() {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)
	r := router.SetupRouter(userHandler)

	addr := ":8928"
	log.Println("Lesson 23 Gin 分层小项目启动：http://localhost" + addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
