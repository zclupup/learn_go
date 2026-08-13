package main

import "testing"

// go test 会自动查找当前 package 下所有 *_test.go 文件中形如 TestXxx(t *testing.T) 的函数。
// main_test.go 和 main.go 都是 package main，所以测试里可以直接调用 Add、Grade、Divide、Max。
// 同一个 package 的顶层函数、变量、常量、类型名不能重复；不同函数内部的局部变量可以同名。

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Fatalf("Add(2, 3) = %d, want %d", got, want)
	}
}

func TestIsPassingScore(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  bool
	}{
		{name: "60 分及格", score: 60, want: true},
		{name: "59 分不及格", score: 59, want: false},
		{name: "90 分及格", score: 90, want: true},
	}

	for _, tt := range tests {
		// t.Run 会创建并运行一个子测试。
		// 第二个参数 func(t *testing.T) 是匿名函数：先作为参数传给 t.Run，函数体再由 t.Run 调用执行。
		t.Run(tt.name, func(t *testing.T) {
			got := IsPassingScore(tt.score)
			if got != tt.want {
				t.Fatalf("IsPassingScore(%d) = %t, want %t", tt.score, got, tt.want)
			}
		})
	}
}

func TestGrade(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{name: "优秀", score: 95, want: "优秀"},
		{name: "良好", score: 85, want: "良好"},
		{name: "及格", score: 70, want: "及格"},
		{name: "不及格", score: 50, want: "不及格"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Grade(tt.score)
			if got != tt.want {
				t.Fatalf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("Divide(10, 2) unexpected error: %v", err)
	}

	want := 5
	if got != want {
		t.Fatalf("Divide(10, 2) = %d, want %d", got, want)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Fatal("Divide(10, 0) expected error, got nil")
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "a is greater", a: 5, b: 3, want: 5},
		// 这里故意把 want 写成 3，用来观察 go test 失败输出、文件行号和子测试结果。
		// 如果想让全部测试通过，把 want 改成 4。
		{name: "b is greater", a: 2, b: 4, want: 3},
		{name: "a equals b", a: 7, b: 7, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
