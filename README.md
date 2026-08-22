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
8. **每课记录我的疑问（重要）**：我学习时会问很多"为什么"，这些是我的真实盲点。
   每节课结束后，把我提出的疑问和你的回答重点，归纳成简明列点，写进该课「知识点总结」
   后面的 **"💬 本课疑问与答疑"** 小节。以后每一课都要这样做，不要只写知识点、漏掉答疑。

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
| Lesson 15 | context 上下文与超时控制 | `lesson15_context/` | ✅ 完成 |
| Lesson 16 | JSON 编码/解码 + struct tag | `lesson16_json/` | ✅ 完成 |
| Lesson 17 | 文件读写 + JSON 文件 | `lesson17_file_json/` | ✅ 完成 |
| Lesson 18 | 包 package、目录结构与模块复习 | `lesson18_package_module/` | ✅ 完成 |
| Lesson 19 | 测试入门 | `lesson19_testing/` | ✅ 完成 |
| Lesson 20 | 标准库 net/http 入门 | `lesson20_http_server/` | ✅ 完成 |
| Lesson 21 | 标准库 HTTP 小项目 | `lesson21_http_task_api/` | ✅ 完成 |
| Lesson 22 | Gin 入门 | `lesson22_gin_intro/` | ✅ 完成 |
| Lesson 23 | Gin 分层小项目 | `lesson23_gin_layered/` | ✅ 完成 |
| Lesson 24 | 数据库入门：database/sql + MySQL | `lesson24_database_sql/` | ✅ 完成 |
| Lesson 25 | interface + repo 模式 | `lesson25_repo_pattern/` | ✅ 完成 |
| Lesson 26 | GORM 入门：连接、模型、CRUD | `lesson26_gorm_intro/` | ✅ 完成 |
| Lesson 27 | 配置文件读取：yaml + struct | `lesson27_yaml_config/` | ✅ 完成 |
| Lesson 28 | 依赖注入思想：手写 NewXxx 组装对象 | `lesson28_di/` `lesson28_di_manual/` | ✅ 完成 |
| Lesson 29 | Go 泛型入门 | `lesson29_generics/` | ✅ 完成 |
| Lesson 30 | sync.RWMutex 读写锁 | `lesson30_rwmutex/` | ✅ 完成 |
| Lesson 31 | log/slog 结构化日志 | `lesson31_slog/` | ✅ 完成 |
| Lesson 32 | Kafka 消费者入门（Sarama） | `lesson32_sarama/` | ⬜ 待复习 |
| Lesson 33 | jsonparser JSON 快速解析 | `lesson33_jsonparser/` | ✅ 完成 |
| Lesson 34 | JWT 解析与校验 | `lesson34_jwt/` | ⬜ 下一课 |

**当前进度：已完成 Lesson 01–31 + 33，下一课是 Lesson 34。已经掌握 Go 基础语法、分层架构、依赖注入、泛型、RWMutex、slog、jsonparser，下一步学习 JWT 解析与校验，对照 issue-consumer 中的 util/jwt.go。**

---

## 🧭 issue-consumer 对照路线

当前学习目标转向看懂 `issue-consumer` 项目。后续课程针对 issue-consumer 中用到的、还未学到的 Go 特性。

### 后续课程路线

| 课程 | 主题 | 对应 issue-consumer |
|------|------|--------------------|
| Lesson 28 | 依赖注入思想：手写 NewXxx 组装对象 | 对照 `cmd/issue-consumer/main.go` 的组装方式 |
| Lesson 29 | Go 泛型入门 | 对照 `HTTPResponse[T any]`、`Result[T any]` |
| Lesson 30 | sync.RWMutex 读写锁 | 对照 `auth.go` 中 token 缓存的并发安全 |
| Lesson 31 | log/slog 结构化日志 | 对照项目中所有 `slog.Info/Error` 调用 |
| Lesson 32 | Kafka 消费者入门（Sarama） | 对照 `consumer_group.go` 的消费模式 |
| Lesson 33 | jsonparser JSON 快速解析 | 对照 `mark_tool.go` 中按路径取 JSON 字段 |
| Lesson 34 | JWT 解析与校验 | 对照 `util/jwt.go` 的 token 过期判断 |
| Lesson 35 | TOML 配置读取 | 对照 `configs/config.go` 的 TOML 解码 |

### 每课固定对照方式

每一课结束后，都追加一个“真实项目对照”小环节：

- 本课学了什么；
- `issue_api` 里哪里用到了；
- 只看 1-2 个真实文件；
- 先建立映射，不急着改业务代码。

例如学完 HTTP 小项目后，只看：

```text
issue_api/internal/server/gin.go
issue_api/internal/service/service.go
```

只关注路由注册、handler 入口、请求参数绑定、统一响应格式，不看复杂业务细节。

学完分层和 repo 后，再对照：

```text
issue_api/internal/service/task.go
issue_api/internal/biz/task/task.go
issue_api/internal/biz/repo/task.go
issue_api/internal/data/issue_tracking_task.go
```

学完 GORM 后，再看 `internal/model` 和 `internal/data`；学完 Wire 后，再看：

```text
issue_api/cmd/issue_api/wire.go
issue_api/cmd/issue_api/wire_gen.go
```

### 追 issue_api 代码时的问题清单

看真实接口时，不要求一开始读懂所有业务，只练习追这条链路：

- 路由在哪里注册？
- 请求参数怎么绑定？
- 调用哪个 service handler？
- service 调哪个 usecase / biz？
- usecase 调哪个 repo？
- repo 接口在哪里定义？
- data 层怎么实现这个 repo？
- 最终查哪张表或调用哪个外部系统？
- 错误怎么包装并返回给前端？
- 有没有 goroutine、cron、Redis 队列这类异步处理？

重要原则：现在不要直接一口气学 Kratos、Wire、GORM、Gin 全套。先用标准库和小项目把底层模型搭起来：函数、结构体、接口、方法、context、HTTP handler、JSON、数据库 CRUD。框架只是把这些东西组织起来。

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

## 📚 知识点总结（Lesson 01–26）

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

