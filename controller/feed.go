package controller

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/dao"
	"SimpleTikTok/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	commom.Response
	VideoList []commom.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// if lastTime value doesn't exist use currentTime
	// lastTime : > the returning video's newest creating time
	var lastTime int64 = time.Now().Unix()
	var err error
	lastTimeInput := c.Query("lastest_Time")
	if lastTimeInput != "" {
		lastTime, err = strconv.ParseInt(lastTimeInput, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK,
				commom.Response{StatusCode: 1, StatusMsg: "Wrong LastTime value"})
			return
		}
	}
	var videoList []commom.Video
	var newTimeStamp int64
	token := c.Query("token")
	if token == "" {
		newTimeStamp, videoList, err = service.GetGuestAccountFeed(lastTime)
		if err != nil {
			c.JSON(http.StatusOK,
				commom.Response{StatusCode: 1, StatusMsg: fmt.Sprint(err)})
			return

		}

	} else if _, err := dao.JwtAuth(token); err != nil {
		newTimeStamp, videoList, err = service.GetGuestAccountFeed(lastTime)

	} else {
		id, _ := dao.JwtAuth(token)
		newTimeStamp, videoList, err = service.GetLoginAccountFeed(id, lastTime)

	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  commom.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  newTimeStamp,
	})
}
