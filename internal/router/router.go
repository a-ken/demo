package router

import (
	"github.com/gin-gonic/gin"
	endpoint "ken.com.tw/demo/internal/router/endoint"
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