### Lesson 15 — context 上下文与超时控制
- `context` 是 Go 里的"任务控制器"，用来传递取消信号、超时信号、请求生命周期。
- `context.Background()`：最基础的空 context，一般作为起点。
- `context.WithTimeout(parent, duration)`：在父 context 基础上加超时功能，返回两个值：新 context 和取消函数。
- `context.WithCancel(parent)`：在父 context 基础上加手动取消功能，只有调用 `cancel()` 才会触发取消。
- ⭐ `WithTimeout` 返回值必须用两个变量接：第一个是新 context，第二个是取消函数 `cancel`。
- ⭐ 第一个参数是"父 context"，可以用前面声明的 `ctx`，也可以直接写 `context.Background()`。
- `ctx.Done()`：返回一个 channel，context 取消/超时时会收到信号。
- `ctx.Err()`：返回取消原因——`context deadline exceeded`（超时）或 `context canceled`（手动取消）。
- ⭐ `WithTimeout` 到时间**自动**触发 `ctx.Done()`；`WithCancel` 需要**手动**调用 `cancel()` 才触发。
- 固定写法：`ctx, cancel := context.WithTimeout(...); defer cancel()` —— 不管成功失败都释放资源。
- `select` 配合 `ctx.Done()` 和 `time.After()` 实现"任务完成 vs 超时取消"的竞争等待。
- Python 对比：像给任务设置 timeout，也像传一个"取消信号"给后台协程。
- ⭐ 匿名结构体切片：`[]struct{name string; delay time.Duration}{...}` 是声明+赋值一体写法，适合临时数据。
- ⭐ 闭包陷阱：`for` 循环里启动 goroutine 时，要把循环变量作为参数传进去，否则所有 goroutine 共享同一个变量。
- ⭐ 并发执行顺序不固定：goroutine 启动顺序由调度器决定，但结果顺序通常按完成时间（先完成先进 channel）。

### Lesson 16 — JSON 编码/解码 + struct tag
- `encoding/json` 是 Go 标准库里处理 JSON 的包：`json.Marshal(v)` 把 Go 数据转成 JSON 字节，`json.Unmarshal(data, &v)` 把 JSON 字节解析进 Go 变量。
- ⭐ Go 的公开/私有规则：标识符首字母大写表示导出（公开），小写表示未导出（私有）。这个规则适用于：变量、常量、函数、类型、结构体字段、方法。`encoding/json` 是另一个包，所以只能访问结构体里的大写字段；小写字段即使写了 `json` tag 也不会被处理。
- `json.Marshal` 返回 `[]byte` 和 `error`；打印 JSON 文本时常用 `string(data)` 把字节切片转成字符串。
- struct tag 写在结构体字段后面，例如 ``Name string `json:"name"` ``；它规定 JSON 对象里的 key 名。Marshal 时用它决定输出 key，Unmarshal 时也用它把 JSON key 匹配回结构体字段。
- tag 只能写在 struct 字段声明上；`json` tag 是给 `encoding/json` 看的，其他包也可能读取自己的 tag，例如 `db`、`form`、`validate`。
- `omitempty` 只影响编码方向：Marshal/MarshalIndent 时字段为零值就不输出；它不影响 Unmarshal。JSON 缺少字段时，结构体字段保持零值；JSON 里有空值时，会照样解析为空值。
- `json:"-"` 表示忽略字段：Marshal 不输出，Unmarshal 也不会写入，适合密码、token 等不该暴露的数据。
- `json.MarshalIndent(对象, 前缀, 缩进)` 用来格式化 JSON；第二个参数 `prefix` 是每一行前面额外加的字符串，通常传 `""`；第三个参数 `indent` 是每一级缩进，常用两个空格 `"  "`。
- `[]byte(s)` 是 Go 的类型转换语法：`目标类型(值)`。例如 `[]byte(productJson)` 表示把 JSON 字符串转成字节切片；反过来 `string(data)` 可以把字节切片转成字符串。
- ⭐ `json.Unmarshal(data, &v)` 第二个参数必须是指针：意思是把字节数据解析后，写入这个指针指向的内存地址，和 `fmt.Scan(&age)` 的原理类似。
- ⭐ 为什么不能传普通变量：Go 函数参数是值传递，传 `user` 进去时函数拿到的是一份拷贝；即使函数内部给这份拷贝赋值，外面的 `user` 也不会变。传 `&user` 才是把原变量地址交给函数，函数才能把解析结果写回原变量。
- `map[string]interface{}` 可用于结构不固定的动态 JSON；Go 1.18+ 可以写成 `map[string]any`，因为 `any` 是 `interface{}` 的别名。
- 动态 JSON 解析到 `interface{}`/`any` 时，数字默认是 `float64`，取值后如果要当具体类型使用，通常需要类型断言。
- JSON、map、struct 的区别：JSON 是一种文本数据格式；map 是 Go 里的键值对容器，类似 Python dict；struct 是固定字段的数据结构，适合定义稳定的请求/响应模型。
- `fmt.Printf("%T", value)` 会打印值的类型；`main.Product` 里的 `main` 是包名 `package main`，不是文件名 `main.go`。`package main` 表示这是可执行程序包，`func main()` 是程序入口。

### Lesson 17 — 文件读写 + JSON 文件
- `os.WriteFile(path, data, perm)` 用来写文件；`data` 类型是 `[]byte`，所以写字符串时常用 `[]byte(text)` 转换。
- `os.ReadFile(path)` 用来一次性读取整个文件，返回 `[]byte` 和 `error`；如果要当文本打印，常用 `string(data)` 转成字符串。
- ⭐ 相对路径是相对于“运行命令时所在的目录”，不是相对于 `main.go` 所在目录。例如从项目根目录执行 `go run ./lesson17_file_json` 时，`lesson17_file_json/students.json` 指向项目根目录下的该文件；如果进入 `lesson17_file_json` 再运行 `go run main.go`，同一个相对路径会变成 `lesson17_file_json/lesson17_file_json/students.json`，可能导致找不到目录。
- 本课推荐从项目根目录运行：`go run ./lesson17_file_json`。
- `0644` 是文件权限：拥有者可读写，其他用户只读。学习阶段先记作“普通文本/JSON 文件常用权限”。
- `json.MarshalIndent(students, "", "  ")` 可以把结构体切片转成格式化 JSON 字节，再配合 `os.WriteFile` 保存到 `.json` 文件。
- `os.ReadFile` 读出 JSON 文件后，可以直接传给 `json.Unmarshal(data, &target)`，解析回结构体或结构体切片。
- 写文件、读文件、JSON 转换都会返回 `error`；每一步都要 `if err != nil { ...; return }`，否则后续代码可能拿到空数据或错误数据继续运行。
- 练习：新增 `Book` 结构体，把 `[]Book` 写入 `book.json`，再读回来解析成 `[]Book` 并遍历打印标题和价格。

