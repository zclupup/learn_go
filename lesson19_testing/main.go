package main

import "errors"

// ======== Lesson 19：测试入门 ========
//
// Go 标准库自带 testing 包，不需要安装第三方测试框架。
// 测试文件命名必须以 _test.go 结尾，例如 main_test.go。
// 测试函数命名必须以 Test 开头，形如 func TestXxx(t *testing.T)。
//
// 常用命令：
// - go test ./lesson19_testing       → 运行这一课的测试
// - go test -v ./lesson19_testing    → 显示每个测试函数的详细结果
// - go test ./...                    → 运行整个 module 下所有测试
// - go test -run TestMax ./lesson19_testing → 只运行名字匹配 TestMax 的测试
//
// Python 对比：
// - Go 的 testing 包类似 Python 的 unittest/pytest
// - t.Fatalf 类似 assert 失败后立刻停止当前测试
// - 表格驱动测试类似 pytest.mark.parametrize

func Add(a int, b int) int {
	return a + b
}

func IsPassingScore(score int) bool {
	return score >= 60
}

func Grade(score int) string {
	if score >= 90 {
		return "优秀"
	}
	if score >= 80 {
		return "良好"
	}
	if score >= 60 {
		return "及格"
	}
	return "不及格"
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 测试课的重点在 main_test.go。
	// main 函数这里留空，避免运行 go run 时影响测试示例。
}
