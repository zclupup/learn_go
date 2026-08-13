package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// ======== Lesson 17：文件读写 + JSON 文件 ========
//
// Go 里常用 os 包做简单文件读写：
// - os.WriteFile(path, data, perm) → 写文件
// - os.ReadFile(path)              → 读文件
//
// 注意：
// - 文件内容通常用 []byte 表示，所以字符串写入前要用 []byte(text) 转换。
// - 0644 是文件权限：文件拥有者可读写，其他人只读。
// - 相对路径是相对于运行命令时所在的目录，不一定是 main.go 所在目录。
//
// Python 对比：
// - os.WriteFile 类似 open(path, "w").write(...)
// - os.ReadFile 类似 open(path).read()
// - json.MarshalIndent + os.WriteFile 类似 json.dump(..., indent=2)
// - os.ReadFile + json.Unmarshal 类似 json.load(...)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type Book struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

const (
	notePath     = "lesson17_file_json/lesson17_note.txt"
	studentsPath = "lesson17_file_json/students.json"
	bookPath     = "lesson17_file_json/book.json"
)

func main() {
	fmt.Println("=== 1. 写入普通文本文件 ===")

	text := "hello file\n这是 Go 写入的文本\n"
	err := os.WriteFile(notePath, []byte(text), 0644)
	if err != nil {
		fmt.Println("写入文本失败:", err)
		return
	}
	fmt.Println("文本写入成功")

	fmt.Println("\n=== 2. 读取普通文本文件 ===")

	textData, err := os.ReadFile(notePath)
	if err != nil {
		fmt.Println("读取文本失败:", err)
		return
	}
	fmt.Println(string(textData))

	fmt.Println("=== 3. 把结构体切片保存成 JSON 文件 ===")

	students := []Student{
		{ID: 1, Name: "张三", Score: 95},
		{ID: 2, Name: "李四", Score: 88},
		{ID: 3, Name: "王五", Score: 76},
	}

	jsonData, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		fmt.Println("转 JSON 失败:", err)
		return
	}

	err = os.WriteFile(studentsPath, jsonData, 0644)
	if err != nil {
		fmt.Println("写入 JSON 文件失败:", err)
		return
	}
	fmt.Println("students.json 写入成功")

	fmt.Println("\n=== 4. 从 JSON 文件读取并解析回结构体 ===")

	readJSONData, err := os.ReadFile(studentsPath)
	if err != nil {
		fmt.Println("读取 JSON 文件失败:", err)
		return
	}

	var loadedStudents []Student
	err = json.Unmarshal(readJSONData, &loadedStudents)
	if err != nil {
		fmt.Println("解析 JSON 文件失败:", err)
		return
	}

	fmt.Printf("解析结果: %+v\n", loadedStudents)

	fmt.Println("\n=== 5. 遍历从文件里读出来的数据 ===")
	for _, student := range loadedStudents {
		fmt.Printf("%s 的成绩是 %d\n", student.Name, student.Score)
	}

	books := []Book{
		{Title: "Go 入门", Price: 39.9},
		{Title: "Python 高级", Price: 49.5},
	}

	fmt.Println("\n=== 6. 把结构体切片保存成 JSON 文件（书籍） ===")
	jsonBooks, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		fmt.Println("转 JSON 失败:", err)
		return
	}

	err = os.WriteFile(bookPath, jsonBooks, 0644)
	if err != nil {
		fmt.Println("写入 JSON 文件失败:", err)
		return
	}
	fmt.Println("book.json 写入成功")

	fmt.Println("\n=== 7. 从 JSON 文件读取并解析回结构体（书籍） ===")
	readBookData, err := os.ReadFile(bookPath)
	if err != nil {
		fmt.Println("读取 JSON 文件失败:", err)
		return
	}

	var loadedBooks []Book
	err = json.Unmarshal(readBookData, &loadedBooks)
	if err != nil {
		fmt.Println("解析 JSON 文件失败:", err)
		return
	}

	fmt.Printf("解析结果: %+v\n", loadedBooks)

	fmt.Println("\n=== 8. 遍历从文件里读出来的书籍数据 ===")
	for _, book := range loadedBooks {
		fmt.Printf("%s 的价格是 %.2f\n", book.Title, book.Price)
	}
}
