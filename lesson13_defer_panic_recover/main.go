package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// ======== defer：延迟执行 ========
// defer 后面的语句会"延迟"到函数即将返回时才执行。
// 常用于收尾工作：关闭文件、释放资源。类似 Python 的 finally / with。

func deferDemo() {
	fmt.Println("1. 函数开始")
	defer fmt.Println("3. defer 执行（函数结束前）") // 被推迟到最后
	fmt.Println("2. 函数主体")
}

// 多个 defer：后进先出（栈）。最后写的 defer 最先执行。
func multiDefer() {
	defer fmt.Println("defer A（最先写，最后执行）")
	defer fmt.Println("defer B")
	defer fmt.Println("defer C（最后写，最先执行）")
	fmt.Println("函数主体先执行")
}

// ======== panic：主动"崩溃" ========
// panic 会让程序立即停止当前流程（类似 Python 抛出未捕获的异常）。

// ======== recover：从 panic 中恢复 ========
// recover 只能在 defer 里生效，能"抓住" panic，让程序不崩溃继续运行。
func safeDivide(a, b int) (result int, err error) {
	// 用 defer + recover 兜底：万一发生 panic，就把它变成 error 返回
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("发生了 panic: %v", r)
		}
	}()

	result = a / b // 如果 b 是 0，这里会 panic
	return result, nil
}

// practice safeGet
func safeGet(nums []int, index int) (val int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("occur panic: %v", r)
		}
	}()
	val = nums[index] // 可能会 panic
	return val, nil
}

func main() {
	fmt.Println("=== 1. defer 基础 ===")
	deferDemo()

	fmt.Println("\n=== 2. 多个 defer（后进先出）===")
	multiDefer()

	fmt.Println("\n=== 3. panic + recover 兜底 ===")
	// 正常情况
	r1, err1 := safeDivide(10, 2)
	fmt.Println("10 / 2 =", r1, "err =", err1)

	// 除以 0：本会 panic，但被 recover 抓住变成 error
	r2, err2 := safeDivide(10, 0)
	fmt.Println("10 / 0 =", r2, "err =", err2)

	fmt.Println("\n程序没有崩溃，继续正常结束 ✅")

	fmt.Println("\n=== 4. safeGet ===")
	nums := []int{1, 2, 3}
	val, err := safeGet(nums, 1)
	fmt.Println("safeGet(nums, 1) =", val, "err =", err)

	val, err = safeGet(nums, 5)
	fmt.Println("safeGet(nums, 5) =", val, "err =", err)

	fmt.Println("\n=== 5. %v / %+v / %T ===")
	person := Person{Name: "Alice", Age: 18}
	fmt.Printf("%%v  -> %v\n", person)
	fmt.Printf("%%+v -> %+v\n", person)
	fmt.Printf("%%T  -> %T\n", person)
}