### Lesson 18 — 包 package、目录结构与模块复习
- Go 项目可以先理解成三层：`module` 是整个项目，由 `go.mod` 里的 `module learn_go` 定义；`package` 是代码包，通常一个目录就是一个 package；`.go` 文件属于它所在目录的 package。
- 项目内 import 路径不是相对当前文件写的，而是 `go.mod` 的 module 名 + 包目录相对项目根目录的路径。例如 `module learn_go` 加上目录 `lesson18_package_module/student`，得到 `import "learn_go/lesson18_package_module/student"`。
- ⭐ “同一个目录下的 `.go` 文件必须写同一个 package 名”只包含直接放在该目录里的 `.go` 文件，不包含子目录；子目录是新的目录，可以是新的 package。
- `package main` 是特殊包，表示这个目录可以编译成可执行程序；`func main()` 是程序入口。
- 普通 package 不能自己作为程序入口，常用来放可复用代码，再由 `main` 包 import 使用。
- 公开/私有规则仍然适用：首字母大写的类型、函数、方法、字段可以跨 package 使用；首字母小写的只能在当前 package 内部使用。
- 小写私有函数不能被其他 package 直接调用，但可以被同 package 内部的公开函数或公开方法调用；外部再调用这个公开函数，这就是封装。
- `func (c Course) Summary() string {}` 里的 `(c Course)` 叫方法接收者，表示 `Summary` 是 `Course` 类型的方法；`c` 类似 Python 方法里的 `self`。
- Go 命名习惯：项目/module 常用短小、全小写路径；目录和 package 名通常短小、全小写、尽量不用下划线；文件名全小写，多个词常用下划线，如 `student_service.go`；变量/函数/方法用驼峰，如 `studentName`、`NewStudent`；常量也常用驼峰，少用全大写；接口名常用 `Reader`、`Writer` 这种 `-er` 形式。
- 包名一般和目录名一致，虽然不是强制，但强烈建议一致，读代码和 import 时最清楚。

### Lesson 19 — 测试入门
- Go 标准库自带 `testing` 包，不需要安装第三方测试框架；测试文件必须以 `_test.go` 结尾，例如 `main_test.go`。
- `go test ./lesson19_testing` 会编译指定目录的普通 `.go` 文件和 `*_test.go` 文件，并自动查找 `func TestXxx(t *testing.T)` 形式的测试函数。
- `go test -v ./lesson19_testing` 是详细模式，会打印每个测试和子测试的运行结果；`go test ./...` 会运行整个 module 下所有 package 的测试；`go test -run TestMax ./lesson19_testing` 可以只运行名字匹配 `TestMax` 的测试。
- 同一个 package 下，`main_test.go` 可以直接调用 `main.go` 里的 `Add`、`Grade`、`Divide`、`Max`，因为它们都属于 `package main`，不需要包名前缀。
- 同一个 package 的顶层函数、变量、常量、类型名不能重复；不同函数内部的局部变量可以同名；不同类型的方法可以同名。
- `*testing.T` 是测试框架传入的测试控制器，常用 `t.Fatal`、`t.Fatalf`、`t.Error`、`t.Errorf`、`t.Run`、`t.Log` 报告测试状态。
- `t.Run(name, func(t *testing.T) { ... })` 用来创建子测试；匿名函数先作为参数传给 `t.Run`，真正的函数体由 `t.Run` 在运行子测试时调用。
- 表格驱动测试：用 `[]struct{...}{...}` 准备多组输入和期望值，再用 `for _, tt := range tests` 配合 `t.Run(tt.name, ...)` 逐组验证，类似 Python 的参数化测试。
- `go test -v` 输出规则：`=== RUN` 表示开始运行测试或子测试；`--- PASS` 表示通过；`--- FAIL` 表示失败；缩进的 PASS/FAIL 是子测试结果；`文件名:行号` 指出失败位置；最后的 `ok/FAIL package/path 耗时` 是整个 package 的测试结果和总耗时。
- 输出里的耗时如 `(0.00s)` 是单个测试或子测试耗时；最后 `learn_go/lesson19_testing 0.007s` 是整个 package 的测试总耗时。
- 子测试里调用 `t.Fatalf` 只会停止当前子测试，不会阻止同一个表格里的其他子测试继续执行；如果任意子测试失败，父测试和整个 package 最终都会失败。
- 本课 `TestMax` 中保留了一个故意失败的用例：`Max(2, 4)` 实际返回 `4`，但期望值写成 `3`，用于练习阅读失败输出；如果想让全部测试通过，把该用例的 `want` 改为 `4`。

### Lesson 20 — 标准库 net/http 入门
- `net/http` 是 Go 标准库自带的 HTTP 包；`http.HandleFunc(path, handler)` 注册路由，`http.ListenAndServe(addr, nil)` 启动服务。
- handler 形如 `func xxxHandler(w http.ResponseWriter, r *http.Request)`：`w` 用来写 HTTP 响应，`r` 用来读取请求信息。
- `fmt.Fprintln(w, "...")` 是写到指定目标 `w`，在 HTTP 中就是返回给客户端；`fmt.Println("...")` 是写到服务端终端。
- `r.URL.Query().Get("name")` 读取 query 参数，例如 `/hello?name=张三`。
- `json.NewEncoder(w).Encode(data)` 会把 `data` 编成 JSON 并直接写入 `w`；和 `json.Marshal(data)` 不同，它不会把 `[]byte` 返回给你。
- `writeJSON` 里先设置 `Content-Type`，再 `WriteHeader(statusCode)`，最后编码 JSON；如果编码失败，只能用 `log.Println` 记到服务端终端。
- 可以用 `make(chan int)` 这类 JSON 不支持的值模拟编码失败，触发 `log.Println("JSON 响应失败:", err)`。
- `time.Now().Format("2006-01-02 15:04:05")` 使用 Go 的固定参考时间格式化：`2006` 年、`01` 月、`02` 日、`15` 时、`04` 分、`05` 秒。分隔符和排列可改，但不能写成 `YYYY-MM-DD`。
- `if r.Method != http.MethodGet { ...; return }` 是方法限制。写完错误响应后要 `return`，避免后面继续执行正常响应逻辑，导致同一次请求试图写两次响应。
- 在标准库 `net/http` 中，注册 `/` 会兜底匹配没有更具体 handler 的路径；如果想让未知路径返回 404，需要在 `/` handler 中判断 `r.URL.Path != "/"`。
- 常用调试命令：`curl -s http://localhost:8928/products` 请求接口；`curl -s -X POST http://localhost:8928/products` 模拟非 GET；`ss -ltnp | grep :8928` 查看端口；前台服务用 `Ctrl+C` 停止。
- 练习：新增 `/products` 接口，只允许 GET，返回 `[]Product` JSON 数组。

