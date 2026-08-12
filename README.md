# learn_go — 从零开始学 Go（Python 背景）

这是一个 Go 语言入门学习项目，学习者有 Python 基础，从最简单的例子逐课学习 Go。

---

## 🤖 给 AI 助手的说明（换电脑后请先读这里）

我是一名 **Go 小白**，只有 **Python 基础**，正在从简单例子逐步学习 Go。请遵守以下规则：

1. **始终用小白视角**：变量、类型、语法等都要解释清楚，尽量对比 Python 来讲。
2. **一课一个主题**：每课新建一个**独立子文件夹** `lessonNN/main.go`（不要放在根目录，
   否则同目录多个 `func main()` 会因"同一文件夹=同一个包"而冲突报错）。
3. **每课流程**：先讲概念 → 写示例代码 → 运行验证 → 给一个小练习。
4. **回答我的疑问**：我经常会问"为什么"，请耐心解释底层原理。
5. **运行命令**：本机 Go 路径为 `D:\software\Go\bin\go.exe`（终端里 `go` 有时不在 PATH）。
   运行单课：`D:\software\Go\bin\go.exe run ./lessonNN`（按目录运行，非单个文件）。
   国内代理已设为 `goproxy.cn`；语言服务器 `gopls` 已安装（跳转/悬停可用）。
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

| 课程 | 主题 | 目录 | 状态 |
|------|------|------|------|
| Lesson 01 | 变量声明与基本类型 | `lesson01_variables/` | ✅ 完成 |
| Lesson 02 | 输入 + if 判断 | `lesson02_input_if/` | ✅ 完成 |
| Lesson 03 | 格式化输出 printf | `lesson03_printf/` | ✅ 完成 |
| Lesson 04 | for 循环 + 猜数字游戏 | `lesson04_for_guess/` | ✅ 完成 |
| Lesson 05 | 数组与切片 slice | `lesson05_slice/` | ✅ 完成 |
| Lesson 06 | 函数 func | `lesson06_function/` | ✅ 完成 |
| Lesson 07 | 结构体 struct | `lesson07_struct/` | ✅ 完成 |
| Lesson 08 | map（字典 dict） | `lesson08_map/` | ✅ 完成 |
| Lesson 09 | 错误处理 error | `lesson09_error/` | ✅ 完成 |
| Lesson 10 | 指针 pointer | `lesson10_pointer/` | ✅ 完成 |
| Lesson 11 | 方法的指针接收者 | `lesson11_method_pointer_receiver/` | ✅ 完成 |
| Lesson 12 | 接口 interface | `lesson12_interface/` | ✅ 完成 |
| Lesson 13 | defer / panic / recover | `lesson13_defer_panic_recover/` | ✅ 完成 |
| Lesson 14 | goroutine + channel 并发入门 | `lesson14_goroutine_channel/` | ✅ 完成 |
| Lesson 15 | *待定* | *待开始* | ⬜ 下一课 |

**当前进度：已完成 Lesson 01–14，下一课是 Lesson 15。后续会面向 Gin 实战方向，先继续夯实 Go 基础和常用后端能力。**

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
4. **设置国内代理**（在中国大陆很重要，否则下载依赖会卡住/失败）：
   ```powershell
   D:\software\Go\bin\go.exe env -w GOPROXY=https://goproxy.cn,direct
   ```
5. **安装 gopls（Go 语言服务器）**：VS Code 的悬停提示、函数签名、跳转定义（`F12`）
   都依赖它。新电脑装完 Go 后需手动装一次：
   ```powershell
   D:\software\Go\bin\go.exe install golang.org/x/tools/gopls@latest
   ```
   - 装好后二进制在 `C:\Users\<用户名>\go\bin\gopls.exe`（即 `GOPATH\bin`）。
   - 验证：`gopls version`（若命令找不到，用全路径 `& "$env:USERPROFILE\go\bin\gopls.exe" version`）。
   - 让 VS Code 识别：`Ctrl+Shift+P` → 运行 `Go: Restart Language Server`（或 `Developer: Reload Window`）。
   - 注意：安装过程可能几分钟（下载+编译），期间终端看似无输出属正常。

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

## �📚 知识点总结（Lesson 01–14）

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

### Lesson 10 — 指针 pointer
- 指针：存放另一个变量**内存地址**的变量（把变量想成盒子，地址就是盒子位置）。
- `&x`：取 x 的地址（找到盒子在哪）。
- `*p`：取地址上的值或修改它（打开盒子）。
- `*int`：表示“一个指向 int 的指针”，常用于函数参数。
- ⭐ 值传递 vs 指针传递：传普通值是复印件（改不到外面）；传指针是真地址（能改原变量）。
- ⭐ 这就是 `fmt.Scan(&age)` 里 `&` 的原因：Scan 需要真地址才能把输入写回 `age`。

### Lesson 11 — 方法的指针接收者
- 方法接收者 `(a Account)` 就是方法所属的结构体。
- 值接收者 `(a Account)`：拿到的是**复印件**，只适合只读方法（如 `Show`），改不到原结构体。
- 指针接收者 `(a *Account)`：拿到的是**真地址**，能修改结构体字段（如 `Deposit`、`Withdraw`）。
- ⭐ 要改结构体就用指针接收者；只读用值接收者（实际代码常统一用指针）。
- 便利细节：调用 `acc.Deposit(50)` 时 Go 自动取地址，无需写 `(&acc).Deposit(50)`。
- ⭐ 顶层函数/方法的书写顺序无关：`main` 放前面或后面都行（与 Python 不同）。
- 模块 vs 包：`learn_go` 是**模块**（整个项目，由 `go.mod` 定义）；每个课程文件夹是一个**包**。

