package main

import "fmt"

// ======== 用指针让函数修改"原始变量" ========
// 传普通值：函数拿到的是"复印件"，改了不影响外面
func addOneByValue(n int) {
	n = n + 1 // 只改了复印件
}

// 传指针：函数拿到的是"真变量的地址"，能改到外面
// *int 表示"一个指向 int 的指针"
func addOneByPointer(n *int) {
	*n = *n + 1 // *n 表示"打开地址，操作里面的真值"
}

func main() {
	// ======== 1) & 取地址，* 取值 ========
	age := 18
	fmt.Println("age 的值:", age)   // 18
	fmt.Println("age 的地址:", &age) // 类似 0xc00001c030（每次运行可能不同）

	// p 是一个指针，存着 age 的地址
	p := &age
	fmt.Println("p 存的地址:", p)     // 和 &age 相同
	fmt.Println("p 指向的值 *p:", *p) // 18（打开盒子看里面）

	// ======== 2) 通过指针修改原变量 ========
	*p = 20                           // 打开地址，把里面的值改成 20
	fmt.Println("通过指针改后, age =", age) // 20（原变量真的变了！）

	fmt.Println("---")

	// ======== 3) 值传递 vs 指针传递 ========
	x := 100
	addOneByValue(x)
	fmt.Println("addOneByValue 后, x =", x) // 还是 100（没变）

	addOneByPointer(&x)
	fmt.Println("addOneByPointer 后, x =", x) // 101（变了）

	fmt.Println("---")

	// ======== 4) 这就是 fmt.Scan(&age) 里 & 的原因 ========
	// Scan 需要"真变量的地址"，才能把你输入的值写回到 age
	fmt.Println("所以 fmt.Scan(&x) 里的 & 就是把 x 的地址交给 Scan")

	// practice
	a := 1
	b := 2
	fmt.Println("before swap, a =", a, "b =", b) // a = 1, b = 2
	swap(&a, &b)
	fmt.Println("after swap, a =", a, "b =", b) // a = 2, b = 1
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