### Lesson 21 — 标准库 HTTP 小项目
- 本课目标：用标准库 `net/http` 做一个内存版任务 API，小项目接口包括 `GET /`、`GET /tasks`、`POST /tasks`、`GET /tasks/{id}`、`PUT /tasks/{id}`、`DELETE /tasks/{id}`。
- `http.HandleFunc(pattern, handler)` 注册路由；标准库路由会按“更具体/更长的匹配优先”选择 handler。`/tasks` 处理列表和创建，`/tasks/` 处理 `/tasks/1`、`/tasks/2` 这类详情路径；`/` 是兜底路由，所以首页 handler 里要判断 `r.URL.Path != "/"` 返回 404。
- `tasksHandler` 里用 `switch r.Method` 区分 `GET` 和 `POST`，类似 Python/Django 里判断 `request.method`；不支持的方法返回 `http.StatusMethodNotAllowed`，也就是 405。
- `sync.Mutex` 是 Go 进程内的“互斥锁”，不是数据库事务隔离级别。它只保证同一个 Go 进程里的多个 goroutine 在进入 `Lock()` 到 `Unlock()` 之间的代码时互斥执行。
- 和 Django/Gunicorn 对比：Gunicorn 多 worker 是多进程，每个进程内存隔离；Go HTTP 服务通常是一个进程，内部每个请求由 goroutine 并发处理。`tasks`、`nextTaskID` 这种包级全局变量在同一个 Go 进程内被所有 goroutine 共享；如果启动多个 Go 进程，它们之间不会共享这些内存变量。
- `Mutex` 不会自动知道自己保护哪些变量，保护范围由程序员决定：凡是在 `tasksMu.Lock()` 和 `tasksMu.Unlock()` 之间读写的共享数据，都算被这把锁保护。本课里约定 `tasksMu` 保护 `tasks` 和 `nextTaskID`。
- 把 `tasks`、`nextTaskID`、`tasksMu` 放在同一个 `var (...)` 块里不是语法要求，而是代码组织习惯：把“共享数据”和“保护它的锁”放近一点，读代码的人更容易知道它们是一组。
- `json.NewDecoder(r.Body).Decode(&req)` 表示从 HTTP 请求体 `r.Body` 读取 JSON，并把解析结果写入 `req`。第二个参数要传 `&req`，因为 Decode 需要修改原变量，原理类似 `fmt.Scan(&age)`。
- `writeJSON` 有两种常见写法：`json.NewEncoder(w).Encode(data)` 是直接编码并写入响应；`json.Marshal(data)` + `w.Write(jsonData)` 是先在内存里生成 JSON 字节，成功后再写响应。后者可以在编码失败时还没写响应头之前改回 500，更方便处理错误。
- `http.Error(w, "JSON 响应失败", http.StatusInternalServerError)` 会给前端返回一个错误响应，状态码是 500，响应体是错误文本。
- HTTP 响应可理解为三部分：`w.Header().Set(...)` 设置响应头，`w.WriteHeader(statusCode)` 设置状态码，`w.Write(...)` 写响应体。如果不显式调用 `WriteHeader`，第一次 `Write` 时 Go 会默认发送 200。
- ⭐ `writeJSON(...)` 只是把响应写出去，不会自动结束当前 handler 函数；如果这是错误分支或提前结束分支，后面必须紧跟 `return`，避免后续代码继续执行导致同一次请求写两次响应。
- `strings.TrimPrefix(r.URL.Path, "/tasks/")` 用来从 `/tasks/1` 中去掉前缀得到字符串 `"1"`；`strconv.Atoi(idText)` 把字符串数字转成 `int`。如果路径是 `/tasks/abc`，转换失败，就应该返回 400。
- `w.Write(jsonData)` 返回两个值：第一个是实际写入的字节数 `int`，第二个是写入错误 `error`。`_, _ = w.Write(jsonData)` 表示两个返回值都暂时忽略。

### Lesson 22 — Gin 入门
- `c.Query("name")` 读取 URL 问号后面的 query 参数，例如 `/hello?name=张三`；`c.Param("id")` 读取路由路径参数，例如路由 `/users/:id` 匹配 `/users/1` 时，`c.Param("id")` 得到 `"1"`。
- Query 参数更像可选筛选条件，如分页、搜索关键字；Param 参数通常是资源标识，如用户 ID、订单 ID。
- 记忆口诀：路径参数 Param 用来“定位资源”，Query 参数用来“筛选/修饰资源”。例如 `/users/12/orders/99?page=2&status=paid` 中，`12` 和 `99` 是路径参数，`page/status` 是 query 参数。
- 路径参数的名字由路由定义决定：`/users/:id` 用 `c.Param("id")`，`/orders/:orderID` 用 `c.Param("orderID")`；一个 URL 里也可以有多个路径参数。
- `*gin.Context` 是 Gin 给每次请求创建的“请求上下文对象”，handler 通过它读取请求参数、请求体、Header，也通过它写 JSON 响应、状态码。
- `gin.Context` 的作用域是“一次 HTTP 请求”：每个请求都有自己的 `c`，请求结束后不要把 `c` 保存到全局变量或后台 goroutine 里长期使用。
- `gin.Default()` 等于创建 Gin 引擎，并默认挂上 `Logger` 和 `Recovery` 中间件。启动时终端会打印请求日志；handler 里 panic 时，`Recovery` 会拦截 panic，避免整个服务进程崩掉。
- 如果不想默认中间件，可以用 `gin.New()` 创建空引擎，再手动 `r.Use(gin.Logger(), gin.Recovery())`。
- Gin 默认是 debug 模式，启动时会打印路由表、debug 提示、proxy 警告等信息；这是正常的开发模式日志。代码里不要随便 `log.Println("r:", r)` 打印整个引擎对象，否则会输出一大串内部结构。
- Gin 模式可以通过环境变量或代码控制：终端临时用 `GIN_MODE=release go run ./lesson22_gin_intro`，代码里用 `gin.SetMode(gin.ReleaseMode)`。开发学习用 debug 方便看路由和日志；线上通常用 release 减少调试输出。
- `r.Group("/api")` 是路由分组，表示给这一组接口统一加前缀。`api := r.Group("/api")` 后写 `api.GET("/status", ...)`，最终路径就是 `/api/status`；真实项目常用 `/api/v1`、`/admin` 这类分组。
- `binding:"required"` 是 Gin/validator 读取的校验 tag，常配合 `c.ShouldBindJSON(&req)` 使用，表示这个字段在请求 JSON 中必填；缺少或零值时，`ShouldBindJSON` 会返回 error。
- `json:"name"` 负责 JSON 字段名映射，`binding:"required"` 负责请求参数校验；它们写在同一个 struct 字段后面，但由不同逻辑读取。
- Gin Logger 请求日志格式如 `[GIN] 2026/08/15 - 17:14:36 | 200 | 136.393µs | 127.0.0.1 | PUT "/users/1"`：依次表示日志来源、请求时间、HTTP 状态码、处理耗时、客户端 IP、请求方法、请求路径。
- `c.JSON(...)` 会把 JSON 响应写给客户端，但不会自动结束当前 handler 函数；如果当前分支写完响应后不应该继续往下执行，就必须手动 `return`，避免同一次请求重复写响应或继续修改数据。
- 本课里的 `users` 和 `nextUserID` 是包级全局变量，在同一个 Go 进程内会被所有请求 goroutine 共享；严格来说并发写入时也应该用 `sync.Mutex` 保护。教学 demo 暂时简化，真实项目要么加锁，要么把数据放进数据库。
- Gin 只是简化 HTTP 写法，不会自动替你解决全局变量并发安全问题；共享内存是否加锁，仍然是程序员负责。

