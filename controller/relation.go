package controller

import (
	"SimpleTikTok/commom"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	commom.Response
	UserList []commom.Userinfo `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, commom.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, commom.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: commom.Response{
			StatusCode: 0,
		},
		UserList: []commom.Userinfo{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: commom.Response{
			StatusCode: 0,
		},
		UserList: []commom.Userinfo{DemoUser},
	})
}
