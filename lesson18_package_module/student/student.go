package student

import "fmt"

// Student 首字母大写，所以其他 package 可以使用这个类型。
type Student struct {
	ID    int
	Name  string
	Score int
}

// NewStudent 首字母大写，所以 main package 可以调用它。
func NewStudent(id int, name string, score int) Student {
	return Student{
		ID:    id,
		Name:  name,
		Score: score,
	}
}

// Summary 是公开方法，因为方法名首字母大写。
func (s Student) Summary() string {
	return fmt.Sprintf("学生：%s，成绩：%d", s.Name, s.Score)
}

// PrintPackageInfo 是公开函数，可以被 main package 调用。
func PrintPackageInfo() {
	fmt.Println("这里是 lesson18_package_module/student 包")
	fmt.Println("普通 package 用来放可复用代码，不能自己作为程序入口")
}

// HelperMessage 是公开函数，main package 可以调用。
// 它在 student package 内部调用了未公开的 helper，演示“外部不能直接调私有函数，但可以通过公开函数间接使用”。
func HelperMessage() string {
	return helper()
}

// helper 是未公开函数，只能在 student package 内部使用。
func helper() string {
	return "这个函数首字母小写，其他 package 不能直接调用"
}
