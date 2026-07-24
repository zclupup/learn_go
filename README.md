# learn_go — 从零开始学 Go（Python 背景）

这是一个 Go 语言入门学习项目，学习者有 Python 基础，从最简单的例子逐课学习 Go。

---

## 🤖 给 AI 助手的说明（换电脑后请先读这里）

我是一名 **Go 小白**，只有 **Python 基础**，正在从简单例子逐步学习 Go。请遵守以下规则：

1. **始终用小白视角**：变量、类型、语法等都要解释清楚，尽量对比 Python 来讲。
2. **一课一个主题**：每课新建一个 `lessonNN_主题.go` 文件，保持简单可运行。
3. **每课流程**：先讲概念 → 写示例代码 → 运行验证 → 给一个小练习。
4. **回答我的疑问**：我经常会问"为什么"，请耐心解释底层原理。
5. **运行命令**：本机 Go 路径为 `D:\software\Go\bin\go.exe`（终端里 `go` 有时不在 PATH）。
   运行单课：`D:\software\Go\bin\go.exe run .\lessonNN_xxx.go`
6. **每课自动更新 README（重要）**：**每当完成一节新课**（示例写完、运行通过、我做完练习后），
   你必须**主动更新本 README**，无需我额外提醒，具体包括：
   - 在「学习进度」表格中把该课状态改为 ✅ 完成，并新增下一课占位行（⬜ 下一课）；
   - 更新「当前进度」那句话；
   - 在「知识点总结」中追加这一课的重点（对比 Python，简明列点）；
   - 然后执行 `git add . && git commit -m "完成 LessonXX" && git push` 推送到 GitHub。
7. **触发口令**：我可能会用这些话触发上述更新，含义相同——
   "更新进度" / "记录进度" / "总结一下这节课" / "写入 README"。
   看到类似口令，或每完成一课，都按第 6 条执行。

---

## 📊 学习进度

| 课程 | 主题 | 文件 | 状态 |
|------|------|------|------|
| Lesson 01 | 变量声明与基本类型 | `lesson01_variables.go` | ✅ 完成 |
| Lesson 02 | 输入 + if 判断 | `lesson02_input_if.go` | ✅ 完成 |
| Lesson 03 | 格式化输出 printf | `lesson03_printf.go` | ✅ 完成 |
| Lesson 04 | for 循环 + 猜数字游戏 | `lesson04_for_guess.go` | ✅ 完成 |
| Lesson 05 | 数组与切片 slice | `lesson05_slice.go` | ✅ 完成 |
| Lesson 06 | 函数 func | `lesson06_function.go` | ✅ 完成 |
| Lesson 07 | 结构体 struct | `lesson07_struct.go` | ✅ 完成 |
| Lesson 08 | map（字典 dict） | `lesson08_map.go` | ✅ 完成 |
| Lesson 09 | 错误处理 error | `lesson09_error.go` | ✅ 完成 |
| Lesson 10 | *待定* | *待开始* | ⬜ 下一课 |

**当前进度：已完成 Lesson 01–09，下一课是 Lesson 10。**

---

## � 换新电脑后如何上手（重要）

**Q: 新电脑需要执行 `git init` 吗？**
> ❌ 不需要。`git init` 只在"第一次把普通文件夹变成 git 仓库"时用一次。
> 新电脑直接 `git clone` 拉取，`.git` 版本信息会一起下载，天生就是 git 仓库。

**Q: 新电脑需要安装 Go 环境吗？**
> ✅ 需要。Go 是运行程序的工具，不随代码走。每台要跑代码的电脑都得装一次。

### 新电脑操作步骤

1. **安装 Go**：到 https://go.dev/dl/ 下载安装（Windows 一路下一步）。
   安装后新开终端执行 `go version` 能看到版本号即成功。
2. **克隆项目**（不要 `git init`）：
   ```powershell
   git clone https://github.com/zclupup/learn_go.git
   cd learn_go
   ```
3. **确认 Go 路径**：本项目示例用的是 `D:\software\Go\bin\go.exe`；
   新电脑的安装路径可能不同，若终端里直接 `go` 可用，就用 `go run` 即可，
   否则把命令里的路径换成新电脑的 Go 安装路径。

### 每日同步流程（避免多台电脑进度冲突）

```powershell
# 开始学习前：先拉取最新进度
git pull

# 学完之后：保存并推送
git add .
git commit -m "完成 LessonXX"
git push
```

---

## �📚 知识点总结（Lesson 01–08）

### Lesson 01 — 变量声明
- 三种声明方式：
  - 完整：`var name string = "zhangcl"`
  - 推断：`var name = "zhangcl"`
  - 短声明 ⭐：`name := "zhangcl"`（最常用，只能在函数内用）
- `:=` 第一次声明用；`=` 给已存在的变量改值。
- 基本类型：`string`、`int`、`float64`、`bool`。
- 零值：声明未赋值时有默认值（int→0，string→""，bool→false）。

### Lesson 02 — 输入 + if
- 读取输入：`fmt.Scan(&age)`，**变量前必须加 `&`**（传变量地址，才能改到原变量）。
- `if / else if / else`：条件不用括号，不加冒号，花括号 `{` 跟同行。

