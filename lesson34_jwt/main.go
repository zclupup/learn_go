package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ================================================================
// Lesson 34 — JWT 解析与校验
// ================================================================
//
// JWT（JSON Web Token）是一种 token 格式，由三部分组成：
//   header.payload.signature
//
// 各部分：
//   header：{"alg": "HS256", "typ": "JWT"}  → 加密算法
//   payload：{"sub": "123", "name": "zhangcl", "exp": 1700000000}  → 数据 + 过期时间
//   signature：对 header + payload 的签名，防止篡改
//
// 对照 issue-consumer 中的 JWT 用法：
//   util/jwt.go → IsTokenExpired 函数：判断 token 是否过期
//
// ================================================================

func main() {
	fmt.Println("========== Lesson 34: JWT 解析与校验 ==========")
	fmt.Println()

	// ============ 1. 生成一个 JWT（模拟登录服务返回的 token）============
	fmt.Println("--- 1. 生成一个 JWT ---")

	// 创建一个 token，1 小时后过期
	claims := jwt.MapClaims{
		"sub":  "1001",
		"name": "zhangcl",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("my-secret-key"))
	fmt.Printf("生成的 JWT:\n%s\n\n", signedToken)

	// 拆开看三部分
	fmt.Println("JWT 三部分（用 . 分隔）：")
	parts := splitToken(signedToken)
	for i, p := range parts {
		fmt.Printf("  第 %d 部分: %s\n", i+1, p)
	}
	fmt.Println()

	// ============ 2. 解析 JWT（不验证签名）============
	fmt.Println("--- 2. 解析 JWT（不验证签名）---")

	// 和 issue-consumer 一样，用 ParseUnverified
	// 为什么？因为只是查 token 过期时间，不需要验证签名
	token, _, err := new(jwt.Parser).ParseUnverified(signedToken, jwt.MapClaims{})
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Printf("sub:  %v\n", claims["sub"])
		fmt.Printf("name: %v\n", claims["name"])
		fmt.Printf("iat:  %v (时间戳)\n", claims["iat"])
		fmt.Printf("exp:  %v (时间戳)\n", claims["exp"])
	}
	fmt.Println()

	// ============ 3. 判断 token 是否过期 ============
	fmt.Println("--- 3. 判断 token 是否过期 ---")

	// 对应 issue-consumer 的 IsTokenExpired 函数
	aliveToken := signedToken
	fmt.Printf("有效 token 是否过期: %v (应该为 false)\n", IsTokenExpired(aliveToken))

	// 创建一个已过期的 token（过期时间设为过去）
	expiredClaims := jwt.MapClaims{
		"exp": time.Now().Add(-1 * time.Hour).Unix(), // 1 小时前已过期
	}
	expiredToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims).SignedString([]byte("key"))
	fmt.Printf("过期 token 是否过期: %v (应该为 true)\n", IsTokenExpired(expiredToken))

	// 无效 token
	fmt.Printf("无效 token 是否过期: %v (应该为 true)\n", IsTokenExpired("invalid.token.here"))
	fmt.Println()

	// ============ 4. 对照 issue-consumer 真实代码 ============
	fmt.Println("--- 4. 对照 issue-consumer 真实代码 ---")
	fmt.Println("util/jwt.go 中的 IsTokenExpired 函数：")
	fmt.Println()
	fmt.Println("  func IsTokenExpired(tokenString string) bool {")
	fmt.Println("      token, _, err := new(jwt.Parser).ParseUnverified(")
	fmt.Println("          tokenString, jwt.MapClaims{})")
	fmt.Println("      if err != nil {")
	fmt.Println("          return true    // 解析失败视为过期")
	fmt.Println("      }")
	fmt.Println("      if claims, ok := token.Claims.(jwt.MapClaims); ok {")
	fmt.Println("          if exp, ok := claims[\"exp\"].(float64); ok {")
	fmt.Println("              now := float64(time.Now().Unix())")
	fmt.Println("              return now > exp")
	fmt.Println("          }")
	fmt.Println("      }")
	fmt.Println("      return false    // 没有 exp 字段视为永久有效")
	fmt.Println("  }")
	fmt.Println()

	// ============ 5. JWT 生命周期 ============
	fmt.Println("--- 5. JWT 在 issue-consumer 中的完整流程 ---")
	fmt.Println()
	fmt.Println("  auth.go: GetAuthToken")
	fmt.Println("    ↓ 请求 auth API 获取 token")
	fmt.Println("    ↓ 缓存到 cacheToken（用 RWMutex 保护）")
	fmt.Println("    ↓")
	fmt.Println("  auth.go: 每次请求前读缓存 token")
	fmt.Println("    ↓")
	fmt.Println("  issue_tracking_result.go: 在请求头中带上 token")
	fmt.Println("    header[\"Authorization\"] = fmt.Sprintf(\"Bearer %s\", token)")
	fmt.Println("    ↓")
	fmt.Println("  issue_api 收到请求后验证 token")
	fmt.Println("    ↓ 如果 token 过期，返回 401（unAuthCode）")
	fmt.Println("    ↓ refresh=true，重新获取 token 再试一次")
	fmt.Println()

	// ============ 6. 总结 ============
	fmt.Println("========== 总结 ==========")
	fmt.Println()
	fmt.Println("JWT 结构：header.payload.signature")
	fmt.Println("核心字段：exp（过期时间）、sub（用户标识）、iat（签发时间）")
	fmt.Println()
	fmt.Println("issue-consumer 中的 JWT 用途：")
	fmt.Println("  1. main 启动时，通过 GetAuthToken 获取 token")
	fmt.Println("  2. 每次请求 issue-api 时，带上 token 做鉴权")
	fmt.Println("  3. 如果 token 过期，重新获取再重试")
	fmt.Println("  4. IsTokenExpired 只是检查 exp 字段，不验证签名")
}

// IsTokenExpired 判断 JWT 是否过期。
// 和 issue-consumer util/jwt.go 的实现完全一致。
func IsTokenExpired(tokenString string) bool {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return true // 解析失败视为过期或无效
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			now := float64(time.Now().Unix())
			return now > exp
		}
	}
	return false // 没有 exp 字段，视为永久有效
}

// splitToken 把 JWT 按 . 分割成三部分，方便查看。
func splitToken(token string) []string {
	parts := make([]string, 0, 3)
	current := ""
	for _, ch := range token {
		if ch == '.' {
			parts = append(parts, current)
			current = ""
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}