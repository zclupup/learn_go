package main

import (
	"fmt"
	"sync"
	"time"
)

var outsideMessage = "我是 demoChannel 外部的包级变量"

type StudentScore struct {
	Name  string
	Score int
}

type ScoreResult struct {
	Name   string
	Score  int
	Passed bool
	Level  string
}

func getScoreLevel(score int) string {
	if score >= 90 {
		return "优秀"
	}
	if score >= 80 {
		return "良好"
	}
	if score >= 60 {
		return "及格"
	}
	return "不及格"
}

func sayHello() {
	fmt.Println("goroutine: 你好，我是在另一个 goroutine 里执行的")
}

func demoGoroutine() {
	fmt.Println("=== 1. goroutine：让函数并发执行 ===")

	go sayHello()

	fmt.Println("main: 我是主 goroutine，继续往下执行")
	time.Sleep(100 * time.Millisecond)
}

func demoChannel() {
	fmt.Println("\n=== 2. channel：goroutine 之间传数据 ===")

	insideMessage := "我是 demoChannel 内部的局部变量"
	messageChan := make(chan string)

	go func() {
		fmt.Println("匿名函数能看到外部变量:", outsideMessage)
		fmt.Println("匿名函数也能看到局部变量:", insideMessage)
		messageChan <- "这条消息来自另一个 goroutine"
	}()

	message := <-messageChan
	fmt.Println("main 收到:", message)
}

func printNumber(number int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("处理数字:", number)
}

func demoWaitGroup() {
	fmt.Println("\n=== 3. sync.WaitGroup：等待多个 goroutine 完成 ===")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printNumber(i, &wg)
	}

	wg.Wait()
	fmt.Println("所有 goroutine 都执行完了")
}

type SquareResult struct {
	Number int
	Square int
}

func sendSquare(n int, result chan<- SquareResult) {
	square := n * n
	fmt.Printf("数字: %d, 平方: %d\n", n, square)
	result <- SquareResult{Number: n, Square: square}
}

func demoFixedCountReceive() {
	fmt.Println("\n=== 3.1. 已知结果数量：固定接收 5 次 ===")

	result := make(chan SquareResult, 5)

	for i := 1; i <= 5; i++ {
		go sendSquare(i, result)
	}

	results := make([]SquareResult, 0, 5)
	for i := 0; i < 5; i++ {
		value := <-result
		results = append(results, value)
	}

	fmt.Println("固定次数收集到的结果:", results)
}

func demoCloseAndRangeReceive() {
	fmt.Println("\n=== 3.2. 不手动数结果：WaitGroup + close + range ===")

	var wg sync.WaitGroup
	result := make(chan SquareResult, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			sendSquare(n, result)
		}(i)
	}

	wg.Wait()
	close(result)

	results := make([]SquareResult, 0, 5)
	for value := range result {
		results = append(results, value)
	}

	fmt.Println("range 收集到的结果:", results)
}

func printNumberWithCopy(number int, wg sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("错误示范：处理数字:", number)
}

func demoWaitGroupCopyExperiment() {
	fmt.Println("\n=== 4. 小实验：WaitGroup 传值会拷贝，不会改到原来的计数器 ===")

	var wg sync.WaitGroup
	wg.Add(1)

	go printNumberWithCopy(1, wg)

	finished := make(chan bool)
	go func() {
		wg.Wait()
		finished <- true
	}()

	select {
	case <-finished:
		fmt.Println("原来的 wg 归零了")
	case <-time.After(300 * time.Millisecond):
		fmt.Println("等待超时：Done 操作的是副本，原来的 wg 没有归零")
	}
}

// practice
func demoStudentScorePractice() {
	fmt.Println("\n=== 5. practice：并发计算学生成绩 ===")

	students := []StudentScore{
		{Name: "Alice", Score: 85},
		{Name: "Bob", Score: 72},
		{Name: "Charlie", Score: 90},
		{Name: "David", Score: 65},
	}
	resultChan := make(chan ScoreResult, len(students))
	for _, student := range students {
		go func(s StudentScore) {
			resultChan <- ScoreResult{
				Name:   s.Name,
				Score:  s.Score,
				Passed: s.Score >= 60,
				Level:  getScoreLevel(s.Score),
			}
		}(student)

	}
	for i := 0; i < len(students); i++ {
		result := <-resultChan
		fmt.Println("学生成绩结果:", result)
	}

}

func main() {
	demoGoroutine()
	demoChannel()
	demoWaitGroup()
	demoFixedCountReceive()
	demoCloseAndRangeReceive()
	demoWaitGroupCopyExperiment()
	demoStudentScorePractice()
}