### Lesson 23 — Gin 分层小项目
- 分层职责口诀：`router` 管 URL 和 HTTP 方法，`handler` 管 HTTP 输入/输出，`service` 管业务逻辑，`model` 管数据结构，`main` 管组装和启动。
- `NewUserHandler(userService)` 返回的是 `*UserHandler`，不是 `UserService`；它把 service 包在 handler 里面，让 handler 方法可以通过 `h.userService.Xxx(...)` 调业务逻辑。
- 为什么要 handler 包 service：router 需要注册的是 Gin handler 方法，例如 `userHandler.CreateUser`；service 方法不接收 `*gin.Context`，不能直接当 HTTP handler 用。这样可以让 service 不依赖 Gin，后续更容易测试和复用。
- service 不接收 `*gin.Context` 的好处：单元测试时可以直接构造 `CreateUserRequest` 调 `userService.CreateUser(req)`，不用伪造 HTTP 请求；以后如果换成命令行、定时任务、gRPC、消息队列消费者，也能复用同一套 service 业务逻辑。
- `userService := service.NewUserService()` 创建的是当前 Go 进程内的一份内存数据；如果启动两个服务进程，它们各有各的 `users`，一个进程新增用户不会自动同步到另一个进程。真实项目要共享数据，需要 MySQL/Redis 等外部存储。
- 同一个服务进程、同一个端口下，多次请求会共享同一个 `userService` 实例里的内存数据；第一次 `POST` 新增用户后，第二次 `GET` 列表能看到新增用户。重启服务后内存数据会恢复初始值。
- `PUT /api/v1/users/:id` 同时使用路径参数和请求体：路径里的 `:id` 定位要修改哪个用户，请求 JSON 里的 `name/age` 提供要改成什么值。
- `UpdateUser(c *gin.Context)` 里先用 `c.Param("id")` 取路径参数，再用 `c.ShouldBindJSON(&req)` 读取请求体，最后调用 `service.UpdateUser(id, req)` 执行业务更新。
- Go import 后默认用“包名”访问，不一定是路径最后一段，但通常包名和目录名一致。如果两个导入包的包名相同，需要给其中一个起别名，例如 `userService "learn_go/xxx/service"`，然后用 `userService.NewUserService()` 访问。
- `main()` 启动时只创建一次 `userService`、`userHandler` 和 Gin engine；后续每个请求进来时，Gin 按路由匹配到对应 handler，handler 复用启动时注入的同一个 service 实例。
- 错误在 service 定义、HTTP 状态码在 handler 决定：service 最清楚业务失败原因（如用户不存在、name 为空），handler 负责把业务错误翻译成 HTTP 响应（如 404、400、500）。
- `ListUsers` 里用 `make + copy` 返回切片副本，是为了避免外部拿到 service 内部 `s.users` 的底层数组后直接修改内部数据；这叫保护内部状态/封装。
- 切片本身只是一个描述结构（指向底层数组的指针、长度、容量）。直接 `return s.users` 会把“能访问同一底层数组的切片”交出去；copy 后外部拿到的是独立底层数组，更安全。
- service 层单元测试可以直接 `userService := NewUserService()`，再调用 `CreateUser/GetUser/DeleteUser` 等方法；这正是“业务逻辑不依赖 Gin”的好处，不需要启动 HTTP 服务也能测。
- `_test.go` 文件如果写 `package service`，就和被测代码属于同一个包，可以直接访问 `ErrUserNotFound`、`ErrInvalidName` 这类同包变量；如果写 `package service_test`，就更像外部用户，只能访问大写导出的标识符。
- 判断错误时，`errors.Is(err, ErrUserNotFound)` 比 `err == ErrUserNotFound` 更稳。现在两者都能工作，但以后如果错误被 `fmt.Errorf("xxx: %w", ErrUserNotFound)` 包一层，`errors.Is` 仍然能识别出来。
- 测试函数名要以 `Test` 开头，参数固定是 `t *testing.T`；文件名必须以 `_test.go` 结尾，`go test` 才会自动识别并运行。
- 局部变量通常用小写开头，例如 `userService`；大写开头一般表示导出标识符，更多用于类型、函数、结构体字段等需要给其他包访问的名字。

### Lesson 24 — 数据库入门：database/sql + MySQL
- `database/sql` 是 Go 标准库提供的数据库抽象层；它本身不直接懂 MySQL，需要通过空白导入 `_ "github.com/go-sql-driver/mysql"` 注册 MySQL driver。
- `sql.Open("mysql", dsn)` 主要是创建数据库连接池对象，不一定立刻真正连库；`db.PingContext(ctx)` 才是常用的“确认数据库能连上”的检查。
- DSN 是数据库连接字符串，里面通常包含用户名、密码、地址、端口、数据库名和参数。本课用环境变量 `LEARN_GO_MYSQL_DSN` 保存 DSN，避免把密码写进代码和 git。
- `context.WithTimeout(...)` 可以给数据库操作加超时，避免网络或数据库卡住时程序一直等待。真实后端里，请求传下来的 context 常会一路传到数据库层。
- `db.ExecContext(ctx, sql, args...)` 用来执行不需要返回多行结果的 SQL，例如 `CREATE TABLE`、`INSERT`、`UPDATE`、`DELETE`。
- SQL 里的 `?` 是占位符，参数通过 `ExecContext/QueryContext` 后面的 `args...` 传入；不要自己拼接用户输入到 SQL 字符串里，否则容易产生 SQL 注入风险。
- `result.LastInsertId()` 可以拿到自增主键插入后的 ID；这类似创建一条记录后数据库告诉你新记录的编号。
- `db.QueryContext(...)` 用来查询多行数据，返回 `*sql.Rows`；用完必须 `defer rows.Close()` 释放资源。
- `for rows.Next()` 一行一行读取查询结果，`rows.Scan(&student.ID, &student.Name, ...)` 把当前行的列写入结构体字段。这里必须传指针，因为 Scan 要修改变量。
- 循环结束后还要检查 `rows.Err()`，因为遍历过程中也可能发生错误；这一步容易漏，但真实项目里很重要。
- 本课仍然保留了小测试：空名字直接返回错误，不需要真实数据库也能测。这体现了一个原则：能在进入外部依赖前验证的业务规则，应该尽量单独测。
- 和 `issue_api` 的关系：后续看 `internal/data` 和 GORM 时，可以把 GORM 理解成在这些底层动作之上封装了模型映射、CRUD、事务和查询构造。

