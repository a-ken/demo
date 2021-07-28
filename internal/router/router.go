package router

import (
	"github.com/gin-gonic/gin"
	"ken.com.tw/demo/internal/router/endpoint"
)

func Initial() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		u := endpoint.UsersEndpoint{}
		u.Initial(v1)
	}
	return router
}