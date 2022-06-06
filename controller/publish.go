package controller

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/dao"
	"SimpleTikTok/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	commom.Response
	VideoList []commom.Video `json:"video_list"`
}

// check if logged in
func CheckLogIn(token string, c *gin.Context) bool {
	if _, err := dao.JwtAuth(token); err != nil {
		c.JSON(http.StatusOK, commom.Response{StatusCode: 1, StatusMsg: "Not Logged In"})
		return false
	}
	return true
}

// Publish check token then save uploaded file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	// use token to check if logged In
	if islogin := CheckLogIn(token, c); !islogin {
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, commom.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	title := c.PostForm("title")
	id, _ := dao.JwtAuth(token)
	err = service.PublishAction(data, title, id, c)
	if err != nil {
		c.JSON(http.StatusOK, commom.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, commom.Response{
		StatusCode: 0,
		StatusMsg:  "Upload Succeed!",
	})

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	tokenId, err := dao.JwtAuth(token)
	if err != nil {
		c.JSON(http.StatusOK, commom.Response{StatusCode: 1, StatusMsg: "Not Logged In"})
		return

	}
	id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, commom.Response{
			StatusCode: 1,
			StatusMsg:  "Incorrect user id form",
		})
		return

	}
	if tokenId != int64(id) {
		c.JSON(http.StatusOK, commom.Response{
			StatusCode: 1,
			StatusMsg:  "Inconsistent User Id and Token",
		})
		return

	}
	videos, err := service.FindPublishList(int64(id))
	if err != nil {
		c.JSON(http.StatusOK, commom.Response{
			StatusCode: 1,
			StatusMsg:  "unable to find videos",
		})
		return
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: commom.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