### Lesson 03 — 格式化输出
- `fmt.Printf("我叫%s，今年%d岁\n", name, age)`（类似 Python f-string）。
- 占位符：`%s` 字符串、`%d` 整数、`%f` 小数（`%.1f` 保留1位）、`%t` 布尔、`\n` 换行。
- `Print` 不换行 / `Println` 换行且参数间加空格 / `Printf` 格式化。
- `fmt.Sprintf`：不打印，而是**生成字符串存到变量**。

### Lesson 04 — for 循环
- Go 只有 `for`：
  - 计数：`for i := 0; i < 3; i++ {}`
  - 条件（当 while 用）：`for count < 3 {}`
  - 无限：`for {}` + `break`
- `i++` 即 `i = i + 1`；`import "math/rand"` → `rand.Intn(10)` 生成 0~9。

### Lesson 05 — 切片 slice（Go 版 list）
- 创建：`fruits := []string{"苹果", "香蕉"}`
- 追加：`fruits = append(fruits, "葡萄")`（必须用返回值接住）
- 遍历：`for i, v := range fruits {}`
- 切片操作：`numbers[1:4]`（含头不含尾），和 Python 一致。
- 数组长度固定：`var arr [3]int`；切片长度可变。

### Lesson 06 — 函数 func
- 定义：`func add(a int, b int) int { return a + b }`（返回值类型写在最后）。
- 多返回值：`func divide(a, b int) (int, int)`。
- `_` 忽略不想要的返回值（Go 声明的变量必须被使用）。
- ⭐ Go **不自动转换类型**：`float64(sum) / float64(len(nums))`。

### Lesson 07 — 结构体 struct（Go 版 class）
- 定义：`type Person struct { Name string; Age int }`
- 方法：`func (p Person) Introduce() string {}`，`p` 相当于 Python 的 `self`。
- 创建：`p := Person{Name: "zhangcl", Age: 23}`；访问：`p.Name`。
- 首字母大写=公开，小写=私有。
- ⭐ 分行写的复合字面量，**最后一个元素也要加逗号**（因为 Go 自动补分号机制）。

### Lesson 08 — map（Go 版 dict）
- 创建：`ages := map[string]int{"Alice": 30}`（`map[键类型]值类型`）。
- 取值/新增/改：`ages["Bob"] = 26`（存在就改，不存在就增）；取值 `ages["Alice"]`。
- 删除：`delete(ages, "Cathy")`（Python 用 `del`）。
- ⭐ 逗号 ok 判断键是否存在：`v, ok := ages["x"]`，`ok` 为 true/false。
  本质是 map 取值可返回 1 或 2 个值，写几个变量就返回几个（Go 没有 `in`）。
- 遍历：`for k, v := range ages {}`；长度 `len(ages)`。
- ⭐ map 遍历顺序是**随机的**（和 Python dict 保持插入顺序不同）。

### Lesson 09 — 错误处理 error
- ⭐ 思维转变：Go 不用 `try/except`，而是把 `error` 作为**最后一个返回值**。
- 固定套路：出错 `return 0, err`；正常 `return v, nil`（`nil` 类似 Python 的 `None`）。
- 调用后立刻检查：`if err != nil { ... }`（Go 最常见的一行）。
- 创建错误两种方式：
  - `errors.New("固定信息")`
  - `fmt.Errorf("带变量 %d", x)`（可格式化）
- 类型区别：`errors.New(...)` 和 `fmt.Errorf(...)` 返回 `error`；`fmt.Sprintf(...)` 返回 `string`。
- 多包导入用括号：`import ( "errors"; "fmt" )`。
- 看标准库源码：光标放在函数名上按 `F12`（跳转）或 `Alt+F12`（内联预览）。

---

## 🚀 如何运行

Go 运行代码有两种方式，理解区别很有用：

### 方式一：`go run`（直接运行，最常用）
一步到位：临时编译 + 立即执行，**不留下文件**。学习时都用这个。
```powershell
# 把 NN_xxx 换成对应文件
D:\software\Go\bin\go.exe run .\lesson01_variables.go
```

### 方式二：`go build`（先编译成 exe，再执行）
分两步：先生成一个独立的 `.exe` 可执行文件，再运行它。
```powershell
# 1) 编译：生成 lesson01_variables.exe（默认用源文件名）
D:\software\Go\bin\go.exe build .\lesson01_variables.go

# 2) 执行编译产物
.\lesson01_variables.exe
```
> `-o` 可以自定义输出名：`go build -o app.exe .\lesson01_variables.go`
> 编译产物 `.exe` 已在 `.gitignore` 中忽略，不会被提交到 git。

### 两者区别（对比 Python 理解）
| | go run | go build |
|--|--------|----------|
| 作用 | 编译 + 立即运行 | 只编译成 exe（要手动再运行） |
| 产物 | 不生成文件 | 生成 `.exe` 可执行文件 |
| 场景 | 学习/调试（快速看结果） | 交付程序（exe 拷到别的电脑也能跑，无需装 Go） |
| Python 类比 | 类似 `python xxx.py` | 类似打包成 exe |

> 注意：每个 `lessonNN_*.go` 都是独立的 `package main`，用 `go run 单个文件` 运行，
> 不要 `go build` 整个目录（会因为多个 `main` 函数冲突）。
