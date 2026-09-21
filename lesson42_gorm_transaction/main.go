package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ========== Lesson 42: GORM 事务控制与批量操作 ==========
//
// 目标：
// 1. 理解为什么需要事务（ACID 中的原子性：要么全成功，要么全回滚）
// 2. 掌握 GORM 的两种事务写法：
//    - 方式 A：闭包自动事务 db.Transaction(func(tx *gorm.DB) error)
//    - 方式 B：手动显式事务 tx.Begin() + defer 回滚/提交（对照 issue_api 真实写法）
// 3. 掌握 GORM 批量插入（CreateInBatches）与性能优化
// 4. 对照 issue_api 的 internal/data/transaction.go 和 internal/biz/task/task.go

// 模拟两个表：账户表与转账记录表
type Account struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:50"`
	Balance int    // 余额（元）
}

type TransferLog struct {
	ID     uint `gorm:"primaryKey"`
	FromID uint
	ToID   uint
	Amount int
}

func setupTestDB() *gorm.DB {
	// 使用内存 sqlite 模拟数据库，运行零依赖且速度极快
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 静默日志，输出更清爽
	})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移建表
	_ = db.AutoMigrate(&Account{}, &TransferLog{})

	// 初始化测试数据：张三 100 元，李四 50 元
	db.Create(&Account{ID: 1, Name: "张三", Balance: 100})
	db.Create(&Account{ID: 2, Name: "李四", Balance: 50})

	return db
}

// ----------------------------------------------------------------
// 模式 1：GORM 官方推荐的闭包事务 db.Transaction(...)
// 特点：闭包返回 error 就自动 Rollback，返回 nil 就自动 Commit，极其安全防漏！
// ----------------------------------------------------------------
func transferWithClosure(db *gorm.DB, fromID, toID uint, amount int, mockError bool) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. 扣除转出方余额（注意：所有的 SQL 操作必须用 tx，不能用 db！）
		res := tx.Model(&Account{}).Where("id = ? AND balance >= ?", fromID, amount).
			Update("balance", gorm.Expr("balance - ?", amount))
		if res.RowsAffected == 0 {
			return errors.New("扣款失败：余额不足或账户不存在")
		}

		// 2. 模拟突然发生的业务中断异常（比如中途调用外部接口失败）
		if mockError {
			return errors.New("模拟转账中途网络中断故障！")
		}

		// 3. 增加收款方余额
		if err := tx.Model(&Account{}).Where("id = ?", toID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		// 4. 插入转账日志
		if err := tx.Create(&TransferLog{FromID: fromID, ToID: toID, Amount: amount}).Error; err != nil {
			return err
		}

		// 返回 nil，GORM 自动提交事务
		return nil
	})
}

// ----------------------------------------------------------------
// 模式 2：手动事务 + defer 统一提交/回滚
// 【重点】：这完全是 issue_api 在 internal/biz/task/task.go 中的实战写法！
// ----------------------------------------------------------------
func transferWithManualTx(db *gorm.DB, fromID, toID uint, amount int, mockError bool) (err error) {
	// 1. 手动开启事务
	tx := db.Begin()

	// 2. 利用 defer 和具名返回值 err，在函数退出时统一检查
	defer func() {
		if err != nil {
			fmt.Println("   [事务警报] 检测到 error，执行 Rollback 回滚！")
			tx.Rollback()
			return
		}
		fmt.Println("   [事务完成] 无任何 error，执行 Commit 提交！")
		tx.Commit()
	}()

	// 3. 业务逻辑操作（全部使用 tx）
	res := tx.Model(&Account{}).Where("id = ? AND balance >= ?", fromID, amount).
		Update("balance", gorm.Expr("balance - ?", amount))
	if res.RowsAffected == 0 {
		err = errors.New("扣款失败：余额不足")
		return err
	}

	if mockError {
		err = errors.New("模拟后续操作抛出严重错误")
		return err // 提前 return，触发 defer 中的 Rollback
	}

	err = tx.Model(&Account{}).Where("id = ?", toID).
		Update("balance", gorm.Expr("balance + ?", amount)).Error
	if err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------
// 模式 3：GORM 批量插入 (CreateInBatches)
// ----------------------------------------------------------------
func demonstrateBatchInsert(db *gorm.DB) {
	fmt.Println("━━━━━ 模式 3: GORM 高性能批量插入 (CreateInBatches) ━━━━━")

	// 构造 10 条测试数据
	var users []Account
	for i := 10; i < 20; i++ {
		users = append(users, Account{
			ID:      uint(i),
			Name:    fmt.Sprintf("用户_%d", i),
			Balance: 1000,
		})
	}

	// 单条插入 vs 批量插入：
	// 如果用 for 循环 db.Create，要向数据库发 10 次请求；
	// 用 CreateInBatches，每次打包 batchSize 条拼成一条批量 INSERT 语句：
	// INSERT INTO accounts (...) VALUES (...), (...), (...)
	batchSize := 5
	result := db.CreateInBatches(users, batchSize)
	if result.Error != nil {
		fmt.Printf("批量插入失败: %v\n", result.Error)
		return
	}
	fmt.Printf("分批插入完成！总共受影响行数: %d\n\n", result.RowsAffected)
}

func printBalances(db *gorm.DB, title string) {
	var accounts []Account
	db.Find(&accounts)
	fmt.Printf("=== %s ===\n", title)
	for _, a := range accounts {
		if a.ID <= 2 {
			fmt.Printf("  账户 %d (%s): 余额 %d 元\n", a.ID, a.Name, a.Balance)
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("========== Lesson 42: GORM 事务控制与批量操作 ==========")
	fmt.Println()

	db := setupTestDB()
	printBalances(db, "初始状态")

	// 测试 1：闭包事务成功转账 20 元
	fmt.Println(">>> 测试 1: 闭包事务执行正常转账 20 元（张三 -> 李四）")
	err := transferWithClosure(db, 1, 2, 20, false)
	fmt.Printf("转账结果: err = %v\n", err)
	printBalances(db, "测试 1 后状态（张三扣20，李四加20）")

	// 测试 2：闭包事务中途失败（回滚验证）
	fmt.Println(">>> 测试 2: 闭包事务中途出现故障，验证原子回滚")
	err = transferWithClosure(db, 1, 2, 20, true) // mockError = true
	fmt.Printf("转账结果: err = %v\n", err)
	printBalances(db, "测试 2 后状态（张三金额没有被扣减，数据完好无损！）")

	// 测试 3：手动事务 + defer 写法（issue_api 模式）
	fmt.Println(">>> 测试 3: 手动事务 + defer 回滚（对照 issue_api 真实用法）")
	err = transferWithManualTx(db, 1, 2, 30, true) // 故意模拟失败
	fmt.Printf("转账结果: err = %v\n", err)
	printBalances(db, "测试 3 后状态（触发 Rollback，张三依然完好）")

	// 测试 4：批量插入
	demonstrateBatchInsert(db)
}
