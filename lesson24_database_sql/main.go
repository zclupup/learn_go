package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ======== Lesson 24：数据库入门：database/sql + MySQL ========
//
// 这一课先不用 GORM，先学 Go 标准库的 database/sql。
// 这样以后看 issue_api/internal/data 里的 GORM 代码时，会更容易理解：
// - 连接数据库
// - 执行 SQL
// - 查询多行数据
// - 把数据库列 Scan 到结构体字段
// - 用 context 控制超时
//
// 本课需要一个 MySQL DSN 才能真正连库。为了避免把密码写进代码，用环境变量：
// export LEARN_GO_MYSQL_DSN='user:password@tcp(127.0.0.1:3306)/learn_go?charset=utf8mb4&parseTime=True&loc=Local'
// go run ./lesson24_database_sql

type Student struct {
	ID        int
	Name      string
	Score     int
	CreatedAt time.Time
}

const createTableSQL = `
CREATE TABLE IF NOT EXISTS lesson24_students (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    score INT NOT NULL,
    created_at DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
`

func main() {
	dsn := os.Getenv("LEARN_GO_MYSQL_DSN")
	if strings.TrimSpace(dsn) == "" {
		fmt.Println("未设置 LEARN_GO_MYSQL_DSN，跳过真实数据库连接。")
		fmt.Println("示例：export LEARN_GO_MYSQL_DSN='user:password@tcp(127.0.0.1:3306)/learn_go?charset=utf8mb4&parseTime=True&loc=Local'")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("创建 DB 对象失败:", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	fmt.Println("数据库连接成功")

	if err := createTable(ctx, db); err != nil {
		log.Fatal("创建表失败:", err)
	}

	studentID, err := createStudent(ctx, db, "张三", 95)
	if err != nil {
		log.Fatal("创建学生失败:", err)
	}
	fmt.Println("创建学生成功，ID =", studentID)

	students, err := listStudents(ctx, db)
	if err != nil {
		log.Fatal("查询学生失败:", err)
	}

	for _, student := range students {
		fmt.Printf("ID=%d Name=%s Score=%d CreatedAt=%s\n",
			student.ID,
			student.Name,
			student.Score,
			student.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}
}

func createTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, createTableSQL)
	return err
}

func createStudent(ctx context.Context, db *sql.DB, name string, score int) (int64, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return 0, errors.New("name 不能为空")
	}

	result, err := db.ExecContext(
		ctx,
		"INSERT INTO lesson24_students (name, score, created_at) VALUES (?, ?, ?)",
		name,
		score,
		time.Now(),
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func listStudents(ctx context.Context, db *sql.DB) ([]Student, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, name, score, created_at FROM lesson24_students ORDER BY id DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := make([]Student, 0)
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Score, &student.CreatedAt); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
