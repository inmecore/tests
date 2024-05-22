package router

import (
	"github.com/gin-gonic/gin"
	"tests/middleware"
)

type Router struct {
	Api ApiGroup
}

func New() *gin.Engine {
	router := gin.New()

	// middleware
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(middleware.Cors())

	r := new(Router)
	{
		apiGroup := router.Group("/api")

		r.Api.UserRouter.Init(apiGroup)
	}

	return router
}
