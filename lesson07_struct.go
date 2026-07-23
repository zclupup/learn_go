package main

import "fmt"

// ======== 定义结构体（类似 Python 的 class）========
// 把"一个人"的相关信息打包在一起
type Person struct {
	Name string
	Age  int
	City string
}

// ======== 给结构体绑定方法（类似 Python 的类方法）========
// (p Person) 表示这个方法属于 Person，p 相当于 Python 的 self
func (p Person) Introduce() string {
	return fmt.Sprintf("我叫%s，今年%d岁，来自%s", p.Name, p.Age, p.City)
}

// definion struct Book
type Book struct {
	Title string
	Author string
	Price float64
}

func (b Book) Info() string {
	return fmt.Sprintf("book name: %s, author: %s, price: %.2f", b.Title, b.Author, b.Price)
}


func main() {
	// ======== 1) 创建结构体（写字段名，推荐）========
	p1 := Person{
		Name: "zhangcl",
		Age:  23,
		City: "Chengdu",
	}
	fmt.Println("p1:", p1)

	// ======== 2) 访问字段（用点 . ）========
	fmt.Println("名字:", p1.Name)
	fmt.Println("年龄:", p1.Age)

	// ======== 3) 修改字段 ========
	p1.Age = 24
	fmt.Println("修改后的年龄:", p1.Age)

	// ======== 4) 调用方法 ========
	fmt.Println(p1.Introduce())

	// ======== 5) 创建多个（放进切片里，就像一个通讯录）========
	people := []Person{
		{Name: "Alice", Age: 30, City: "Beijing"},
		{Name: "Bob", Age: 25, City: "Shanghai"},
	}


	fmt.Println("\n=== 通讯录 ===")
	for _, p := range people {
		fmt.Println(p.Introduce())
	}

	// ========6) prictice create two book struct
	books := []Book{
		{Title: "study go", Author: "zhangcl", Price: 28.61},
		{Title: "review go", Author: "ca", Price: 20.00},
	}

	fmt.Println("\n====books====")
	for _, b := range books {
		fmt.Println(b.Info())
	}
}
