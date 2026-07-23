package main

import "fmt"

func main() {
	var age int

	fmt.Print("请输入你的年龄: ")
	fmt.Scan(&age)

	if age >= 60 {
		fmt.Println("你是老年人")
	} else if age >= 18 {
		fmt.Println("你是成年人")
	} else {
		fmt.Println("你还未成年")
	}
}
