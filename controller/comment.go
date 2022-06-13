package controller

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/dao"
	"SimpleTikTok/model"
	"SimpleTikTok/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	commom.Response
	CommentList []commom.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	commom.Response
	Comment commom.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	content := c.Query("comment_text")
	commentId, err := strconv.Atoi(c.Query("comment_id"))
	if err != nil {
		c.JSON(http.StatusOK, commom.Response{StatusCode: 1, StatusMsg: "commentId 出错"})
	}
	videoId := c.Query("video_id")
	db := commom.GetDB()
	if actionType == "1" {
		var comment model.Comments
		var userId int64
		if token != "" {
			userId, err = dao.JwtAuth(token)
			if err != nil {
				c.JSON(http.StatusOK, commom.Response{StatusCode: 1, StatusMsg: "token 出错"})
			}
		} else {
			userId = 1
		}
		comment.UserId = int(userId)
		//comment.ID = commentId
		comment.Content = content
		comment.VideoId = videoId
		comment.CreateDate = time.Now()

		db.Table("comments").Select("id").Find(&commentId)
		comment.ID = commentId + 1
		db.Table("comments").Create(&comment)

		var commentCommon commom.Comment
		commentCommon.Id = int64(comment.ID)
		var userinfo model.Users
		db.Table("users").Where("id=?", comment.UserId).Find(&userinfo)
		commentCommon.User.Id = userinfo.ID
		commentCommon.User.Name = userinfo.Username
		commentCommon.User.FollowerCount = userinfo.FansCounts
		commentCommon.User.FollowCount = userinfo.FollowCounts
		commentCommon.User.IsFollow = false
		commentCommon.CreateDate = comment.CreateDate.Format("2006.01.02 15:04:05")
		commentCommon.Content = comment.Content

		c.JSON(http.StatusOK, CommentActionResponse{
			Response: commom.Response{StatusCode: 0, StatusMsg: "成功"},
			Comment:  commentCommon})
	} else if actionType == "2" {
		var comment model.Comments
		var userId int64
		if token != "" {
			userId, err = dao.JwtAuth(token)
			if err != nil {
				c.JSON(http.StatusOK, commom.Response{StatusCode: 1, StatusMsg: "token 出错"})
			}
		} else {
			userId = 1
		}
		comment.UserId = int(userId)
		//comment.ID = commentId
		comment.Content = content
		comment.VideoId = videoId
		comment.CreateDate = time.Now()
		//
		//db.Table("comments").Select("id").Find(&commentId)
		//comment.ID = commentId + 1
		//db.Table("comments").Create(&comment)

		var commentCommon commom.Comment
		commentCommon.Id = int64(comment.ID)
		var userinfo model.Users
		db.Table("users").Where("id=?", comment.UserId).Find(&userinfo)
		commentCommon.User.Id = userinfo.ID
		commentCommon.User.Name = userinfo.Username
		commentCommon.User.FollowerCount = userinfo.FansCounts
		commentCommon.User.FollowCount = userinfo.FollowCounts
		commentCommon.User.IsFollow = false
		commentCommon.CreateDate = comment.CreateDate.Format("2006.01.02 15:04:05")
		commentCommon.Content = comment.Content

		//db.Table("comments").Select("id").Find(&commentId)
		comment.ID = commentId
		db.Table("comments").Delete(&comment)
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: commom.Response{StatusCode: 0, StatusMsg: "成功"},
			Comment:  commentCommon})
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		fmt.Printf("参数videoId必须为整数")
	}
	comments, err := service.CommentsList(int64(videoId))

	if comments == nil {

	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    commom.Response{StatusCode: 0},
		CommentList: comments,
	})
}
