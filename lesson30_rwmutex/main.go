package main

import (
	"fmt"
	"sync"
	"time"
)

// ================================================================
// Lesson 30 — sync.RWMutex 读写锁
// ================================================================
//
// 回顾 Lesson 21：sync.Mutex 是互斥锁，Lock() 之后所有人都得等。
// 读操作和读操作之间并不冲突，完全没必要排队。
// 但 Mutex 不管：读也 Lock，写也 Lock，一个在锁里，其他人全在外面等。
//
// RWMutex 就是解决这个问题的：
//   - RLock() / RUnlock()：读锁，多个 goroutine 可以同时持有
//   - Lock() / Unlock()：写锁，独占，持有写锁时不能读也不能写
//
// 规则：
//   读锁互斥写锁，但不互斥读锁（可以同时读）
//   写锁互斥所有锁（独占）
//
// 对照 issue-consumer 中 auth.go 的真实场景：
//   - 多个 Kafka 消费者同时需要 token
//   - token 有缓存，大多数时候是读缓存
//   - 只有 token 过期时才需要写缓存（重新获取）
//   - RWMutex 是最优选择：读不互斥，写才互斥
//
// ================================================================

func main() {
	fmt.Println("========== Lesson 30: sync.RWMutex 读写锁 ==========")
	fmt.Println()

	// ============ 1. 用 Mutex 做缓存（反例）============
	fmt.Println("--- 1. 用 Mutex 做缓存：读也互斥 ---")
	fmt.Println("Mutex 读缓存：每个读都要 Lock，同时只能一个 goroutine 读")
	fmt.Println("问题：读操作之间不冲突，但被锁串行化了，浪费性能")
	fmt.Println()

	// ============ 2. RWMutex 基本用法 ============
	fmt.Println("--- 2. RWMutex 基本用法 ---")

	var rw sync.RWMutex
	counter := 0

	// 写操作：Lock + Unlock（独占）
	go func() {
		rw.Lock()
		counter = 42
		fmt.Println("写：counter =", counter)
		rw.Unlock()
	}()

	time.Sleep(10 * time.Millisecond)

	// 读操作：RLock + RUnlock（可以多个同时读）
	rw.RLock()
	fmt.Println("读：counter =", counter)
	rw.RUnlock()

	fmt.Println()

	// ============ 3. 模拟 issue-consumer 的 token 缓存 ============
	fmt.Println("--- 3. 模拟 issue-consumer 的 token 缓存 ---")

	// 问题：多个 goroutine 同时需要 token，token 有缓存
	// 大多数时候从缓存读（快），偶尔过期需要重新获取（慢）
	// 读多写少场景，RWMutex 最合适

	tokenCache := &TokenCache{
		rw: new(sync.RWMutex),
	}

	var wg sync.WaitGroup

	// 模拟 5 个 goroutine 同时需要 token
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			token, err := tokenCache.GetToken(id)
			if err != nil {
				fmt.Printf("goroutine %d: 获取 token 失败: %v\n", id, err)
			} else {
				fmt.Printf("goroutine %d: 拿到 token: %s\n", id, token)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println()

	// ============ 4. 对比 Mutex vs RWMutex 的性能差异 ============
	fmt.Println("--- 4. 读多写少场景：Mutex vs RWMutex ---")

	// 场景：1000 次读，1 次写
	readCount := 1000

	// 用 Mutex
	start := time.Now()
	benchMutex(readCount)
	fmt.Printf("Mutex 耗时:   %v\n", time.Since(start))

	// 用 RWMutex
	start = time.Now()
	benchRWMutex(readCount)
	fmt.Printf("RWMutex 耗时: %v\n", time.Since(start))

	fmt.Println()

	// ============ 5. RWMutex 的规则总结 ============
	fmt.Println("--- 5. RWMutex 规则总结 ---")
	fmt.Println("  写Lock     读RLock   |  结果")
	fmt.Println("  ─────────────────────┼──────────")
	fmt.Println("  持有中      等待       |  写独占，读要等")
	fmt.Println("  等待        持有中     |  读先执行，写要等")
	fmt.Println("  等待        等待       |  写优先（Go 保证写者不会饿死）")
	fmt.Println("  不持有      不持有     |  各走各的")
	fmt.Println()

	// ============ 6. 对照 issue-consumer auth.go ============
	fmt.Println("--- 6. 对照 issue-consumer auth.go ---")
	fmt.Println("auth.go 中的 token 缓存：")
	fmt.Println("  var rwmux sync.RWMutex")
	fmt.Println("  var cacheToken = \"\"")
	fmt.Println()
	fmt.Println("  读缓存：rwmux.RLock() → 读 cacheToken → rwmux.RUnlock()")
	fmt.Println("  写缓存：rwmux.Lock()   → 写 cacheToken → rwmux.Unlock()")
	fmt.Println()
	fmt.Println("  为什么用 RWMutex 而不是 Mutex？")
	fmt.Println("  因为 token 缓存是典型的「读多写少」：")
	fmt.Println("  - 每次发 HTTP 请求都要读 token（高频）")
	fmt.Println("  - 只有 token 过期才写（低频，可能几小时一次）")
	fmt.Println("  - 用 Mutex 的话，大量读操作互相阻塞，白白浪费性能")
	fmt.Println()

	fmt.Println("========== 总结 ==========")
	fmt.Println("| 锁类型    | 读操作          | 写操作        | 适用场景     |")
	fmt.Println("|----------|----------------|--------------|------------|")
	fmt.Println("| Mutex    | Lock（互斥）    | Lock（互斥）  | 读写都改数据 |")
	fmt.Println("| RWMutex  | RLock（可并发） | Lock（独占）  | 读多写少     |")
	fmt.Println()
	fmt.Println("记忆口诀：看操作会不会「改数据」")
	fmt.Println("  只读不出 → RLock（大家一起读，互不影响）")
	fmt.Println("  要改数据 → Lock（一个人改，其他人全等着）")
}

// ============================================================
// 模拟 issue-consumer auth.go 的 token 缓存
// ============================================================

// TokenCache 用 RWMutex 保护缓存的 token。
// 这和 auth.go 中的 rwmux + cacheToken 模式完全一致。
type TokenCache struct {
	rw    *sync.RWMutex
	token string
}

// GetToken 获取 token，缓存命中直接返回，没命中则重新获取。
func (tc *TokenCache) GetToken(id int) (string, error) {
	// 第 1 步：用读锁检查缓存（多个 goroutine 可以同时进入）
	tc.rw.RLock()
	if tc.token != "" {
		token := tc.token
		tc.rw.RUnlock()
		fmt.Printf("goroutine %d: 缓存命中!\n", id)
		return token, nil
	}
	tc.rw.RUnlock() // 释放读锁，准备获取写锁

	// 第 2 步：缓存为空，用写锁重新获取
	tc.rw.Lock()
	defer tc.rw.Unlock()

	// 双重检查：可能刚才别人已经写入了
	if tc.token != "" {
		return tc.token, nil
	}

	// 模拟从外部获取 token（如 HTTP 请求 auth API）
	fmt.Printf("goroutine %d: 缓存未命中，重新获取 token...\n", id)
	time.Sleep(50 * time.Millisecond) // 模拟网络延迟
	tc.token = fmt.Sprintf("token_%d", time.Now().UnixMilli())
	return tc.token, nil
}

// ============================================================
// 性能对比
// ============================================================

// benchMutex 用 Mutex 保护读写。
func benchMutex(readCount int) {
	var mu sync.Mutex
	counter := 0

	// 写操作
	mu.Lock()
	counter = 42
	mu.Unlock()

	// 1000 次读操作，每次都要 Lock（互斥！）
	var wg sync.WaitGroup
	for i := 0; i < readCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			_ = counter
			mu.Unlock()
		}()
	}
	wg.Wait()
}

// benchRWMutex 用 RWMutex 保护读写。
func benchRWMutex(readCount int) {
	var rw sync.RWMutex
	counter := 0

	// 写操作
	rw.Lock()
	counter = 42
	rw.Unlock()

	// 1000 次读操作，用 RLock（可以并发！）
	var wg sync.WaitGroup
	for i := 0; i < readCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rw.RLock()
			_ = counter
			rw.RUnlock()
		}()
	}
	wg.Wait()
}