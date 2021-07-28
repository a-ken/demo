package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ken.com.tw/demo/internal/service"
)

type UserSignin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersEndpoint struct {
}

func (u *UsersEndpoint) Initial(rg *gin.RouterGroup) {
	r := rg.Group("/users")
	r.POST("/signin", u.signin)
}

func (u *UsersEndpoint) signin(c *gin.Context) {
	var json UserSignin
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	userService := service.UserService{}
	if _, err := userService.Signin(json.Username, json.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
	return
}