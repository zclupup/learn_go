package main

import "fmt"

func main() {
	name := "zhangcl"
	age := 23
	city := "Chengdu"
	height := 172.5
	hobby := "coding"

	// 对比 Python 的 f-string:
	// Python:   print(f"我叫{name}，今年{age}岁")
	//
	// Go 的 fmt.Printf 类似但写法不同：
	// %s = 字符串, %d = 整数, %f = 小数

	fmt.Println("我叫%s，今年%d岁，来自%s", name, age, city)

	// 保留1位小数: %.1f
	fmt.Printf("身高: %.1f cm\n", height)

	fmt.Printf("我叫%s, 今年%d岁，喜欢%s\n", name, age, hobby)

	// 也可以拼接多行:
	info := fmt.Sprintf("姓名: %s, 年龄: %d", name, age)
	fmt.Println(info)
}
