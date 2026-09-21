package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
)

// ========== Lesson 41: 自定义错误 CError 与统一响应 Return ==========
//
// 目标：
// 1. 理解 Go 自定义 error 的实现原理（实现 error 接口）
// 2. 理解 runtime.Caller(1) 如何自动捕获文件名与报错行号
// 3. 理解 service.Return 怎么根据是否为 *CError 提取错误码和错误信息
// 4. 对照 issue_api 的 internal/utils/error.go 和 internal/service/service.go

// ----------------------------------------------------
// 1. 模拟 issue_api 的错误码定义
// ----------------------------------------------------
const (
	SuccessCode      = 0
	ErrRequestParams = 400000 // 参数错误
	ErrTokenError    = 414    // Token 错误
	ErrVinTooShort   = 400001 // VIN 长度不足 17 位
	ErrCommon        = 500500 // 通用系统错误
)

// ----------------------------------------------------
// 2. 模拟 issue_api 的 CError 结构体
// ----------------------------------------------------
type CError struct {
	file   string
	line   int
	errNo  int64
	errMsg string
}

// 实现 Go 的标准 error 接口（只要实现了 Error() string 方法就是 error）
func (e *CError) Error() string {
	return fmt.Sprintf("[CError] File:%s Line:%d No:%d Msg:%s",
		e.file, e.line, e.errNo, e.errMsg)
}

func (e *CError) ErrNo() int64 {
	return e.errNo
}

func (e *CError) ErrMsg() string {
	return e.errMsg
}

// 核心黑魔法：runtime.Caller(1)
// 1 表示获取“调用 NewCError 的那一行代码”的文件名和行号！
func NewCError(errNo int64, msg string) *CError {
	if errNo == 0 || msg == "" {
		return nil
	}
	e := &CError{
		errNo:  errNo,
		errMsg: msg,
	}
	// skip=1: 拿到上一层调用者的文件和行号
	_, e.file, e.line, _ = runtime.Caller(1)
	return e
}

// ----------------------------------------------------
// 3. 模拟 HTTP 响应结构（前端收到的 JSON 结构）
// ----------------------------------------------------
type Response struct {
	ErrNo   int64       `json:"err_no"`
	ErrMsg  string      `json:"err_msg"`
	Results interface{} `json:"results"`
}

// 模拟 service.Return 统一返回处理函数
func MockReturn(results interface{}, err error, fallbackErrCode int64) {
	resp := Response{}

	if err != nil {
		// 类型断言：判断 err 是不是我们自定义的 *CError
		if cErr, ok := err.(*CError); ok {
			// 如果是 CError，精准提取错误码和具体错误信息
			resp.ErrNo = cErr.ErrNo()
			resp.ErrMsg = cErr.ErrMsg()
			// 服务端控制台打印详细排错信息（带文件和行号）
			fmt.Printf("【服务端排错日志】%s\n", cErr.Error())
		} else {
			// 普通 Go error（如 errors.New 或 db 报错），使用兜底错误码
			resp.ErrNo = fallbackErrCode
			resp.ErrMsg = err.Error()
			fmt.Printf("【普通系统错误】%s\n", err.Error())
		}
	} else {
		resp.ErrNo = SuccessCode
		resp.ErrMsg = "success"
	}

	resp.Results = results

	// 序列化成 HTTP JSON 响应
	jsonBytes, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println("【返回给客户端的 HTTP JSON】:")
	fmt.Println(string(jsonBytes))
	fmt.Println()
}

// ----------------------------------------------------
// 4. 模拟两个真实业务场景
// ----------------------------------------------------

// 场景 A：参数校验失败，业务层抛出 CError
func checkVin(vin string) error {
	if len(vin) < 17 {
		// 这一行调用 NewCError，runtime.Caller(1) 会精确记录这一行！
		return NewCError(ErrVinTooShort, "VIN 码格式不合规，不足 17 位")
	}
	return nil
}

// 场景 B：底层库抛出普通原生 error
func queryFromThirdParty() error {
	return errors.New("connection timeout to third-party api")
}

func main() {
	fmt.Println("========== Lesson 41: 自定义错误 CError 与统一响应 Return ==========")
	fmt.Println()

	// 案例 1：成功请求
	fmt.Println(">>> 案例 1：业务成功")
	MockReturn(map[string]string{"user": "zhangcl", "role": "admin"}, nil, ErrCommon)

	// 案例 2：业务校验失败（带行号和特定业务错误码的 CError）
	fmt.Println(">>> 案例 2：参数校验失败（CError）")
	err1 := checkVin("ABC12345") // 故意传个短的
	MockReturn(nil, err1, ErrCommon)

	// 案例 3：底层原生普通 error（触发 fallbackErrCode 兜底）
	fmt.Println(">>> 案例 3：第三方调用失败（普通原生 error）")
	err2 := queryFromThirdParty()
	MockReturn(nil, err2, ErrCommon)
}
