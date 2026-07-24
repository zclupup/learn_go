package main

import "fmt"
import "math/rand"

func main() {
	// ======== 第一部分：for 循环三种用法 ========

	fmt.Println("=== 1. 计数循环（类似 Python 的 for i in range(3)）===")
	for i := 0; i < 3; i++ {
		fmt.Println("第", i+1, "次循环")
	}

	fmt.Println("\n=== 2. 条件循环（类似 Python 的 while）===")
	count := 0
	for count < 3 {
		fmt.Println("count =", count)
		count++ // 相当于 count = count + 1
	}

	fmt.Println("\n=== 3. 无限循环 + break ===")
	n := 0
	for {
		fmt.Println("n =", n)
		n++
		if n >= 3 {
			break // 跳出循环（类似 Python 的 break）
		}
	}

	// ======== 第二部分：猜数字小游戏 ========
	fmt.Println("\n=== 猜数字小游戏 ===")

	// secret := 7
	secret := rand.Intn(10) + 1 // 生成 1-10 的随机整数
	var guess int
	attempts := 0

	for {
		fmt.Print("猜一个数字（1-10）: ")
		fmt.Scan(&guess)
		attempts++

		if guess < secret {
			fmt.Println("太小了，再试试")
		} else if guess > secret {
			fmt.Println("太大了，再试试")
		} else {
			fmt.Println("猜对了！你真棒！")
			break // 猜对就退出循环
		}

		if attempts >= 3 {
			fmt.Println("机会用完了，正确答案是:", secret)
			break
		}
	}
}
