package main

import (
	"errors"
	"fmt"
)

// ======== Go 的错误处理核心思想 ========
// Python 用 try/except 抛异常；
// Go 不用异常，而是把 error 当成"函数的一个返回值"返回给你，
// 你自己检查 "if err != nil" 来判断有没有出错。

// 1) 一个可能出错的函数：除法（除数为 0 就返回错误）
// 约定：error 永远是"最后一个返回值"
func divide(a, b int) (int, error) {
	if b == 0 {
		// 用 errors.New 创建一个错误
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil // 没出错就返回 nil（表示"没有错误"）
}

// 2) 用 fmt.Errorf 创建带变量的错误信息（类似 Printf 的格式化）
func checkAge(age int) error {
	if age < 0 {
		return fmt.Errorf("年龄不能是负数, 你输入的是 %d", age)
	}
	if age > 150 {
		return fmt.Errorf("年龄 %d 太大了, 不合理", age)
	}
	return nil // 检查通过，没有错误
}

// practice
func getScore(scores []int, index int) (int, error) {
	if index < 0 || index >= len(scores) {
		return 0, fmt.Errorf("index %d exceeds the range of scores (0-%d)", index, len(scores)-1)
	}
	return scores[index], nil
}

func main() {
	// ======== 处理返回 error 的函数 ========

	// 情况1：正常
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("出错了:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// 情况2：出错（除数为 0）
	result2, err := divide(10, 0)
	if err != nil {
		fmt.Println("出错了:", err) // 这里会打印
	} else {
		fmt.Println("10 / 0 =", result2)
	}

	fmt.Println("---")

	// ======== 只关心有没有出错的函数 ========
	ages := []int{25, -5, 200}
	for _, age := range ages {
		err := checkAge(age)
		if err != nil {
			fmt.Println("年龄检查失败:", err)
		} else {
			fmt.Printf("年龄 %d 合法\n", age)
		}
	}

	// practice
	scores := []int{90, 85, 78}
	result3, err := getScore(scores, 1)
	if err != nil {
		fmt.Println("wrong:", err)
	} else {
		fmt.Println("score at index 1:", result3)
	}

	result4, err := getScore(scores, 5)
	if err != nil {
		fmt.Println("wrong:", err)
	} else {
		fmt.Println("score at index 5:", result4)
	}
}
