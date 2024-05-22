package router

import (
	"github.com/gin-gonic/gin"
	"tests/middleware"
	"tests/service"
)

type UserRouter struct{}

func (*UserRouter) Init(router *gin.RouterGroup) {
	g := router.Group("/user")
	{
		g.POST("/login", middleware.Code(), service.UserService.Login)
		g.GET("", middleware.Auth(), service.UserService.Info)
	}
}
