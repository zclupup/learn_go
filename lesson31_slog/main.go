package main

import (
	"fmt"
	"log/slog"
	"os"
)

// ================================================================
// Lesson 31 — log/slog 结构化日志
// ================================================================
//
// 之前学过的日志方式：
//   - fmt.Println("用户", name, "支付", amount)    → 非结构化，字符串拼接
//   - log.Println("Error:", err)                    → 有时间戳，但没有结构化字段
//
// slog（Go 1.21+ 标准库）解决的核心问题：
//   日志不只是给人看的，更是给机器（日志系统）检索的。
//   结构化日志 = 每条日志是"键值对的集合"，可以按字段查询。
//
// 对照 issue-consumer 中的 slog 用法：
//   slog.Info("Message claimed",
//     slog.String("value", string(message.Value)),
//     slog.Time("timestamp", message.Timestamp),
//     slog.String("topic", message.Topic))
//
// ================================================================

func main() {
	// ============ 1. 对比：非结构化 vs 结构化 ============
	fmt.Println("========== 1. 非结构化 vs 结构化 ==========")

	// 非结构化写法（以前的习惯）
	fmt.Println("非结构化写法的输出：")
	fmt.Println("2026-08-22 10:30:00 INFO 用户 zhangcl 支付 100 元")
	fmt.Println(`→ 要搜索"支付100元"的记录？只能 grep 全文，没法按字段查`)
	fmt.Println()

	// 结构化写法（slog）
	fmt.Println("结构化写法的输出（JSON 格式）：")
	fmt.Println(`{"time":"2026-08-22T10:30:00","level":"INFO","msg":"支付成功","user":"zhangcl","amount":100}`)
	fmt.Println("→ 可以按字段查询：user=zhangcl AND amount>50")
	fmt.Println()

	// ============ 2. 基础用法：创建 Logger ============
	fmt.Println("========== 2. 基础用法 ==========")

	// 方式一：用默认 Logger（最简单）
	slog.Info("这是一条默认 Info 日志")
	slog.Error("这是一条默认 Error 日志")

	// 方式二：创建自定义 Logger（最常用，和 issue-consumer 一样）
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo, // 只输出 Info 及以上级别
		AddSource: false,          // 是否输出文件名和行号
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)

	logger.Info("用自定义 Logger 输出 Info 日志")
	logger.Error("用自定义 Logger 输出 Error 日志")
	// Debug 级别低于 Info，不会输出
	logger.Debug("这条 Debug 不会输出，因为 Level 是 Info")
	fmt.Println()

	// ============ 3. 关键字段：slog.String / slog.Int 等 ============
	fmt.Println("========== 3. 关键字段 ==========")

	// 每条日志后面可以附加多个键值对
	logger.Info("支付成功",
		slog.String("user", "zhangcl"),
		slog.Int("amount", 100),
		slog.Float64("discount", 0.95),
		slog.Bool("vip", true),
		slog.Any("items", []string{"item1", "item2"}),
	)
	fmt.Println()

	// ============ 4. 日志级别 ============
	fmt.Println("========== 4. 日志级别 ==========")

	// 四个级别（从低到高）：
	// Debug < Info < Warn < Error
	logger.Debug("Debug：开发调试用，生产环境通常不输出")
	logger.Info("Info：重要信息，如请求处理成功")
	logger.Warn("Warn：警告，如重试中")
	logger.Error("Error：错误，但仍然继续运行（不 panic）")
	fmt.Println()

	// ============ 5. JSONHandler vs TextHandler ============
	fmt.Println("========== 5. JSONHandler vs TextHandler ==========")

	// TextHandler：适合本地开发，人读
	textLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	textLogger.Info("TextHandler 输出", slog.String("key", "value"))

	// JSONHandler：适合生产环境，机器解析
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	jsonLogger.Info("JSONHandler 输出", slog.String("key", "value"))
	fmt.Println()

	// ============ 6. AddSource：显示文件名和行号 ============
	fmt.Println("========== 6. AddSource：显示文件名和行号 ==========")

	// 和 issue-consumer 完全一致
	// main.go 第 24-26 行：
	//   opts := &slog.HandlerOptions{
	//       Level:     slog.LevelInfo,
	//       AddSource: true,  // ← 输出文件名和行号
	//   }
	sourceLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}))
	sourceLogger.Info("这条日志带文件名和行号")
	fmt.Println()

	// ============ 7. slog.SetDefault：设置全局 Logger ============
	// 和 issue-consumer main.go 第 33-34 行完全一致
	//   logger := slog.New(handler)
	//   slog.SetDefault(logger)
	//
	// 设置后，项目里任何地方直接写 slog.Info(...) 都使用这个 logger，
	// 不需要把 logger 作为参数到处传。

	// ============ 8. 对照 issue-consumer 的真实用法 ============
	fmt.Println("========== 8. 对照 issue-consumer ==========")
	fmt.Println("mark_tool.go 第 47 行：")
	fmt.Println(`  slog.Info("Message claimed",`)
	fmt.Println(`    slog.String("value", string(message.Value)),`)
	fmt.Println(`    slog.Time("timestamp", message.Timestamp),`)
	fmt.Println(`    slog.String("topic", message.Topic))`)
	fmt.Println()
	fmt.Println("mark_tool.go 第 50 行：")
	fmt.Println(`  slog.Error("get json tag filed error", slog.String("error", err.Error()))`)
	fmt.Println()
	fmt.Println("issue_tracking_result.go 第 96 行：")
	fmt.Println(`  slog.Error("BatchImportIssue failed", slog.Int64("err_no", resp.ErrNo), slog.String("err_msg", resp.ErrMsg))`)
	fmt.Println()

	// ============ 9. 总结 ============
	fmt.Println("========== 总结 ==========")
	fmt.Println("| 概念 | 说明 |")
	fmt.Println("|------|------|")
	fmt.Println("| slog.Info(msg, kvs...) | 输出 Info 级别日志 |")
	fmt.Println("| slog.Error(msg, kvs...) | 输出 Error 级别日志 |")
	fmt.Println("| slog.String(k, v) | 字符串键值对 |")
	fmt.Println("| slog.Int(k, v) | 整数键值对 |")
	fmt.Println("| slog.Int64(k, v) | int64 键值对 |")
	fmt.Println("| slog.Any(k, v) | 任意类型键值对 |")
	fmt.Println("| TextHandler | 人读的文本格式 |")
	fmt.Println("| JSONHandler | 机器解析的 JSON 格式 |")
	fmt.Println("| HandlerOptions.Level | 控制输出级别 |")
	fmt.Println("| HandlerOptions.AddSource | 是否输出文件名行号 |")
	fmt.Println("| slog.SetDefault(logger) | 设为全局默认 Logger |")
	fmt.Println()
	fmt.Println("和 Python 的对比：")
	fmt.Println("  Python logging: logger.info('xxx', extra={'user': 'zhangcl'})")
	fmt.Println("  Go slog:        slog.Info('xxx', slog.String('user', 'zhangcl'))")
}