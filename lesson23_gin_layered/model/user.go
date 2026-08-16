package model

// User 是对外返回的用户模型。
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// CreateUserRequest 是创建用户时的请求体。
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

// UpdateUserRequest 是更新用户时的请求体。
type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

// APIResponse 是通用消息响应。
type APIResponse struct {
	Message string `json:"message"`
}
