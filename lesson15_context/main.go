package main

import (
	"context"
	"fmt"
	"time"
)

// ======== Lesson 15：context 上下文与超时控制 ========
//
// context 可以理解成 Go 里的"任务控制器"。
// 它常用来告诉 goroutine：任务超时了、请求取消了，你该停下来了。
//
// 核心 API：
// - context.Background()  → 最基础的空 context，作为起点
// - context.WithTimeout(parent, duration) → 返回 (新ctx, cancel函数)，到时间自动取消
// - context.WithCancel(parent)            → 返回 (新ctx, cancel函数)，需手动调 cancel() 取消
// - ctx.Done()  → 返回 channel，取消/超时时收到信号
// - ctx.Err()   → 返回取消原因："context deadline exceeded"(超时) 或 "context canceled"(手动取消)
//
// ⭐ 重要：
// - WithTimeout 返回两个值，必须都接收：ctx, cancel := context.WithTimeout(...)
// - 第一个参数是"父 context"，可以用前面声明的 ctx，也可以直接写 context.Background()
// - WithTimeout 到时间自动触发 ctx.Done()；WithCancel 必须手动 cancel() 才触发
// - 固定写法：defer cancel() —— 不管成功失败都释放资源
//
// Python 对比：
// - 像给任务设置 timeout
// - 也像传一个"取消信号"给后台线程/协程

// slowTask 模拟一个耗时任务。
// 它会同时等待两件事：
// 1) 任务自己做完
// 2) context 发来取消/超时信号
func slowTask(ctx context.Context, name string, workTime time.Duration) error {
	fmt.Println(name, "开始执行")

	select {
	case <-time.After(workTime):
		// time.After(workTime) 会在指定时间后往 channel 里发一个值。
		fmt.Println(name, "执行完成")
		return nil
	case <-ctx.Done():
		// ctx.Done() 返回一个 channel。
		// 如果 context 被取消或超时，这个 channel 就会收到信号。
		// ctx.Err() 会告诉你原因：
		//   - "context deadline exceeded" → WithTimeout 超时
		//   - "context canceled"          → WithCancel 手动取消
		return ctx.Err()
	}
}

type ServiceResult struct {
	Name string
	Data string
	Err  error
}

// practice: queryService 模拟请求某个服务
// 返回 ServiceResult，包含服务名、数据、错误信息
func queryService(ctx context.Context, name string, delay time.Duration) ServiceResult {
	fmt.Printf("%s 开始请求\n", name)

	select {
	case <-time.After(delay):
		// 任务在超时前完成
		return ServiceResult{Name: name, Data: "data from " + name, Err: nil}
	case <-ctx.Done():
		// context 超时或被取消
		return ServiceResult{Name: name, Data: "", Err: ctx.Err()}
	}
}

func main() {
	fmt.Println("=== 1. Background：创建最基础的 context ===")
	ctx := context.Background()
	fmt.Println("ctx =", ctx)

	fmt.Println("\n=== 2. WithTimeout：任务在超时前完成 ===")
	// WithTimeout 返回两个值：
	// 1) fastCtx：新的 context，带超时功能
	// 2) fastCancel：取消函数，必须调用以释放资源
	// 第一个参数 context.Background() 是"父 context"
	fastCtx, fastCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer fastCancel() // 固定写法：不管成功失败都 defer cancel()

	err := slowTask(fastCtx, "快速任务", 1*time.Second)
	if err != nil {
		fmt.Println("快速任务失败:", err)
	} else {
		fmt.Println("快速任务成功")
	}

	fmt.Println("\n=== 3. WithTimeout：任务超时被取消 ===")
	// 这里超时时间 1s < 任务耗时 3s，所以会自动触发 ctx.Done()
	// ⭐ WithTimeout 到时间自动触发，不需要手动 cancel()
	slowCtx, slowCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer slowCancel()

	err = slowTask(slowCtx, "慢速任务", 3*time.Second)
	if err != nil {
		fmt.Println("慢速任务失败:", err)
	} else {
		fmt.Println("慢速任务成功")
	}

	fmt.Println("\n=== 4. WithCancel：手动取消任务 ===")
	// WithCancel 只有手动调用 cancel() 才会触发 ctx.Done()
	// ⭐ 和 WithTimeout 的区别：WithTimeout 到时间自动触发，WithCancel 必须手动触发
	cancelCtx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("准备手动取消任务")
		cancel() // 手动调用 cancel()，此时 ctx.Done() 才会收到信号
	}()

	err = slowTask(cancelCtx, "可取消任务", 2*time.Second)
	if err != nil {
		fmt.Println("可取消任务失败:", err)
	} else {
		fmt.Println("可取消任务成功")
	}

	fmt.Println("\n=== 5. 练习：并发请求多服务（带超时控制）===")
	// 设置 2 秒总超时
	practiceCtx, practiceCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer practiceCancel()

	// 定义 3 个服务的耗时
	services := []struct {
		name  string
		delay time.Duration
	}{
		{"服务A", 500 * time.Millisecond},
		{"服务B", 1500 * time.Millisecond},
		{"服务C", 3000 * time.Millisecond}, // 会超时
	}

	resChan := make(chan ServiceResult, len(services))

	// 并发请求所有服务
	for _, svc := range services {
		go func(name string, delay time.Duration) {
			result := queryService(practiceCtx, name, delay)
			resChan <- result
		}(svc.name, svc.delay) // ⭐ 把变量作为参数传进去，避免闭包陷阱
	}

	// 收集所有结果
	for i := 0; i < len(services); i++ {
		result := <-resChan
		if result.Err != nil {
			fmt.Printf("%s: 失败, %v\n", result.Name, result.Err)
		} else {
			fmt.Printf("%s: 成功, %s\n", result.Name, result.Data)
		}
	}
}
