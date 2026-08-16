package service

import (
	"errors"
	"testing"

	"learn_go/lesson23_gin_layered/model"
)

func TestCreateUser(t *testing.T) {
	userService := NewUserService()
	// 测试创建用户
	user, err := userService.CreateUser(model.CreateUserRequest{Name: "Alice", Age: 30})
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	if user.ID != 3 || user.Name != "Alice" || user.Age != 30 {
		t.Fatalf("CreateUser returned unexpected user: %+v", user)
	}
}

func TestCreateUserInvalidName(t *testing.T) {
	userService := NewUserService()

	_, err := userService.CreateUser(model.CreateUserRequest{Name: "   ", Age: 18})
	if !errors.Is(err, ErrInvalidName) {
		t.Fatalf("CreateUser expected ErrInvalidName, got: %v", err)
	}
}

func TestGetUserNotFound(t *testing.T) {
	userService := NewUserService()
	// 测试获取不存在的用户
	_, err := userService.GetUser(999)
	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("GetUser expected ErrUserNotFound, got: %v", err)
	}
}

func TestSearchUsersByMinAge(t *testing.T) {
	userService := NewUserService()

	users := userService.SearchUsersByMinAge(21)
	if len(users) != 1 {
		t.Fatalf("SearchUsersByMinAge expected 1 user, got: %d", len(users))
	}
	if users[0].ID != 2 || users[0].Name != "李四" || users[0].Age != 25 {
		t.Fatalf("SearchUsersByMinAge returned unexpected user: %+v", users[0])
	}
}

func TestDeleteUser(t *testing.T) {
	userService := NewUserService()

	if err := userService.DeleteUser(1); err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}
	_, err := userService.GetUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("GetUser after DeleteUser expected ErrUserNotFound, got: %v", err)
	}
}
