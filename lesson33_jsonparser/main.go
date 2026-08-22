package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/buger/jsonparser"
)

// ================================================================
// Lesson 33 — jsonparser JSON 快速解析
// ================================================================
//
// 之前学过的 JSON 解析方式（Lesson 16）：
//   json.Unmarshal(data, &struct)  → 必须定义结构体
//   json.Unmarshal(data, &map)     → 不用结构体，但取值要类型断言
//
// jsonparser 是另一种方式：
//   - 不用定义结构体
//   - 按路径直接取字段（类似 Python 的 jsonpath）
//   - 性能比标准库快
//   - 适合"从大 JSON 中取几个字段"的场景
//
// 对照 issue-consumer 中的 jsonparser 用法：
//   mark_tool.go:48  → jsonparser.GetString(message.Value, "user_id")
//   mark_tool.go:66  → jsonparser.ArrayEach(message.Value, func(...), "主菜单")
//
// ================================================================

func main() {
	fmt.Println("========== Lesson 33: jsonparser JSON 快速解析 ==========")
	fmt.Println()

	// 模拟一条来自 Kafka 的 JSON 消息（和 issue-consumer 类似）
	jsonData := []byte(`{
		"user_id": "product_validation",
		"标记时间": "2026-03-09 11:13:10",
		"主菜单": ["冲出车道", "偏离车道"],
		"一级菜单": ["横向控制"],
		"二级菜单": ["分流场景", "失败", "安全接管"],
		"备注信息": "测试车辆偏离车道",
		"用户": "zhangcl",
		"项目": "H47A-J6p",
		"车辆编号": "233",
		"软件版本": "0307",
		"测试主题": "冒烟",
		"VIN": "LDP95C96XTY007720",
		"nested": {
			"level1": {
				"level2": "deep_value"
			}
		}
	}`)

	// ============ 1. 标准库方式：先定义结构体 ============
	fmt.Println("--- 1. 标准库方式：需要定义结构体 ---")

	type Message struct {
		UserID    string   `json:"user_id"`
		EventTime string   `json:"标记时间"`
		MainMenu  []string `json:"主菜单"`
		Desc      string   `json:"备注信息"`
		VIN       string   `json:"VIN"`
	}
	var msg Message
	json.Unmarshal(jsonData, &msg)
	fmt.Printf("UserID: %s, VIN: %s, 备注: %s\n", msg.UserID, msg.VIN, msg.Desc)
	fmt.Println("问题：每来一种新消息，就得定义一个新结构体")
	fmt.Println()

	// ============ 2. jsonparser：不用结构体，直接取字段 ============
	fmt.Println("--- 2. jsonparser：按路径直接取值 ---")

	userID, _ := jsonparser.GetString(jsonData, "user_id")
	eventTime, _ := jsonparser.GetString(jsonData, "标记时间")
	desc, _ := jsonparser.GetString(jsonData, "备注信息")
	vin, _ := jsonparser.GetString(jsonData, "VIN")

	fmt.Printf("user_id:   %s\n", userID)
	fmt.Printf("标记时间:   %s\n", eventTime)
	fmt.Printf("备注信息:   %s\n", desc)
	fmt.Printf("VIN:       %s\n", vin)
	fmt.Println()

	// ============ 3. 取嵌套字段 ============
	fmt.Println("--- 3. 取嵌套字段 ---")

	// 嵌套路径用多个参数：GetString(data, "nested", "level1", "level2")
	deepValue, _ := jsonparser.GetString(jsonData, "nested", "level1", "level2")
	fmt.Printf("nested.level1.level2 = %s\n", deepValue)
	fmt.Println("标准库要定义嵌套结构体，jsonparser 直接写路径即可")
	fmt.Println()

	// ============ 4. 遍历数组 ============
	fmt.Println("--- 4. 遍历数组 ---")
	fmt.Println("主菜单:")

	jsonparser.ArrayEach(jsonData, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		// value 是数组元素的原始字节，直接转 string
		fmt.Printf("  - %s\n", string(value))
	}, "主菜单")

	fmt.Println("二级菜单:")
	jsonparser.ArrayEach(jsonData, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Printf("  - %s\n", string(value))
	}, "二级菜单")
	fmt.Println()

	// ============ 5. 取其他类型 ============
	fmt.Println("--- 5. 其他类型 ---")

	// 取数字
	numberData := []byte(`{"count": 42, "price": 99.9, "active": true}`)

	count, _ := jsonparser.GetInt(numberData, "count")
	price, _ := jsonparser.GetFloat(numberData, "price")
	active, _ := jsonparser.GetBoolean(numberData, "active")

	fmt.Printf("count=%d (int), price=%.1f (float), active=%t (bool)\n", count, price, active)
	fmt.Println()

	// ============ 6. 判断字段是否存在 ============
	fmt.Println("--- 6. 判断字段是否存在 ---")

	value, dataType, _, err := jsonparser.Get(jsonData, "VIN")
	if err != nil {
		fmt.Println("VIN 字段不存在")
	} else {
		fmt.Printf("VIN 存在，类型=%v，值=%s\n", dataType, string(value))
	}

	// 不存在的字段
	_, _, _, err = jsonparser.Get(jsonData, "不存在的字段")
	if err != nil {
		fmt.Printf("不存在的字段：%v\n", err)
	}
	fmt.Println()

	// ============ 7. 对照 issue-consumer 真实代码 ============
	fmt.Println("--- 7. 对照 issue-consumer 真实代码 ---")
	fmt.Println("mark_tool.go 中 jsonparser 的用法：")
	fmt.Println()
	fmt.Println("  // 取单个字段")
	fmt.Println("  userID, err := jsonparser.GetString(message.Value, \"user_id\")")
	fmt.Println("  eventTime, err := jsonparser.GetString(message.Value, \"标记时间\")")
	fmt.Println("  vin, err := jsonparser.GetString(message.Value, \"VIN\")")
	fmt.Println()
	fmt.Println("  // 遍历数组")
	fmt.Println("  jsonparser.ArrayEach(message.Value, func(action []byte, ...) {")
	fmt.Println("      mainTag = append(mainTag, string(action))")
	fmt.Println("  }, \"主菜单\")")
	fmt.Println()

	// ============ 8. 对比总结 ============
	fmt.Println("========== 对比总结 ==========")
	fmt.Println()
	fmt.Println("| 场景 | 用 json.Unmarshal | 用 jsonparser |")
	fmt.Println("|------|------------------|--------------|")
	fmt.Println("| 结构固定，字段多 | ✅ 定义结构体一次搞定 | ❌ 每个字段都要写 GetString |")
	fmt.Println("| 只需要 1-2 个字段 | ❌ 也要定义结构体 | ✅ 直接 GetString |")
	fmt.Println("| 字段名是动态的（如中文） | ❌ 结构体不能有中文字段名 | ✅ 直接传字符串路径 |")
	fmt.Println("| 深层嵌套取值 | ❌ 要定义多层嵌套结构体 | ✅ GetString(data, \"a\",\"b\",\"c\") |")
	fmt.Println("| 性能 | 一般 | 更快 |")
	fmt.Println()
	fmt.Println("issue-consumer 为什么用 jsonparser？")
	fmt.Println("  因为 Kafka 消息中的字段名是中文的（\"标记时间\"、\"主菜单\"、\"备注信息\"），")
	fmt.Println("  用标准库虽然能解析（json tag 支持中文），但 jsonparser 的路径写法更直观，")
	fmt.Println("  而且很多字段只需要取一个值，不需要定义整个结构体。")
}

// 确保导入了 strings（虽然没直接使用，但为了展示）
var _ = strings.ToUpper