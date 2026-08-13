package main

import (
	"encoding/json"
	"fmt"
)

// ======== Lesson 16：JSON 编码/解码 + struct tag ========
//
// Go 里处理 JSON 用 encoding/json 包：
// - json.Marshal(v)      → Go 结构体 → JSON 字节
// - json.Unmarshal(data, &v) → JSON 字节 → Go 结构体
//
// struct tag 用来控制 JSON 字段名、是否忽略等：
// - `json:"name"`       → JSON 字段名是 "name"
// - `json:"name,omitempty"` → Marshal 时如果字段是零值就不输出
// - `json:"-"`          → Marshal/Unmarshal 都忽略这个字段
// tag 只能写在 struct 字段上；不同包会读取不同 tag，比如 json、db、form。
//
// Python 对比：
// - json.dumps(obj) → json.Marshal
// - json.loads(s)   → json.Unmarshal
// - dataclass/pydantic → struct + tag

// ========== 1. 基础结构体 ==========

// encoding/json 只能访问“导出字段”：字段名首字母必须大写。
// 如果写成 name、age 这种小写字段，即使有 json tag，也不会被 Marshal/Unmarshal 处理。

// User 用户信息（没有 tag）
type User struct {
	Name  string
	Age   int
	Email string
}

// UserWithTag 用户信息（带 tag）
type UserWithTag struct {
	Name  string `json:"name"`            // Marshal 时输出 key 为 "name"；Unmarshal 时也用 "name" 匹配到 Name 字段
	Age   int    `json:"age"`             // json tag 是 struct 字段的元数据，encoding/json 会读取它
	Email string `json:"email,omitempty"` // omitempty 只影响 Marshal/MarshalIndent：零值时不输出；不影响 Unmarshal
}

// UserPrivate 带私有字段
type UserPrivate struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"-"` // 忽略：Marshal 不输出，Unmarshal 也不会写入这个字段
}

// practice
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock,omitempty"`
}

func main() {
	fmt.Println("=== 1. Marshal：结构体 → JSON ===")

	// 没有 tag，字段名保持原样（首字母大写）
	user1 := User{Name: "张三", Age: 25, Email: "zhangsan@example.com"}
	data1, _ := json.Marshal(user1)
	fmt.Println("无 tag:", data1) // 输出字节切片
	fmt.Println("无 tag:", string(data1))

	// 有 tag，字段名按 tag 来
	user2 := UserWithTag{Name: "李四", Age: 30, Email: "lisi@example.com"}
	data2, _ := json.Marshal(user2)
	fmt.Println("有 tag:", string(data2))

	// omitempty：空值不输出
	user3 := UserWithTag{Name: "王五", Age: 28, Email: ""} // Email 为空
	data3, _ := json.Marshal(user3)
	fmt.Println("omitempty:", string(data3)) // 不会有 email 字段

	// 忽略字段：Password 不会出现
	user4 := UserPrivate{Name: "赵六", Age: 35, Password: "secret123"}
	data4, _ := json.Marshal(user4)
	fmt.Println("忽略字段:", string(data4)) // 没有 password

	fmt.Println("\n=== 2. MarshalIndent：格式化输出 ===")
	// 第二个参数 prefix 表示每一行前面额外加什么前缀；通常传空字符串。
	// 第三个参数 indent 表示每一级缩进用什么字符；这里用两个空格。
	prettyData, _ := json.MarshalIndent(user2, "", "  ")
	fmt.Println(string(prettyData))

	fmt.Println("\n=== 3. Unmarshal：JSON → 结构体 ===")

	jsonStr := `{"name": "孙七", "age": 40, "email": "sunqi@example.com"}`
	var user5 UserWithTag
	// Unmarshal 第二个参数要传指针：把字节数据解析后，写入 &user5 指向的那块内存。
	err := json.Unmarshal([]byte(jsonStr), &user5) // 注意：传指针 &user5
	if err != nil {
		fmt.Println("解析失败:", err)
	} else {
		fmt.Printf("解析结果: %+v\n", user5)
	}

	// JSON 字段比结构体多：多余的字段被忽略
	jsonExtra := `{"name": "周八", "age": 45, "email": "zhouba@example.com", "city": "北京"}`
	var user6 UserWithTag
	json.Unmarshal([]byte(jsonExtra), &user6)
	fmt.Printf("多余字段被忽略: %+v\n", user6)

	// JSON 字段比结构体少：缺少的字段是零值
	jsonMissing := `{"name": "钱九"}`
	var user7 UserWithTag
	json.Unmarshal([]byte(jsonMissing), &user7)
	fmt.Printf("缺少字段用零值: %+v\n", user7) // Age=0, Email=""

	fmt.Println("\n=== 4. 嵌套结构体 ===")

	type Address struct {
		City   string `json:"city"`
		Street string `json:"street"`
	}
	type Person struct {
		Name    string  `json:"name"`
		Address Address `json:"address"` // 嵌套结构体
	}
	person := Person{
		Name: "吴十",
		Address: Address{
			City:   "上海",
			Street: "南京路",
		},
	}
	personData, _ := json.MarshalIndent(person, "", "  ")
	fmt.Println(string(personData))

	fmt.Println("\n=== 5. 切片/数组 ===")

	users := []UserWithTag{
		{Name: "用户A", Age: 20, Email: "a@test.com"},
		{Name: "用户B", Age: 25, Email: ""},
	}
	usersData, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println(string(usersData))

	// 解析 JSON 数组
	jsonArray := `[{"name":"X","age":1},{"name":"Y","age":2}]`
	var userList []UserWithTag
	json.Unmarshal([]byte(jsonArray), &userList)
	fmt.Printf("解析数组: %+v\n", userList)

	fmt.Println("\n=== 6. map[string]interface{} 处理动态 JSON ===")

	// 当你不知道 JSON 结构时，用 map
	dynamicJSON := `{"name": "动态", "score": 99.5, "active": true}`
	// interface{} 表示任意类型；Go 1.18+ 里 any 是 interface{} 的别名，也可以写 map[string]any。
	var result map[string]interface{}
	json.Unmarshal([]byte(dynamicJSON), &result)
	fmt.Printf("动态解析: %v\n", result)
	fmt.Printf("取值: name=%v, score=%v\n", result["name"], result["score"])

	// ⚠️ 注意：数字默认解析为 float64
	fmt.Printf("score 的类型: %T\n", result["score"]) // float64

	product := Product{ID: 1, Name: "zcl", Price: 99.8, Stock: 0}
	productData, _ := json.Marshal(product)
	fmt.Println(string(productData))

	products := []Product{
		{ID: 1, Name: "Product A", Price: 10.5, Stock: 100},
		{ID: 2, Name: "Product B", Price: 20.0, Stock: 50},
	}
	productsData, _ := json.MarshalIndent(products, "", "  ")
	fmt.Println(string(productsData))

	productJson := `{"id":9, "name":"zz", "price": 99.8, "stock": 0}`
	var productFromJson Product
	errPrac := json.Unmarshal([]byte(productJson), &productFromJson)
	if errPrac != nil {
		fmt.Println("解析失败:", errPrac)
	} else {
		fmt.Printf("解析单个产品: %+v\n", productFromJson)
		fmt.Printf("解析单个产品: %T\n", productFromJson)
	}
}
