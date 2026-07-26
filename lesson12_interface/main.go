package main

import "fmt"

// ======== 接口 interface ========
// 接口定义"能做什么"（一组方法），不关心"是谁"。
// 只要某个类型实现了接口里的所有方法，它就"自动"满足这个接口（无需显式声明）。

// 定义一个接口：任何能"发出声音"的东西
type Animal interface {
	Sound() string
}

// ======== 两个结构体，各自实现 Sound 方法 ========
type Dog struct {
	Name string
}

func (d Dog) Sound() string {
	return "汪汪"
}

type Cat struct {
	Name string
}

func (c Cat) Sound() string {
	return "喵喵"
}

// practice
type Cow struct {
	Name string
}

func (c Cow) Sound() string {
	return "哞哞"
}

// ======== 一个函数，接收"接口类型" ========
// 它不关心具体是 Dog 还是 Cat，只要是 Animal 就行
func describe(a Animal) {
	fmt.Println("它发出的声音是:", a.Sound())
}

func main() {
	d := Dog{Name: "旺财"}
	c := Cat{Name: "咪咪"}
	co := Cow{Name: "牛牛"}

	// Dog 和 Cat 都实现了 Sound()，所以都能当作 Animal 使用
	describe(d)
	describe(c)
	describe(co)

	fmt.Println("---")

	// ======== 把不同类型放进同一个接口切片 ========
	animals := []Animal{d, c, co}
	for _, a := range animals {
		fmt.Println("声音:", a.Sound())
	}

	fmt.Println("---")

	// ======== 类型断言：从接口拿回具体类型 ========
	var a Animal = d
	// a.(Dog) 尝试把接口还原成 Dog；ok 表示是否成功
	if dog, ok := a.(Dog); ok {
		fmt.Println("这是一只狗，名字是:", dog.Name)
	}

	a = co
	if cow, ok := a.(Cow); ok {
		fmt.Println("This is a cow, name is:", cow.Name)
	}
}
