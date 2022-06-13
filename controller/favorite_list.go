package controller

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/pack"
	"errors"

	"github.com/gin-gonic/gin"
	"net/http"
)

type FavoriteListResponse struct {
	commom.Response
	VideoList []commom.Video `json:"video_list"`
}

// FavoriteList 收藏（喜欢）列表
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, FavoriteListFunc(
		c.Query("token"),
	))
}

func FavoriteListFunc(token string) FavoriteListResponse {
	// 使用token进行鉴权
	if token == "" {
		return ErrorFavoriteListResponse(errors.New("空令牌或者userId是空的"))
	}
	videos, err := favoriteService.FavoriteList(token)
	if err != nil {
		ErrorFavoriteListResponse(err)
	}
	return FavoriteListResponse{
		Response: commom.Response{
			StatusCode: 0,
			StatusMsg:  "加载收藏（点赞列表）成功!",
		},
		VideoList: pack.VideoPtrs(videos),
	}
}

func ErrorFavoriteListResponse(err error) FavoriteListResponse {
	return FavoriteListResponse{
		Response: commom.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		VideoList: nil,
	}
}
