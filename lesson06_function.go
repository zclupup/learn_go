package main

import "fmt"

// ======== 函数定义 ========

// 1) 无返回值函数（类似 Python 的 def，但没有 return）
func sayHello(name string) {
	fmt.Println("你好,", name)
}

// 2) 有返回值函数（类似 Python 的 def ... return）
func add(a int, b int) int {
	return a + b
}

// 3) 多个参数可以简写类型
func multiply(a, b int) int {
	return a * b
}

// 4) 多个返回值（Python 也能返回多个值，Go 也支持）
func divide(a, b int) (int, int) {
	quo := a / b
	rem := a % b
	return quo, rem
}

// practice
func greet(name string, age int) string {
	return fmt.Sprintf("my name is %s, I am %d years old", name, age)
}

func average(nums []int) (int, float64) {
	sum := 0
	for _, num := range nums{
		sum += num
	}
	avg := float64(sum) / float64(len(nums))
	return sum, avg
}

func main() {
	// 调用函数
	sayHello("zhangcl")

	result := add(10, 20)
	fmt.Println("10 + 20 =", result)

	product := multiply(6, 7)
	fmt.Println("6 * 7 =", product)

	// 接收多个返回值
	q, r := divide(17, 5)
	fmt.Println("17 / 5 =", q, "余", r)

	// 如果只想要其中一个返回值，用 _ 忽略
	_, remainder := divide(20, 3)
	fmt.Println("20 / 3 的余数是:", remainder)

	// practice
	greeting := greet("Alice", 30)
	fmt.Println(greeting)

	nums := []int{10, 20, 31, 47, 54}
	total, avg := average(nums)
	fmt.Println("总和:", total, "平均值:", avg)
}