### Lesson 12 — 接口 interface
- 接口定义“能做什么”（一组方法），不关心“是谁”：`type Animal interface { Sound() string }`。
- ⭐ 隐式实现：只要类型拥有接口的所有方法，就**自动**满足该接口（无需声明 implements）。
- 类似 Python 的鸭子类型，但是在**编译时**检查。
- 接口参数：`func describe(a Animal)` 能接收任何满足 `Animal` 的类型（一个函数处理多种类型）。
- 接口切片 `[]Animal{d, c}`：把不同具体类型放在同一切片里。
- 类型断言 `a.(Dog)`：把接口还原为具体类型；用 `dog, ok := a.(Dog)` 的“逗号 ok”形式更安全（不会 panic）。

### Lesson 13 — defer / panic / recover
- `defer`：将语句延迟到函数返回前执行，常用于收尾（关文件/释放资源），类似 Python 的 `finally`/`with`。
- 多个 `defer`：**后进先出**（像栈），最后写的最先执行。
- `panic`：立即中断正常流程（类似未捕获的异常）；如 `10/0` 会自动 panic。
- `recover`：只能在 `defer` 里生效，能抓住 panic 让程序不崩溃。
- ⭐ 地道模式：`defer func() { if r := recover(); r != nil { err = fmt.Errorf(...) } }()` 把 panic 转成 error。
- 注意：正常代码优先返回 `error`；`panic/recover` 只用于真正异常的情况。

### Lesson 14 — goroutine + channel 并发入门
- `go 函数()`：启动一个 goroutine 并发执行；主 goroutine 不会自动等待它。
- `time.Sleep(...)`：本课里只是教学演示，临时给 goroutine 执行时间；正式等待任务完成应优先用 `sync.WaitGroup` 或 channel。
- `make(chan T)`：创建传递 `T` 类型数据的 channel；`ch <- value` 发送，`value := <-ch` 接收。
- channel 像队列：每个发送进去的值只能被成功接收一次，读出来后就从 channel 中消失；不同于 slice/list 可反复读取同一元素。
- 匿名函数可以访问外层作用域变量（闭包）；`go func(s StudentScore) { ... }(student)` 是把当前循环变量作为参数传进去，避免并发里误用循环变量。
- `sync.WaitGroup` 可理解为等待 goroutine 的计数器：`Add(1)` 增加等待数量，`Done()` 完成一个，`Wait()` 阻塞直到计数器变为 0。
- 只有显式调用 `Add(1)` 才会让 WaitGroup 计数器增加；单纯多写一个 `go func()` 不会自动加一。
- `WaitGroup` 要传指针 `&wg`：多个 goroutine 必须操作同一个计数器；传值会拷贝副本，`Done()` 改不到原来的 `wg`。
- `select` 用来同时等待多个 channel 操作；本课用 `select + time.After(...)` 做超时等待，避免错误实验永久卡住。
- `close(ch)` 表示“不再发送新数据”；已在 channel 里的数据仍可继续读。`for value := range ch` 会读到 channel 被关闭且数据读完才结束。
- 两种收集结果方式：已知结果数量时可固定接收 `len(items)` 次，不必 `close`；不想手动数结果时常用 `WaitGroup + close + range`。
- `make([]T, 0, cap)`：创建长度为 0、容量为 `cap` 的切片，适合后续用 `append` 收集结果；`make([]T, n)` 则会直接生成长度为 `n` 的零值元素。
- 练习：用 `StudentScore` 和 `ScoreResult` 结构体并发计算学生成绩结果，为以后 Gin 中结构体响应、JSON 返回和并发处理打基础。

---

---

## 🚀 如何运行

> ⚠️ 重要：本项目每节课放在**独立子文件夹**里（如 `lesson09_error/main.go`）。
> 因为 Go 规定「同一个文件夹 = 同一个包」，若把多个含 `func main()` 的文件放在同一目录会冲突报错。
> 所以每课单独一个目录，各自是独立的 `package main`。

Go 运行代码有两种方式，理解区别很有用：

### 方式一：`go run`（直接运行，最常用）
一步到位：临时编译 + 立即执行，**不留下文件**。学习时都用这个。
```powershell
# 运行某一课：go run ./目录名（编译该目录整个包并执行其 main 函数）
D:\software\Go\bin\go.exe run ./lesson09_error
```
> `go run ./lesson09_error` 编译该目录里所有 `.go` 文件，运行其中的 `func main()`，
> 与文件叫不叫 `main.go` 无关，入口永远是 `main` 函数。

### 方式二：`go build`（先编译成 exe，再执行）
分两步：先生成一个独立的 `.exe` 可执行文件，再运行它。
```powershell
# 1) 编译：在 lesson09_error 目录生成可执行文件
D:\software\Go\bin\go.exe build -o lesson09.exe ./lesson09_error

# 2) 执行编译产物
.\lesson09.exe
```
> 编译产物 `.exe` 已在 `.gitignore` 中忽略，不会被提交到 git。

### 两者区别（对比 Python 理解）
| | go run | go build |
|--|--------|----------|
| 作用 | 编译 + 立即运行 | 只编译成 exe（要手动再运行） |
| 产物 | 不生成文件 | 生成 `.exe` 可执行文件 |
| 场景 | 学习/调试（快速看结果） | 交付程序（exe 拷到别的电脑也能跑，无需装 Go） |
| Python 类比 | 类似 `python xxx.py` | 类似打包成 exe |

### 格式化代码
```powershell
D:\software\Go\bin\go.exe fmt ./...
```
