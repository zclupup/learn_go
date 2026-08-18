package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// ======== Lesson 27：配置文件读取：yaml + struct ========
//
// 这一课学习如何把 yaml 配置文件解析成 Go 结构体。
// 结构对齐 issue_api：
//   configs/config.yaml  -> yaml 文件
//   conf.Bootstrap       -> 顶层的 Config 结构体
//   server/data/project  -> 嵌套的子结构体
//
// issue_api 用的是 Kratos 的 config 库 + protobuf 生成结构体，
// 但底层思路一样：yaml 文本 -> 结构体字段。
// 本课用标准库方式（gopkg.in/yaml.v3）先掌握这个核心概念。

// Config 是顶层结构体，对应 issue_api 里的 Bootstrap。
type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Data    DataConfig    `yaml:"data"`
	Project ProjectConfig `yaml:"project"`
}

type ServerConfig struct {
	HTTP HTTPConfig `yaml:"http"`
	Env  string     `yaml:"env"`
}

type HTTPConfig struct {
	Addr    string `yaml:"addr"`
	Timeout string `yaml:"timeout"`
}

type DataConfig struct {
	Mysql MysqlConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
}

type MysqlConfig struct {
	DSN string `yaml:"dsn"`
}

type RedisConfig struct {
	Address string `yaml:"address"`
	DB      int    `yaml:"db"`
}

type ProjectConfig struct {
	ProjectName string `yaml:"project_name"`
	Mode        string `yaml:"mode"`
}

func main() {
	// 用 -conf 指定配置文件路径，和 issue_api 的 -conf 用法一致。
	confPath := flag.String("conf", "config.yaml", "配置文件路径")
	flag.Parse()

	data, err := os.ReadFile(*confPath)
	if err != nil {
		log.Fatal("读取配置文件失败:", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatal("解析配置失败:", err)
	}

	fmt.Println("=== 配置内容 ===")
	fmt.Printf("Server.HTTP.Addr    = %s\n", config.Server.HTTP.Addr)
	fmt.Printf("Server.HTTP.Timeout = %s\n", config.Server.HTTP.Timeout)
	fmt.Printf("Server.Env          = %s\n", config.Server.Env)
	fmt.Printf("Data.Mysql.DSN      = %s\n", config.Data.Mysql.DSN)
	fmt.Printf("Data.Redis.Address  = %s\n", config.Data.Redis.Address)
	fmt.Printf("Data.Redis.DB       = %d\n", config.Data.Redis.DB)
	fmt.Printf("Project.ProjectName = %s\n", config.Project.ProjectName)
	fmt.Printf("Project.Mode        = %s\n", config.Project.Mode)
}
