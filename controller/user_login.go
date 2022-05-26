package controller

import (
	"SimpleTikTok/commom"
	service "SimpleTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	commom.Response
	*service.UserLoginResponse
}

func UserLoginHandle(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	login, err := service.QueryUserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: commom.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserRegisterResponse{
		Response: commom.Response{
			StatusCode: 0,
		},
		UserLoginResponse: login,
	})

}