### Lesson 25 — interface + repo 模式
- repo 模式的核心是“业务层只依赖接口，不依赖具体数据实现”。本课里 `repo.TaskRepo` 定义 biz 需要的数据能力，`data.memoryTaskRepo` 负责真正保存和查询数据。
- `type TaskRepo interface { ... }` 只写方法签名，不写具体逻辑；它像一份合同：谁拥有这些方法，谁就能被当成 `TaskRepo` 使用。
- `memoryTaskRepo` 是小写类型，外部不能直接依赖它；`NewMemoryTaskRepo()` 返回 `repo.TaskRepo` 接口，让外部只知道“这是一个任务仓库”，不知道它内部是内存实现。
- `TaskUseCase` 里保存的是 `taskRepo repo.TaskRepo`，所以 biz 层可以调用 `Create/List/MarkDone`，但不关心底层是 slice、MySQL、Redis 还是外部 HTTP。
- `main` 的职责是组装对象：先 `data.NewMemoryTaskRepo()`，再 `biz.NewTaskUseCase(taskRepo)`。这就是依赖注入的最小版本：把依赖从外面传进去，而不是在 usecase 内部自己创建。
- 依赖方向要记住：`biz -> repo 接口`，`data -> repo 接口并实现它`，`main -> 组装 biz 和 data`。不要让 biz 直接 import data，否则业务层会被具体存储方式绑死。
- 测试里用 `fakeTaskRepo` 实现同一个 `repo.TaskRepo` 接口，就能在不启动数据库、不启动 HTTP 服务的情况下测试 `TaskUseCase`。这是真实项目里 repo 接口很重要的原因之一。
- 空标题校验发生在 `CreateTask` 里，校验失败时不会调用 repo；测试通过 `fakeRepo.created` 验证这一点，说明 biz 规则可以独立于 data 实现被测试。
- `context.Context` 继续从 usecase 传到 repo，和真实后端请求链路一致：请求进来后，context 会一路传到数据库或外部调用层。
- 对照 `issue_api`：`internal/biz/repo/pack.go` 里的 `PackRepo` 就像本课的 `TaskRepo`；`internal/data/pack.go` 里的 `packRepo` 就像本课的 `memoryTaskRepo`；`NewPackRepo(...) repo.PackRepo` 就像本课的 `NewMemoryTaskRepo() repo.TaskRepo`。
- 看真实项目时先抓链路，不急着看业务细节：service handler 调 biz/usecase，biz/usecase 调 repo 接口，data 层实现 repo 接口，最后才是 GORM/MySQL 或其他外部依赖。

#### 💬 本课疑问与答疑

- **为什么 `NewMemoryTaskRepo()` 返回 `repo.TaskRepo`，但实际 return 的是 `*memoryTaskRepo` 指针，不矛盾吗？**
  不矛盾。接口不是具体对象，而是一组方法要求。`*memoryTaskRepo` 拥有 `Create/GetByID/List/MarkDone` 四个方法，满足 `repo.TaskRepo`，所以能作为接口返回。目的是对上层隐藏具体实现，只暴露接口能力。

- **`&memoryTaskRepo{...}` 里没写 `mu`，那 `mu` 存在吗？**
  存在。结构体未赋值的字段会自动取零值。`sync.Mutex` 的零值就是"未加锁、可直接使用"的状态，等价于写 `mu: sync.Mutex{}`。int→0、string→""、bool→false、slice→nil、Mutex→可用未加锁，都是零值。

- **为什么 `NewTaskUseCase()` 返回 `*TaskUseCase` 而不是 `TaskUseCase`？**
  因为 `return &TaskUseCase{...}` 返回的是地址，类型就是 `*TaskUseCase`，所以签名必须写 `*TaskUseCase`。若写 `TaskUseCase`，`return` 就要去掉 `&` 写 `return TaskUseCase{...}`。返回指针能避免拷贝整个结构体、语义更清楚，也符合后端"构造函数返回对象地址"的习惯。

- **为什么 `NewMemoryTaskRepo` 不带 `*` 也能 return 地址，`NewTaskUseCase` 却要写 `*`？**
  关键区别：`NewMemoryTaskRepo` 返回的是**接口** `repo.TaskRepo`，接口能装下满足它的 `*memoryTaskRepo`；`NewTaskUseCase` 返回的是**指针类型** `*TaskUseCase`。两者不是同一件事。接口靠"方法够不够"判断能否接住某个值，指针类型则直接要求右边是地址。

- **为什么接口能接住 `*memoryTaskRepo`？**
  因为接口只问"这个值的实际类型方法够不够"。本课四个方法都是指针接收者 `func (r *memoryTaskRepo) ...`，所以只有 `*memoryTaskRepo` 实现了接口。方法接收者规则：值接收者 `(r T)` → `T` 和 `*T` 都实现；指针接收者 `(r *T)` → 只有 `*T` 实现。

- **biz 是不是 handler？data 是不是 service？repo 是不是 service 的上层接口？**
  不完全对。更准确是：`handler` 负责 HTTP 输入输出，`biz` 更像 Lesson23 里的 service（业务逻辑），`repo` 是 biz 与 data 之间的接口合同，`data` 才是真正访问数据库/Redis/外部接口的实现。`issue_api` 里 `internal/service` 反而更像 handler/controller 层，`internal/biz` 才是业务 usecase。

- **为什么 data 和 biz 都要写 New 方法？**
  因为它们创建的是两层不同对象：`data.NewMemoryTaskRepo()` 造"数据仓库"，`biz.NewTaskUseCase(taskRepo)` 造"业务对象"并把仓库塞进去。之所以不在 usecase 内部直接 new repo，是为了让 biz 只依赖接口，换掉 data 实现时 biz 不用改。

- **为什么 data 层每个方法开头都写 `if err := ctx.Err(); err != nil { return ... }`？**
  这是"开工前检查"。真实项目里 data 层可能是查 MySQL/Redis/调外部接口，可能慢或卡住。检查 `ctx.Err()` 能感知"请求是否已取消或超时"，超时了就直接返回，不再继续操作。

- **`main` 里为什么写 `context.WithTimeout(..., 3*time.Second)`？从请求到响应必须 3 秒内完成吗？**
  这是给本次业务链路设置一个"最多 3 秒"的生命周期，3 秒是教学里给的宽松值，不是神奇数字。真实 HTTP 场景中，请求 ctx 常来自 `c.Request.Context()`，超时后 ctx 会标记取消，但 Go 不会自动杀代码，需要代码检查 ctx 或把 ctx 传给数据库库来响应。`defer cancel()` 用于提前结束时释放资源。

- **为什么只在 data 层检查 ctx，biz/main 不检查？**
  因为本课里最可能慢、卡、需要取消的是 data 层。biz 层主要是快速业务判断，所以主要把 ctx 往下传。真实项目中若 biz 有长循环或连续外部调用，也应检查 ctx；这不是"只能在 data 层查"。

