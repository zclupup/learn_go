package main

import (
	"fmt"
)

// ================================================================
// Lesson 29 — Go 泛型（Generics）入门
// ================================================================
//
// 泛型（Generics）是 Go 1.18 引入的特性，核心目的是：减少重复代码。
//
// 场景：你写的多个函数/结构体，逻辑一模一样，只是类型不同。
// 没有泛型时，你要为 int、string、float64 各写一套。
// 有泛型时，写一套，用的时候指定类型即可。
//
// 对照 issue-consumer 中的泛型：
//   issue_tracking_result.go:61 → type HTTPResponse[T any] struct { ... }
//   auth.go:12                → type Result[T any] struct { ... }
//   调用方：httpz.PostJsonBodyFormattedWithCtx[HTTPResponse[HttpBatchImportIssueResp]](...)
//
// ================================================================

func main() {
	fmt.Println("========== Lesson 29: Go 泛型入门 ==========")
	fmt.Println()

	// ============ 1. 没有泛型时的痛点 ============
	fmt.Println("--- 1. 没有泛型：要为每种类型写重复代码 ---")

	fmt.Println(SumInts([]int{1, 2, 3, 4, 5}))
	fmt.Println(SumFloat64s([]float64{1.1, 2.2, 3.3}))
	// 注意：SumInts 和 SumFloat64s 除了类型不同，代码一模一样

	fmt.Println()

	// ============ 2. 泛型函数 ============
	fmt.Println("--- 2. 泛型函数：一套代码，多种类型 ---")

	// 调用泛型函数时，可以显式指定类型参数
	fmt.Println(SumGeneric[int]([]int{1, 2, 3, 4, 5}))
	fmt.Println(SumGeneric[float64]([]float64{1.1, 2.2, 3.3}))

	// 也可以省略类型参数，让编译器自动推断
	fmt.Println(SumGeneric([]int{10, 20, 30}))        // 编译器推断出 T = int
	fmt.Println(SumGeneric([]float64{1.5, 2.5, 3.5})) // 编译器推断出 T = float64

	fmt.Println()

	// ============ 3. 泛型结构体 ============
	fmt.Println("--- 3. 泛型结构体 ---")

	// 创建一个装 int 的盒子
	intBox := NewBox(42)
	fmt.Println(intBox.Get()) // 42

	// 创建一个装 string 的盒子
	strBox := NewBox("hello generics")
	fmt.Println(strBox.Get()) // hello generics

	// 创建一个装 float64 的盒子
	floatBox := NewBox(3.14)
	fmt.Println(floatBox.Get()) // 3.14

	fmt.Println()

	// ============ 4. 对照 issue-consumer 的真实代码 ============
	fmt.Println("--- 4. 对照 issue-consumer 的真实代码 ---")

	// issue-consumer 中的 HTTPResponse[T any] 结构体：
	//   type HTTPResponse[T any] struct {
	//       ErrNo   int64  `json:"err_no"`
	//       ErrMsg  string `json:"err_msg"`
	//       Results T      `json:"results"`
	//   }
	//
	// 调用时：httpz.PostJsonBodyFormattedWithCtx[HTTPResponse[HttpBatchImportIssueResp]](...)
	// 相当于：T = HttpBatchImportIssueResp，所以 Results 字段的类型就是 HttpBatchImportIssueResp

	// 模拟一下：
	loginResp := APIResponse[LoginResult]{
		Code: 0,
		Msg:  "success",
		Data: LoginResult{UserID: 1001, Token: "abc123"},
	}
	fmt.Printf("登录响应: Code=%d, UserID=%d, Token=%s\n", loginResp.Code, loginResp.Data.UserID, loginResp.Data.Token)

	// 同样的 APIResponse 结构体，也可以用于不同的 Data 类型
	articleResp := APIResponse[ArticleData]{
		Code: 0,
		Msg:  "success",
		Data: ArticleData{Title: "Go 泛型入门", Content: "..."},
	}
	fmt.Printf("文章响应: Code=%d, Title=%s\n", articleResp.Code, articleResp.Data.Title)

	fmt.Println()

	// ============ 5. 多个类型参数 ============
	fmt.Println("--- 5. 多个类型参数 ---")

	pair1 := Pair[string, int]{Key: "age", Value: 25}
	pair2 := Pair[string, string]{Key: "name", Value: "zhangcl"}
	fmt.Printf("Pair[string,int]:   Key=%s, Value=%d\n", pair1.Key, pair1.Value)
	fmt.Printf("Pair[string,string]: Key=%s, Value=%s\n", pair2.Key, pair2.Value)

	// 泛型函数也可以有多个类型参数
	fmt.Println(PairToTuple(pair1))
	fmt.Println(PairToTuple(pair2))

	fmt.Println()

	// ============ 6. 类型约束（constraints） ============
	fmt.Println("--- 6. 类型约束 ---")

	// 如果泛型函数里要做比较操作，不能直接用 T any，因为 any 不保证支持 ==
	// 需要加约束：comparable（可比较）或自定义约束
	fmt.Println(Contains([]int{1, 2, 3, 4, 5}, 3))      // true
	fmt.Println(Contains([]int{1, 2, 3, 4, 5}, 10))     // false
	fmt.Println(Contains([]string{"a", "b", "c"}, "b")) // true

	fmt.Println()

	// ============ 7. 自定义类型约束 ============
	fmt.Println("--- 7. 自定义类型约束 ---")

	// 约束一组数字类型，让泛型函数可以同时支持 int/float64
	fmt.Println(Double(5))   // 10
	fmt.Println(Double(3.5)) // 7.0

	fmt.Println()

	// ============ 8. 总结 ============
	fmt.Println("========== 总结 ==========")
	fmt.Println("泛型核心语法：")
	fmt.Println("  函数: func Foo[T any](x T) T        ← [T any] 是类型参数")
	fmt.Println("  结构体: type Box[T any] struct { ... }  ← T 可以在结构体字段中使用")
	fmt.Println("  调用: Foo[int](42) 或省略为 Foo(42)  ← 编译器可推断时省略")
	fmt.Println()
	fmt.Println("对照 issue-consumer:")
	fmt.Println("  HTTPResponse[T any] → T 可以是任意响应类型")
	fmt.Println("  Result[T any]       → T 可以是任意 data 类型")
	fmt.Println("  调用: httpz.PostJsonBodyFormattedWithCtx[HTTPResponse[Xxx]](...)")
	fmt.Println("  这叫做「嵌套泛型」：外层 T = HTTPResponse[Xxx]，内层 Xxx 是具体 response 类型")
	fmt.Println()
	fmt.Println("下一课预告：sync.RWMutex 读写锁 — auth.go 中的 token 缓存并发安全")
	//prcatice
	res := MapSlice([]int{1, 2, 3, 4}, func(x int) string {
		return fmt.Sprintf("Number: %d", x)
	})
	fmt.Println(res)

}

