package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// ================================================================
// Lesson 35 — TOML 配置读取
// ================================================================
//
// 之前学过 YAML 配置读取（Lesson 27），现在学 TOML。
//
// TOML 和 YAML 都是配置格式，区别：
//   YAML：用缩进表示层级，人读友好，但容易出错（缩进敏感）
//   TOML：用 [section] 表示层级，更接近 INI 格式，简单清晰
//
// 对照 issue-consumer：
//   配置格式：TOML（不是 YAML）
//   解析库：github.com/BurntSushi/toml
//   配置文件：configs/config.toml
//   解析代码：configs/config.go → toml.DecodeFile
//   结构体 tag：`toml:"brokers"`
//
// ================================================================

func main() {
	fmt.Println("========== Lesson 35: TOML 配置读取 ==========")
	fmt.Println()

	// ============ 1. TOML vs YAML 对比 ============
	fmt.Println("--- 1. TOML vs YAML 格式对比 ---")
	fmt.Println()
	fmt.Println("YAML 写法（Lesson 27）：")
	fmt.Println("  server:")
	fmt.Println("    host: 0.0.0.0")
	fmt.Println("    port: 8080")
	fmt.Println()
	fmt.Println("TOML 写法（本课）：")
	fmt.Println("  [server]")
	fmt.Println("  host = \"0.0.0.0\"")
	fmt.Println("  port = 8080")
	fmt.Println()
	fmt.Println("YAML 用缩进表示层级，TOML 用 [section] 表示层级")
	fmt.Println("YAML 容易写错缩进，TOML 更接近 INI 格式，更简单")
	fmt.Println()

	// ============ 2. 读取 TOML 配置 ============
	fmt.Println("--- 2. 读取 TOML 配置 ---")

	var config Config
	_, err := toml.DecodeFile("lesson35_toml/config.toml", &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取配置失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Title: %s\n", config.Title)
	fmt.Printf("Server: %s:%d\n", config.Server.Host, config.Server.Port)
	fmt.Printf("Database: %s@%s:%d/%s\n", config.Database.User, config.Database.Host, config.Database.Port, config.Database.Name)
	fmt.Printf("Log: level=%s, format=%s\n", config.Log.Level, config.Log.Format)
	fmt.Println()

	// ============ 3. 读取嵌套配置 ============
	fmt.Println("--- 3. 读取嵌套配置 ---")

	fmt.Printf("Kafka MarkerTool: brokers=%s, topics=%s, group=%s\n",
		config.Kafka.MarkerTool.Brokers,
		config.Kafka.MarkerTool.Topics,
		config.Kafka.MarkerTool.Group)
	fmt.Printf("Kafka ProdMp3: brokers=%s, topics=%s, group=%s\n",
		config.Kafka.ProdMp3.Brokers,
		config.Kafka.ProdMp3.Topics,
		config.Kafka.ProdMp3.Group)
	fmt.Println()

	// ============ 4. 读取数组 ============
	fmt.Println("--- 4. 读取数组（[[users]]）---")

	for _, u := range config.Users {
		fmt.Printf("用户: %s, 角色: %s\n", u.Name, u.Role)
	}
	fmt.Println()

	// ============ 5. 对照 issue-consumer 真实代码 ============
	fmt.Println("--- 5. 对照 issue-consumer 真实代码 ---")
	fmt.Println("configs/config.go 中的结构体：")
	fmt.Println()
	fmt.Println("  type KafkaConfig struct {")
	fmt.Println("      Brokers           string `toml:\"brokers\"`")
	fmt.Println("      Topics            string `toml:\"topics\"`")
	fmt.Println("      Group             string `toml:\"group\"`")
	fmt.Println("      Version           string `toml:\"version\"`")
	fmt.Println("      Oldest            bool   `toml:\"oldest\"`")
	fmt.Println("      Assignor          string `toml:\"assignor\"`")
	fmt.Println("      ...")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  type AppConfigs struct {")
	fmt.Println("      MarkerToolKafka   KafkaConfig `toml:\"markerToolKafka\"`")
	fmt.Println("      ProdMp3Kafka      KafkaConfig `toml:\"prodMp3Kafka\"`")
	fmt.Println("      ...")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  func LoadConfig(configFile string) (*AppConfigs, error) {")
	fmt.Println("      config := &AppConfigs{}")
	fmt.Println("      _, err := toml.DecodeFile(configFile, config)")
	fmt.Println("      return config, err")
	fmt.Println("  }")
	fmt.Println()

	// ============ 6. 总结：YAML vs TOML ============
	fmt.Println("========== 总结：YAML vs TOML ==========")
	fmt.Println()
	fmt.Println("| 对比项 | YAML | TOML |")
	fmt.Println("|--------|------|------|")
	fmt.Println("| 层级表示 | 缩进 | [section] |")
	fmt.Println("| 键值对 | key: value | key = value |")
	fmt.Println("| 数组 | [1, 2, 3] 或 - 1 | [1, 2, 3] |")
	fmt.Println("| 对象数组 | - key: v | [[array]] |")
	fmt.Println("| 解析库 | gopkg.in/yaml.v3 | github.com/BurntSushi/toml |")
	fmt.Println("| tag | `yaml:\"key\"` | `toml:\"key\"` |")
	fmt.Println("| 解析函数 | yaml.Unmarshal | toml.DecodeFile |")
	fmt.Println()
	fmt.Println("原理完全一样：文本格式 → 解析到结构体，靠 tag 映射字段名。")
	fmt.Println("区别只在于格式语法和 tag 名不同。")
}

// ============================================================
// 结构体定义（对应 TOML 配置）
// ============================================================

// Config 对应 config.toml 的顶层结构。
type Config struct {
	Title    string
	Server   ServerConfig
	Database DatabaseConfig
	Log      LogConfig
	Kafka    KafkaConfig
	Users    []UserConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type LogConfig struct {
	Level  string
	Format string
}

type KafkaConfig struct {
	MarkerTool KafkaTopicConfig `toml:"marker_tool"`
	ProdMp3    KafkaTopicConfig `toml:"prod_mp3"`
}

type KafkaTopicConfig struct {
	Brokers  string
	Topics   string
	Group    string
	Version  string
	Oldest   bool
	Assignor string
}

type UserConfig struct {
	Name string
	Role string
}