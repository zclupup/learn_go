package course

// Course 是公开结构体：首字母大写，所以 main package 可以使用。
type Course struct {
	ID      int
	Name    string
	Teacher string
}

// NewCourse 是公开函数：Go 里常用 NewXxx 命名构造函数。
func NewCourse(id int, name string, teacher string) Course {
	return Course{
		ID:      id,
		Name:    name,
		Teacher: teacher,
	}
}

// Summary 是 Course 的方法。
// (c Course) 叫方法接收者，表示这个方法属于 Course 类型；c 类似 Python 方法里的 self。
func (c Course) Summary() string {
	return formatCourseInfo(c)
}

// formatCourseInfo 是未公开函数：首字母小写，只能在 course package 内部使用。
func formatCourseInfo(c Course) string {
	return "课程信息：" + c.Name + "，教师：" + c.Teacher
}
