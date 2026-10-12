package main

import (
	"fmt"
	"sync"
	"time"
)

// ========== Lesson 43: Task 核心实体与状态机流转 ==========
//
// 目标：
// 1. 理解复杂后台任务为什么不能“一次性同步执行”，必须通过【状态机】异步驱动
// 2. 掌握状态常量的枚举定义（iota 与自定义常量）
// 3. 掌握状态迁移的合法性校验（防脏写、防乱序）
// 4. 理解主任务（Task）与子任务/结果项（Result）的状态联动机制
// 5. 对照 issue_api 的 internal/biz/task/task.go

// ----------------------------------------------------
// 1. 模拟 issue_api 的任务状态常量
// ----------------------------------------------------
type TaskStatus int64

const (
	NoStatus    TaskStatus = -1
	ToBeStarted TaskStatus = 0 // 待启动（刚创建时）
	Queuing     TaskStatus = 1 // 排队中
	Running     TaskStatus = 2 // 运行中（正由 Worker 消费处理）
	Success     TaskStatus = 3 // 成功完成
	Fail        TaskStatus = 4 // 执行失败
	Retry       TaskStatus = 5 // 等待重试
	Exit        TaskStatus = 6 // 手动取消/退出
)

// 状态转中文展示
func (s TaskStatus) String() string {
	switch s {
	case ToBeStarted:
		return "待启动 (0)"
	case Queuing:
		return "排队中 (1)"
	case Running:
		return "运行中 (2)"
	case Success:
		return "成功 (3)"
	case Fail:
		return "失败 (4)"
	case Retry:
		return "重试 (5)"
	case Exit:
		return "已取消 (6)"
	default:
		return "未知状态"
	}
}

// ----------------------------------------------------
// 2. 状态机迁移规则校验（防并发脏写、防逆向跃迁）
// ----------------------------------------------------
func canTransition(current, next TaskStatus) bool {
	switch current {
	case ToBeStarted:
		// 待启动 -> 排队中、运行中、已取消
		return next == Queuing || next == Running || next == Exit
	case Queuing:
		// 排队中 -> 运行中、已取消
		return next == Running || next == Exit
	case Running:
		// 运行中 -> 成功、失败、等待重试、已取消
		return next == Success || next == Fail || next == Retry || next == Exit
	case Retry:
		// 等待重试 -> 排队中、运行中、已取消
		return next == Queuing || next == Running || next == Exit
	case Success, Fail, Exit:
		// 终态（Terminal States）：一旦成功、彻底失败或取消，不可再跃迁！
		return false
	default:
		return false
	}
}

// ----------------------------------------------------
// 3. 模拟 Task 主任务与 SubTask 子任务（结果项）
// ----------------------------------------------------
type SubTask struct {
	ID       int64
	Name     string
	Status   TaskStatus
	Finished bool
}

type MainTask struct {
	ID       int64
	Name     string
	Status   TaskStatus
	SubTasks []*SubTask
	mu       sync.Mutex // 状态并发安全锁
}

// 尝试迁移主任务状态
func (m *MainTask) TransitionTo(next TaskStatus) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !canTransition(m.Status, next) {
		return fmt.Errorf("非法状态迁移: 禁止从 [%s] 跃迁到 [%s]", m.Status, next)
	}

	fmt.Printf("[主任务 %d] 状态变更: %s ===> %s\n", m.ID, m.Status, next)
	m.Status = next
	return nil
}

// 子任务完成汇报，并自动计算主任务是否可以标记为 Success
// （对照 issue_api 中的: if findNum == len(results) { fieldsTask["status"] = Success }）
func (m *MainTask) ReportSubTaskFinished(subTaskID int64, success bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	allSuccess := true
	for _, sub := range m.SubTasks {
		if sub.ID == subTaskID {
			if success {
				sub.Status = Success
			} else {
				sub.Status = Fail
			}
			sub.Finished = true
			fmt.Printf("   ├─ 子任务 %s (ID:%d) 结束，状态变为: %s\n", sub.Name, sub.ID, sub.Status)
		}
		if sub.Status != Success {
			allSuccess = false
		}
	}

	// 如果所有子任务全部成功，主任务自动晋升为 Success
	if allSuccess && m.Status == Running {
		fmt.Printf("   └─ [联动触发] 所有子任务均已完成！主任务状态变更为: %s\n", Success)
		m.Status = Success
	}
}

func main() {
	fmt.Println("========== Lesson 43: Task 核心实体与状态机流转 ==========")
	fmt.Println()

	// 1. 初始化一个主任务和两个子任务（如一个追数任务，要查两辆车的包）
	task := &MainTask{
		ID:     101,
		Name:   "ADAS问题数据检索任务_20260921",
		Status: ToBeStarted,
		SubTasks: []*SubTask{
			{ID: 1, Name: "车辆A数据检索", Status: ToBeStarted},
			{ID: 2, Name: "车辆B数据检索", Status: ToBeStarted},
		},
	}
	fmt.Printf("1. 初始任务状态: %s\n\n", task.Status)

	// 2. 模拟合法状态流转：ToBeStarted -> Running
	fmt.Println("2. 工作协程接单，开始执行任务:")
	if err := task.TransitionTo(Running); err != nil {
		fmt.Println("状态迁移失败:", err)
	}
	fmt.Println()

	// 3. 模拟并发非法逆向流转（比如已被取消或成功的任务，某个异步回包试图把它改回 Running）
	fmt.Println("3. 模拟异常：试图把运行中的任务直接倒退回 ToBeStarted（待启动）:")
	if err := task.TransitionTo(ToBeStarted); err != nil {
		fmt.Printf("   [拦截成功] %v\n\n", err)
	}

	// 4. 模拟子任务逐步完成并联动主任务
	fmt.Println("4. 模拟子任务逐步执行完毕:")
	time.Sleep(50 * time.Millisecond)
	task.ReportSubTaskFinished(1, true) // 第一个子任务成功

	time.Sleep(50 * time.Millisecond)
	task.ReportSubTaskFinished(2, true) // 第二个子任务成功，自动触发主任务变 Success
	fmt.Println()

	// 5. 验证终态不可再变
	fmt.Println("5. 验证终态不可再跃迁（任务已经成功，试图将它改为失败）:")
	if err := task.TransitionTo(Fail); err != nil {
		fmt.Printf("   [拦截成功] %v\n", err)
	}
}
