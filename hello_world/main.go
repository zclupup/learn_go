package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 模拟一个业务goroutine
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(1 * time.Second):
				fmt.Println("业务运行中...")
			}
		}
	}(ctx)

	// 信号监听协程
	go func() {
		<-sigChan
		fmt.Println("\n后台协程：捕获到操作系统退出信号！准备调用cancel()")
		cancel()
	}()

	fmt.Println("程序启动，按 Ctrl+C 退出")
	// 主协程等待取消
	<-ctx.Done()
	fmt.Println("主goroutine收到ctx.Done，开始优雅关闭，err:", ctx.Err())
	time.Sleep(2 * time.Second)
	fmt.Println("程序完全退出")
}
