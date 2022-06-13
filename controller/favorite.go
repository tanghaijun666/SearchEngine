package controller

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var favoriteService = service.NewFavoriteServiceInstance()

type FavoriteActionResponse struct {
	commom.Response
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	c.JSON(http.StatusOK, FavoriteActionFunc(
		c.Query("video_id"),
		c.Query("token"),
		c.Query("action_type"),
	))
}

func FavoriteActionFunc(videoId, token, actionType string) FavoriteActionResponse {
	vid, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		return ErrorFavoriteResponse(err)
	}

	if actionType == "1" {
		err = favoriteService.Add(vid, token)
		if err != nil {
			return ErrorFavoriteResponse(err)
		}
		return FavoriteActionResponse{
			Response: commom.Response{
				StatusCode: 0,
				StatusMsg:  "谢谢你的喜欢! ",
			},
		}
	} else if actionType == "2" {
		err := favoriteService.Withdraw(vid, token)
		if err != nil {
			return ErrorFavoriteResponse(err)
		}
		return FavoriteActionResponse{
			Response: commom.Response{
				StatusCode: 0,
				StatusMsg:  "下次请收藏，已经取消收藏! ",
			},
		}
	} else {
		return FavoriteActionResponse{
			Response: commom.Response{
				StatusCode: 1,
				StatusMsg:  "服务出错了!",
			},
		}
	}
}

func ErrorFavoriteResponse(err error) FavoriteActionResponse {
	return FavoriteActionResponse{
		Response: commom.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		},
	}
}
