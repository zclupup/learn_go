package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"learn_go/lesson23_gin_layered/model"
	"learn_go/lesson23_gin_layered/service"
)

// UserHandler 负责 HTTP 层：读参数、读请求体、写响应。
type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users := h.userService.ListUsers()
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, ok := h.parseID(c)
	if !ok {
		return
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		h.writeServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{Message: "请求体不是合法 JSON，或缺少必填字段"})
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		h.writeServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, ok := h.parseID(c)
	if !ok {
		return
	}

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{Message: "请求体不是合法 JSON，或缺少必填字段"})
		return
	}

	user, err := h.userService.UpdateUser(id, req)
	if err != nil {
		h.writeServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, ok := h.parseID(c)
	if !ok {
		return
	}

	if err := h.userService.DeleteUser(id); err != nil {
		h.writeServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{Message: "用户已删除"})
}

func (h *UserHandler) SearchUsers(c *gin.Context) {
	minAgeStr := c.Query("min_age")

	minAge, err := strconv.Atoi(minAgeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{Message: "min_age 必须是数字"})
		return
	}

	users := h.userService.SearchUsersByMinAge(minAge)
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) parseID(c *gin.Context) (int, bool) {
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{Message: "用户 id 必须是数字"})
		return 0, false
	}
	return id, true
}

func (h *UserHandler) writeServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrUserNotFound):
		c.JSON(http.StatusNotFound, model.APIResponse{Message: "用户不存在"})
	case errors.Is(err, service.ErrInvalidName):
		c.JSON(http.StatusBadRequest, model.APIResponse{Message: "name 不能为空"})
	default:
		c.JSON(http.StatusInternalServerError, model.APIResponse{Message: "服务内部错误"})
	}
}
