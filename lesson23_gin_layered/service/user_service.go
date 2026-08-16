package service

import (
	"errors"
	"strings"
	"sync"

	"learn_go/lesson23_gin_layered/model"
)

var ErrUserNotFound = errors.New("用户不存在")
var ErrInvalidName = errors.New("name 不能为空")

// UserService 负责用户相关业务逻辑。
// 这一层不直接关心 HTTP，也不直接操作 gin.Context。
type UserService struct {
	mu     sync.Mutex
	users  []model.User
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users: []model.User{
			{ID: 1, Name: "张三", Age: 20},
			{ID: 2, Name: "李四", Age: 25},
		},
		nextID: 3,
	}
}

func (s *UserService) ListUsers() []model.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]model.User, len(s.users))
	copy(result, s.users)
	return result
}

func (s *UserService) SearchUsersByMinAge(minAge int) []model.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []model.User
	for _, user := range s.users {
		if user.Age >= minAge {
			result = append(result, user)
		}
	}
	return result
}

func (s *UserService) GetUser(id int) (model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.ID == id {
			return user, nil
		}
	}
	return model.User{}, ErrUserNotFound
}

func (s *UserService) CreateUser(req model.CreateUserRequest) (model.User, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return model.User{}, ErrInvalidName
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	user := model.User{
		ID:   s.nextID,
		Name: name,
		Age:  req.Age,
	}
	s.nextID++
	s.users = append(s.users, user)
	return user, nil
}

func (s *UserService) UpdateUser(id int, req model.UpdateUserRequest) (model.User, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return model.User{}, ErrInvalidName
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.users {
		if s.users[i].ID == id {
			s.users[i].Name = name
			s.users[i].Age = req.Age
			return s.users[i], nil
		}
	}
	return model.User{}, ErrUserNotFound
}

func (s *UserService) DeleteUser(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return ErrUserNotFound
}