// ============================================================
// 第一部分：没有泛型 — 重复代码
// ============================================================

// SumInts 只能处理 []int。
func SumInts(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// SumFloat64s 只能处理 []float64，和 SumInts 逻辑一模一样，只是类型不同。
func SumFloat64s(nums []float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	return total
}

// 如果还要支持 int64、float32、uint 等类型，每个都要写一遍……

// ============================================================
// 第二部分：泛型函数 — 一套代码，多种类型
// ============================================================

// SumGeneric 是泛型版本的 Sum。
// [T int | float64] 表示 T 可以是 int 或 float64（类型约束）。
// 注意：这里不能用 T any，因为 any 不支持 + 运算符。
func SumGeneric[T int | float64](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// ============================================================
// 第三部分：泛型结构体
// ============================================================

// Box 是一个泛型结构体，可以装任意类型的值。
type Box[T any] struct {
	value T
}

// NewBox 是 Box 的构造函数。
// 注意：泛型结构体的构造函数也要带 [T any]。
func NewBox[T any](value T) Box[T] {
	return Box[T]{value: value}
}

// Get 是 Box 的方法，返回 T 类型的值。
// 注意：方法接收者也要带 [T any]。
func (b Box[T]) Get() T {
	return b.value
}

// ============================================================
// 第四部分：模拟 issue-consumer 的泛型结构体
// ============================================================

// APIResponse 模拟 issue-consumer 的 HTTPResponse[T any]。
// T 是任意类型，Data 字段的类型就是 T。
type APIResponse[T any] struct {
	Code int
	Msg  string
	Data T
}

// LoginResult 是登录接口的返回数据。
type LoginResult struct {
	UserID int
	Token  string
}

// ArticleData 是文章接口的返回数据。
type ArticleData struct {
	Title   string
	Content string
}

// ============================================================
// 第五部分：多个类型参数
// ============================================================

// Pair 有两个类型参数 K 和 V。
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// PairToTuple 把 Pair 转换成两个单独的值。
// 注意：函数也要声明 [K comparable, V any]。
func PairToTuple[K comparable, V any](p Pair[K, V]) (K, V) {
	return p.Key, p.Value
}

// ============================================================
// 第六部分：类型约束 comparable
// ============================================================

// Contains 检查切片中是否包含某个值。
// [T comparable] 表示 T 必须支持 == 和 !=（int、string、bool 等基本类型都支持）。
// 不能用 [T any]，因为 any 不保证支持 ==。
func Contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target { // 这行需要 T 可比较
			return true
		}
	}
	return false
}

// ============================================================
// 第七部分：自定义类型约束
// ============================================================

// Number 自定义一个类型约束，表示"数字类型"。
// 注意：~ 表示"底层类型"也匹配。比如 type MyInt int 也满足 Number 约束。
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Double 把数字翻倍。
// 用了自定义约束 Number，所以可以同时支持所有数字类型。
func Double[T Number](x T) T {
	return x + x
}

// MapSlice applies a function to each element of a slice and returns a new slice with the results.
func MapSlice[T any, U any](items []T, fn func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}
