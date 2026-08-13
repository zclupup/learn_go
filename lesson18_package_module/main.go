package main

import (
	"fmt"

	// 项目内 import 路径 = go.mod 的 module 名 + 包目录相对项目根目录的路径。
	// go.mod 里是 module learn_go，course 包目录是 lesson18_package_module/course，
	// 所以完整导入路径是 learn_go/lesson18_package_module/course。
	"learn_go/lesson18_package_module/course"
	"learn_go/lesson18_package_module/student"
)

// ======== Lesson 18：包 package、目录结构与模块复习 ========
//
// Go 的组织方式可以先记成三层：
// - module：整个项目，由 go.mod 里的 module learn_go 定义。
// - package：代码包，通常一个目录就是一个 package。
// - file：具体 .go 文件，同一个目录下的 .go 文件必须写同一个 package 名。
//   这里的“同一个目录”只包含直接放在该目录下的 .go 文件，不包含子目录。
//   子目录会成为新的目录，也就可以是新的 package。
//
// package main 是特殊包：表示这个目录可以编译成可执行程序。
// func main() 是程序入口：go run 运行时从这里开始执行。
//
// 普通 package 不能直接运行，通常用来放可复用代码，然后被 main 包 import。
// 本课的 student 目录就是普通 package。

func main() {
	fmt.Println("=== 1. package main：程序入口 ===")
	fmt.Println("当前 main 函数在 package main 里，所以这个目录可以 go run")

	fmt.Println("\n=== 2. 调用另一个 package 里的公开函数 ===")
	student.PrintPackageInfo()

	fmt.Println("\n=== 3. 使用另一个 package 里的公开结构体 ===")
	stu := student.Student{
		ID:    1,
		Name:  "张三",
		Score: 95,
	}

	fmt.Println(stu.Summary())

	fmt.Println("\n=== 4. 调用公开函数创建结构体 ===")
	stu2 := student.NewStudent(2, "李四", 88)
	fmt.Println(stu2.Summary())

	fmt.Println("\n=== 5. 尝试调用未公开函数 ===")
	fmt.Println(student.HelperMessage())
	// fmt.Println(student.helper()) // ❌ 不能直接调用未公开函数；只能通过 student 包内部的公开函数间接使用。

	fmt.Println("\n=== 6. 练习：调用 course package ===")
	c := course.NewCourse(1, "Go 语言入门", "张老师")
	fmt.Println(c.Summary())
}
