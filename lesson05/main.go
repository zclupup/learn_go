package main

import "fmt"

func main() {
	// ======== 数组（长度固定，类似 Python 的 tuple? 不太像，先放着）========

	// Go 的数组长度是固定的，一旦定义不能变
	var arr [3]int = [3]int{10, 20, 30}
	fmt.Println("数组:", arr)
	fmt.Println("第1个元素:", arr[0])  // 索引从 0 开始
	fmt.Println("数组长度:", len(arr)) // len() 获取长度

	// ======== 切片（重点！类似 Python 的 list）========
	// 切片就是"可变长度的数组"，入门阶段你把它当 Python 的 list 用就行

	// 1) 创建切片
	fruits := []string{"苹果", "香蕉", "橘子"}
	fmt.Println("\n水果列表:", fruits)
	fmt.Println("长度:", len(fruits))

	// 2) 追加元素（类似 Python 的 list.append()）
	fruits = append(fruits, "葡萄")
	fmt.Println("追加后:", fruits)

	// 3) 用 for 循环遍历（类似 Python 的 for item in list）
	fmt.Println("\n遍历水果:")
	for i, fruit := range fruits {
		fmt.Printf("第%d个水果是: %s\n", i+1, fruit)
	}

	// 4) 切片操作（类似 Python 的 list[1:3]）
	fmt.Println("\n切片操作:")
	numbers := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("numbers[1:4]:", numbers[1:4]) // [1, 2, 3]  含头不含尾
	fmt.Println("numbers[:3]:", numbers[:3])   // [0, 1, 2]
	fmt.Println("numbers[3:]:", numbers[3:])   // [3, 4, 5]

	// practice: 练习题
	scores := []int{85, 92, 78}
	scores = append(scores, 95)
	sumScores := 0
	for i, score := range scores {
		fmt.Printf("第%d个成绩是:%d\n", i+1, score)
		sumScores += score
	}
	fmt.Println("总成绩:", sumScores)
}
