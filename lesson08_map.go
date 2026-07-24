package main

import "fmt"

func main() {
	// ======== map = Go 版的字典（类似 Python 的 dict）========
	// 格式：map[键的类型]值的类型

	// 1) 创建 map（key 是 string，value 是 int）
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Cathy": 28,
	}
	fmt.Println("整个 map:", ages)

	// 2) 取值（用中括号，类似 Python 的 dict[key]）
	fmt.Println("Alice 的年龄:", ages["Alice"])

	// 3) 新增 / 修改（写法一样：存在就改，不存在就新增）
	ages["David"] = 40 // 新增
	ages["Bob"] = 26   // 修改
	fmt.Println("修改后:", ages)

	// 4) 删除某个键（Python 用 del，Go 用 delete 函数）
	delete(ages, "Cathy")
	fmt.Println("删除 Cathy 后:", ages)

	// 5) 判断键是否存在（Go 特有的 "逗号 ok" 写法）⭐
	// 取值时可以同时拿到第二个返回值 ok（true=存在，false=不存在）
	age, ok := ages["Alice"]
	if ok {
		fmt.Println("Alice 存在，年龄是:", age)
	} else {
		fmt.Println("Alice 不存在")
	}

	value, exists := ages["Cathy"]
	fmt.Println("Cathy 的值:", value, " 是否存在:", exists) // value=0（int零值）exists=false

	// 6) 遍历 map（类似 Python 的 for k, v in dict.items()）
	fmt.Println("\n=== 遍历所有人 ===")
	for name, age := range ages {
		fmt.Printf("%s 的年龄是 %d\n", name, age)
	}

	// 7) map 的长度
	fmt.Println("\n一共有", len(ages), "个人")

	// prictise
	capitals := map[string]string{
		"China":   "Beijing",
		"America": "NewYork",
		"Japan":   "Tokyo",
	}

	fmt.Println("the whole map:", capitals)
	capitals["France"] = "Paris"
	fmt.Println("after add France:", capitals)
	capital, ok := capitals["Japan"]
	if ok {
		fmt.Println("Japan exists, capital is:", capital)
	} else {
		fmt.Println("Japan does not exist")
	}
	for country, capital := range capitals {
		fmt.Printf("%s's capital is %s\n", country, capital)
	}
}
