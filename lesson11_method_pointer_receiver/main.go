package main

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

// ======== 值接收者：拿到的是结构体的"复印件" ========
// 只读操作用值接收者就够了
func (a Account) Show() string {
	return fmt.Sprintf("%s 的余额: %d", a.Owner, a.Balance)
}

// ======== 值接收者尝试修改（改不到原来的）========
// a 是复印件，改了外面看不到
func (a Account) DepositWrong(amount int) {
	a.Balance += amount // 只改了复印件
}

// ======== 指针接收者：拿到的是结构体的"真地址" ========
// 想修改结构体字段，必须用指针接收者 (a *Account)
func (a *Account) Deposit(amount int) {
	a.Balance += amount // 改到真结构体
}

func (a *Account) Withdraw(amount int) error {
	if amount > a.Balance {
		return fmt.Errorf("余额不足: 当前 %d, 想取 %d", a.Balance, amount)
	}
	a.Balance -= amount
	return nil
}

// practice
func (a *Account) Rename(newOwner string) {
	a.Owner = newOwner
}

func main() {
	acc := Account{Owner: "zhangcl", Balance: 100}
	fmt.Println(acc.Show())

	// 值接收者：改不到原来的
	acc.DepositWrong(50)
	fmt.Println("DepositWrong 后:", acc.Show()) // 还是 100

	// 指针接收者：能改到原来的
	acc.Deposit(50)
	fmt.Println("Deposit 后:", acc.Show()) // 150

	// 指针接收者 + error
	err := acc.Withdraw(200)
	if err != nil {
		fmt.Println("取款失败:", err)
	}

	err = acc.Withdraw(30)
	if err != nil {
		fmt.Println("取款失败:", err)
	} else {
		fmt.Println("取款成功:", acc.Show()) // 120
	}

	// practice: Rename the account owner
	acc.Rename("lihua")
	fmt.Println("after rename result:", acc.Show())
}
