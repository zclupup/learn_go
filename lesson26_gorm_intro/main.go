package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"learn_go/lesson26_gorm_intro/biz"
	"learn_go/lesson26_gorm_intro/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ======== Lesson 26：GORM 入门：连接、模型、CRUD ========
//
// 本课把 Lesson25 的 repo 接口思想接到 GORM：
// - model.Task 用 gorm tag 描述字段和表结构
// - data.gormTaskRepo 用 GORM 实现 repo 接口
// - biz.TaskUseCase 仍然只依赖 repo 接口，不直接依赖 GORM
// - main 负责打开数据库、迁移表结构、组装 data 和 biz
//
// 为了让课程不用本地 MySQL 也能跑，这里使用 SQLite 内存数据库。
// 在 issue_api 里，换成 MySQL 连接后，Create/First/Find/Update 这些 GORM 调用模式是一样的。

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := gorm.Open(sqlite.Open("file:lesson26?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}

	if err := data.AutoMigrate(db); err != nil {
		log.Fatal("迁移表结构失败:", err)
	}

	taskRepo := data.NewGormTaskRepo(db)
	taskUseCase := biz.NewTaskUseCase(taskRepo)

	createdTask, err := taskUseCase.CreateTask(ctx, "用 GORM 实现 repo 接口")
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
		fmt.Printf("- ID=%d Title=%s Done=%t CreatedAt=%s\n",
			task.ID,
			task.Title,
			task.Done,
			task.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}
}
