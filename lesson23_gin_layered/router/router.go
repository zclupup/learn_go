package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learn_go/lesson23_gin_layered/handler"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/users", userHandler.ListUsers)
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/search", userHandler.SearchUsers)

		api.GET("/users/:id", userHandler.GetUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)

	}

	return r
}