### Lesson 26 — GORM 入门：连接、模型、CRUD
- GORM 是 Go 常用 ORM，可以把结构体和数据库表建立映射。本课用 SQLite 内存数据库是为了不用安装 MySQL 也能跑通；真实项目 `issue_api` 里通常是 GORM + MySQL。
- `model.Task` 上的 `gorm` tag 描述字段映射，例如 `gorm:"column:title;not null"` 表示这个字段对应数据库列 `title`，并且不能为空。
- `TableName() string` 可以明确指定表名。本课返回 `lesson26_tasks`；`issue_api/internal/model/pack.go` 里 `func (Pack) TableName() string { return "pack" }` 是同一种写法。
- `gorm.Open(sqlite.Open(...), &gorm.Config{})` 是创建 GORM 数据库对象；在 MySQL 场景中会换成 MySQL driver，但后面的 `Create/First/Find/Update` 调用习惯基本一致。
- `AutoMigrate(&model.Task{})` 会根据结构体自动创建或调整表结构。学习项目里很方便；生产项目中是否自动迁移要看团队规范，不能随便改生产表结构。
- `db.WithContext(ctx)` 把请求/任务的 context 传给 GORM，让数据库操作能感知取消或超时；这对应真实后端里 context 一路传到 data 层的习惯。
- `Create(task)` 会插入记录，并把自增 ID 写回 `task.ID`。所以本课 repo 接口的 `Create` 接收 `*model.Task`，方便 GORM 写回 ID。
- `Where("id = ?", id).First(&task)` 查询单条记录；如果没找到，GORM 返回 `gorm.ErrRecordNotFound`。data 层把它转换成自己的 `repo.ErrTaskNotFound`，让 biz 不直接依赖 GORM 错误。
- `Find(&tasks)` 查询多条记录，常配合 `Order/Where/Limit` 等链式方法使用。
- `Model(task).Update("done", true)` 更新单个字段；真实项目里也常见 `Updates(map[string]interface{}{...})` 批量更新字段。
- Lesson26 延续 Lesson25 的依赖方向：`biz` 仍然只依赖 `repo.TaskRepo`，不知道底层是 GORM；`data.gormTaskRepo` 才 import GORM 并实现接口。
- 对照 `issue_api`：`internal/model` 负责结构体和表字段映射，`internal/data` 里常见 `db.Table(...).Where(...).First/Find/Create/Updates`，`internal/biz/repo` 负责定义上层需要的数据能力。
- 测试里使用 SQLite 内存数据库验证 CRUD，比依赖外部 MySQL 更轻；同时仍然能练习 GORM 的真实调用方式。

#### 💬 本课疑问与答疑

- **GORM 是什么？**
  GORM 是一个 ORM（Object-Relational Mapping，对象关系映射）框架。它让你用"操作 Go 结构体"的方式去"操作数据库表"，自动完成结构体与表、字段与列的映射，以及 CRUD 的 SQL 生成，省去手写 SQL 和手动 `Scan`。对比 Lesson24 手写 `SELECT/Scan`，GORM 的 `First(&task)` 会自动把结果填进结构体。

- **`gorm` tag 和 `json` tag 的区别？**
  `gorm:"..."` 是给 GORM 看的，管数据库映射（列名、主键、非空等）；`json:"..."` 是给 `encoding/json` 看的，管 JSON 字段名。两个 tag 各管各的，互不干扰。

- **为什么 `List` 返回 `[]model.Task`，其他方法返回 `*model.Task`？**
  单个对象用指针是为了能用 `nil` 表达"不存在"（如 `GetByID` 找不到返回 `nil`），以及让 GORM 写回自增 ID（`Create` 需要指针才能改 `task.ID`）。列表不需要 `nil` 语义，空切片 `[]model.Task{}` 天然表达"没有数据"，所以用值切片，遍历也更简单。

- **`gorm.Open(sqlite.Open("file:lesson26?mode=memory&cache=shared"), &gorm.Config{})` 是什么意思？**
  拆成两层：外层 `gorm.Open(拨号器, 配置)` 是 GORM 打开数据库的统一入口，返回 `*gorm.DB`；内层 `sqlite.Open("...")` 是 SQLite 的拨号器，告诉 GORM 用 SQLite。字符串里 `mode=memory` 表示用内存数据库（不写磁盘），`cache=shared` 表示多连接共享同一份内存，程序退出数据就没了。`&gorm.Config{}` 是空配置，用默认值。换成 MySQL 就是 `gorm.Open(mysql.Open("user:pass@tcp(...)/db"), &gorm.Config{})`。

- **`AutoMigrate(&model.Task{})` 是什么意思？**
  根据 `Task` 结构体和它的 `gorm` tag 自动创建/调整数据库表。它读取结构体字段和 tag，生成对应的建表语句；表不存在就建，结构有变化就尝试调整。传 `&model.Task{}` 是为了让 GORM 用反射拿到类型信息。注意生产环境不能随便 AutoMigrate，避免意外改动线上表结构。

- **`db.WithContext(ctx)` 是什么？**
  把数据库操作和 context 绑定，让 GORM 在请求取消或超时时能及时停止。这是真实后端里 context 一路传到 data 层的标准做法，固定模式就是 `db.WithContext(ctx).XXX(...).Error`。

- **`.Error` 是什么？**
  GORM 的链式调用返回的是 `*gorm.DB` 对象，不会直接返回 error，要通过最后的 `.Error` 拿错误。成功时 `err == nil`，失败时 `err` 是具体错误。

- **为什么把 `gorm.ErrRecordNotFound` 转成 `repo.ErrTaskNotFound`？**
  因为 biz 层只依赖 `repo.TaskRepo` 接口，不该知道 GORM 的错误。data 层负责把 GORM 错误翻译成 repo 层自己的错误，避免 biz 为了判断错误而 import GORM，从而保持分层干净。

### Lesson 27 — 配置文件读取：yaml + struct
- yaml 是一种可读的配置格式，用缩进表示层级。后端常把数据库地址、Redis 地址、项目参数等写进 yaml，而不是硬编码在代码里，这样改配置不用改代码、不用重新编译。
- 把 yaml 解析到结构体的核心是 `yaml.Unmarshal(data, &config)`：第一个参数是 yaml 文本的字节，第二个参数是要写入的结构体指针。
- 结构体字段用 `yaml:"..."` tag 指定对应 yaml 里的键名。例如 `Addr string \`yaml:"addr"\`` 表示读 yaml 里的 `addr` 字段。
- 嵌套结构体对应 yaml 的层级缩进：顶层 `Config` 有 `server/data/project` 三个字段，`ServerConfig` 里又有 `http/env`，一层套一层。
- 本课结构对齐 `issue_api`：`configs/config.yaml` 对应本课的 `config.yaml`，`conf.Bootstrap` 对应顶层 `Config` 结构体，`server/data/project` 对应嵌套子结构体。
- `issue_api` 用 Kratos 的 config 库 + protobuf 生成 `conf.Bootstrap` 结构体；本课用 `gopkg.in/yaml.v3`，但核心思想一样：yaml 文本 -> 结构体字段。
- 用 `flag.String("conf", "config.yaml", ...)` 指定配置路径，和 `issue_api` 里 `-conf` 参数的用法一致。`flag.Parse()` 之后才能读到参数值。
- ⭐ 相对路径陷阱：`os.ReadFile("config.yaml")` 里的路径是相对"运行命令时所在的目录"，不是相对 main.go 所在目录。从项目根目录运行要用 `-conf lesson27_yaml_config/config.yaml`。

