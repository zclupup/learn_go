package main

import "fmt"

func main() {
	// 1) 变量定义：var 变量名 类型 = 值
	var name string = "zhangcl"
	var age int = 23
	var hobby string = "reading"

	// 2) Go 可以自动推断类型（最常用）
	city := "Chengdu"  // string
	height := 172.5    // float64
	isStudent := true  // bool
	hobby2 := "coding" // string

	// 3) 打印变量
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("city:", city)
	fmt.Println("height:", height)
	fmt.Println("isStudent:", isStudent)
	fmt.Println("hobby:", hobby)
	fmt.Println("hobby2:", hobby2)
}
