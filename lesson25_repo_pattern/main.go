package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"learn_go/lesson25_repo_pattern/biz"
	"learn_go/lesson25_repo_pattern/data"
	"learn_go/lesson25_repo_pattern/model"
)

// ======== Lesson 25：interface + repo 模式 ========
//
// 这一课的重点不是数据库，而是“依赖方向”：
// - repo 包定义接口：biz 层需要什么数据能力
// - data 包实现接口：数据到底从哪里来
// - biz 包只依赖 repo 接口，不直接依赖具体 data 实现
// - main 负责把 data 实现组装进 biz
//
// 对照 issue_api：
// internal/biz/repo/pack.go 定义 PackRepo 接口；
// internal/data/pack.go 里的 packRepo 实现这个接口；
// biz/usecase 拿到的是 repo.PackRepo，而不是具体的 data.packRepo。

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	taskRepo := data.NewMemoryTaskRepo()
	taskUseCase := biz.NewTaskUseCase(taskRepo)

	createdTask, err := taskUseCase.CreateTask(ctx, model.CreateTaskRequest{Title: "对照 issue_api 的 repo 接口"})
	if err != nil {
		log.Fatal("创建任务失败:", err)
	}
	fmt.Printf("创建任务：ID=%d Title=%s Done=%t\n", createdTask.ID, createdTask.Title, createdTask.Done)

	finishedTask, err := taskUseCase.FinishTask(ctx, createdTask.ID)
	if err != nil {
		log.Fatal("完成任务失败:", err)
	}
	fmt.Printf("完成任务：ID=%d Title=%s Done=%t\n", finishedTask.ID, finishedTask.Title, finishedTask.Done)

	tasks, err := taskUseCase.ListTasks(ctx)
	if err != nil {
		log.Fatal("查询任务失败:", err)
	}

	fmt.Println("任务列表：")
	for _, task := range tasks {
		fmt.Printf("- ID=%d Title=%s Done=%t\n", task.ID, task.Title, task.Done)
	}
}