#### 💬 本课疑问与答疑

- **为什么要用 yaml 配置文件？**
  把配置和代码分离。数据库密码、端口、环境参数等会随部署环境变化，写进 yaml 后改配置不需要改代码、不需要重新编译。密码等敏感信息也不该硬编码进代码并提交到 git。

- **`yaml.Unmarshal` 和之前学的 `json.Unmarshal` 是什么关系？**
  原理完全一样，都是"把一种文本格式解析进结构体"，只是格式从 JSON 换成 yaml。JSON 用 `{}` 和 `[]`，yaml 用缩进和 `key: value`。两者都需要传指针 `&config`，都靠 tag 指定字段名映射。

- **`yaml:"addr"` 这个 tag 和之前的 `json:"addr"`、`gorm:"column:addr"` 有什么区别？**
  都是 struct tag，但给不同的库读：`yaml` tag 给 yaml 库读，`json` tag 给 encoding/json 读，`gorm` tag 给 GORM 读。同一个字段可以同时有多个 tag，各管各的，互不影响。

- **`flag.String("conf", "config.yaml", ...)` 是什么意思？**
  定义一个命令行参数叫 `conf`，默认值是 `"config.yaml"`，第三个字符串是帮助说明。`flag.Parse()` 会解析命令行参数，之后用 `*conf` 拿到值（注意是指针）。这样运行时可以用 `-conf 路径` 覆盖默认值。

### Lesson 28 — 依赖注入思想

- 依赖注入的核心：**每个对象通过构造函数声明"我需要什么"，由外部（main）把依赖传进去**。
- 反例：对象内部自己 `new` 依赖，换实现必须改业务代码。
- 正例：`NewXxx(dep1, dep2)` 传入依赖，换实现只改 main 的组装代码。
- 组合根（Composition Root）：整个程序**唯一**创建和组装对象的地方，一般是 `main` 函数。
- 接口定义在"消费者"这一侧：`Logger` 在 biz 层定义（biz 要用日志），`ArticleRepo` 在 repo 层定义（repo 提供数据能力）。
- 对照 issue-consumer：`main.go` 里 `NewMarkToolHandler(config)` + `NewMConsumerLauncher(cg, handler, topics)` 就是 DI。

### Lesson 29 — Go 泛型

- 泛型解决的核心问题：**减少重复代码**——同一套逻辑，多种类型。
- 泛型函数：`func Sum[T int | float64](nums []T) T`，`[T int | float64]` 是类型参数。
- 泛型结构体：`type Box[T any] struct { value T }`，`T` 在字段中使用。
- 构造函数也要带泛型：`func NewBox[T any](value T) Box[T]`。
- 方法接收者引用已声明的 `T`，不能重新声明：`(b Box[T]) Get() T`（不是 `Box[T any]`）。
- 调用时省略类型参数：`Sum([]int{1,2,3})` 编译器推断 `T = int`。
- 多个类型参数：`Pair[K comparable, V any]`。
- `comparable` 约束：表示支持 `==` 和 `!=`，用于 `Contains` 这类需要比较的函数。
- 自定义约束：`type Number interface { ~int | ~float64 }`，`~` 表示底层类型也匹配。
- 嵌套泛型：`HTTPResponse[HttpBatchImportIssueResp]` 表示外层 T 替换为具体类型。
- 函数作为参数：`func MapSlice[T any, U any](items []T, fn func(T) U) []U`，`fn` 是函数类型参数。
- 对照 issue-consumer：`HTTPResponse[T any]` 的 `T` 就是 `Results` 字段的类型；`httpz.PostJsonBodyFormattedWithCtx[HTTPResponse[Xxx]](...)` 是嵌套泛型调用。

### Lesson 30 — sync.RWMutex 读写锁

- `Mutex` 的缺陷：读操作和读操作之间没有冲突，但 Mutex 强制串行化，浪费性能。
- `RWMutex` 区分读锁和写锁：`RLock()` 读锁可并发，`Lock()` 写锁独占。
- 规则：读锁互斥写锁（写等读、读等写），但不互斥读锁（可以同时读）。
- 适用场景：**读多写少**，如 token 缓存、配置缓存、计数器等。
- 核心对比：`sync.Mutex` 读写都 Lock，`sync.RWMutex` 读用 RLock（可并发）写用 Lock（独占）。
- 对照 issue-consumer：`auth.go` 中 `rwmux.RLock()` 读缓存 token，`rwmux.Lock()` 写缓存 token。

### Lesson 31 — log/slog 结构化日志

- slog 是 Go 1.21+ 标准库，核心思想：**日志是键值对，不是字符串**。
- 四种级别：`Debug < Info < Warn < Error`，`HandlerOptions.Level` 控制输出级别。
- 关键字段：`slog.String("key", "value")`、`slog.Int("key", 42)`、`slog.Int64("key", int64)`、`slog.Float64("key", 3.14)`、`slog.Bool("key", true)`、`slog.Time("key", time.Now())`、`slog.Any("key", value)`。
- `TextHandler`：人读的文本格式，适合开发环境。
- `JSONHandler`：机器解析的 JSON 格式，适合生产环境。
- `AddSource: true`：在日志中输出文件名和行号，方便调试定位。
- `slog.SetDefault(logger)`：设为全局默认 Logger，项目中任何地方直接 `slog.Info(...)` 使用。
- 对照 issue-consumer：`main.go` 第 24-34 行创建 JSONHandler + AddSource + SetDefault。

### Lesson 32 — Kafka 消费者入门（Sarama 模式）

- Kafka 是消息队列：生产者 → Topic → 消费者组。
- Sarama 的 `ConsumerGroupHandler` 接口有 3 个方法：`Setup`、`Cleanup`、`ConsumeClaim`。
- Sarama 回调模式：分配分区 → Setup → ConsumeClaim（逐条处理消息）→ Cleanup。
- 消息通过 `claim.Messages()` channel 传递，消费者从 channel 读消息。
- 对照 issue-consumer：`consumer_group.go` 的 `Run` 方法调用 `client.Consume(ctx, topics, handler)`。

### Lesson 33 — jsonparser JSON 快速解析

- `jsonparser` 是第三方 JSON 解析库，**不用定义结构体，按路径直接取值**。
- 取字符串：`jsonparser.GetString(data, "key")`，支持嵌套路径：`GetString(data, "a", "b", "c")`。
- 取其他类型：`GetInt`、`GetFloat`、`GetBoolean`。
- 遍历数组：`jsonparser.ArrayEach(data, func(value []byte, ...) { ... }, "数组键")`。
- 判断字段存在：`jsonparser.Get` 返回 `value, dataType, offset, err`。
- 对比 `json.Unmarshal`：结构固定/字段多时用 Unmarshal，只需几个字段/动态键名时用 jsonparser。
- 对照 issue-consumer：`mark_tool.go` 中用 `GetString` 取中文键名字段，用 `ArrayEach` 遍历数组。

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
