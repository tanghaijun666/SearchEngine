package controller

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/model"
	"SimpleTikTok/service"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	staticURL = "http://localhost:8080/static/"
)

type VideoListResponse struct {
	commom.Response
	VideoList []commom.Video `json:"video_list"`
}

// check if logged in
func CheckLogIn(token string, c *gin.Context) bool {
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, commom.Response{StatusCode: 1, StatusMsg: token})

		return false
	}
	return true
}
func ModeltoCommomStruct(mVideo *model.Videos, author commom.Userinfo) commom.Video {
	video := commom.Video{
		Id:     mVideo.ID,
		Author: author,

		PlayUrl:  filepath.Join(staticURL, mVideo.VideoPath),
		CoverUrl: filepath.Join(staticURL, mVideo.CoverPath),
		// FavoriteCount int64    `json:"favorite_count,omitempty"`
		// CommentCount  int64    `json:"comment_count,omitempty"`
		// IsFavorite    bool     `json:"is_favorite,omitempty"`
		Title: mVideo.VideoTitle,
	}
	return video
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
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
	id := usersLoginInfo[token].Id
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
	login := CheckLogIn(token, c)
	if !login {
		return
	}
	userInMap := usersLoginInfo[token]
	id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil || int64(id) != userInMap.Id {
		c.JSON(http.StatusOK, commom.Response{
			StatusCode: 1,
			StatusMsg:  "unable to find user id",
		})
		return

	}
	videos, err := service.PublishList(userInMap.Id)
	if err != nil {
		c.JSON(http.StatusOK, commom.Response{
			StatusCode: 1,
			StatusMsg:  "unable to find user id",
		})
		return
	}
	videoList := make([]commom.Video, len(videos))
	for i, video := range videos {
		videoList[i] = ModeltoCommomStruct(video, userInMap)
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: commom.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
