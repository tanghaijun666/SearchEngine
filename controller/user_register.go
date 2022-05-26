package controller

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRegisterResponse struct {
	commom.Response
	*service.UserLoginResponse
}

func UserRegisterHandel(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//调用service ,获取registerResponse
	registerRespon, err := service.PostUserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: commom.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	} else {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: commom.Response{
				StatusCode: 0,
			},
			UserLoginResponse: registerRespon,
		})
	}

}
