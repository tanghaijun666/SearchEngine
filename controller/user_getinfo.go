package controller

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/dao"
	"SimpleTikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//暂时直接使用userid去进行用户的查询

type UserInfoRespone struct {
	commom.Response
	*commom.Userinfo `json:"user"`
}

//封装代理查询操作

type ProxyUserinfo struct {
	c *gin.Context
	u *commom.Userinfo
}

func NewProxyUserinfo(c *gin.Context) *ProxyUserinfo {
	return &ProxyUserinfo{c: c}
}

func UserGetInfoHandle(c *gin.Context) {
	p := NewProxyUserinfo(c)
	rawId := c.Query("user_id")
	p.QueryUserInfoById(rawId)

	//暂时不使用token
	//token := c.Query("token")

	//根据userid 查询

}
func (q *ProxyUserinfo) QueryUserInfoById(rawId string) {
	userId, _ := strconv.ParseInt(rawId, 10, 64)
	var user model.Users
	userLoginDao := dao.NewUserLoginDao()
	if err := userLoginDao.QueryUserbyId(userId, &user); err != nil {
		q.ReturnFail(err.Error())
	}
	q.u = &commom.Userinfo{
		Id:            user.ID,
		Name:          user.Username,
		FollowCount:   user.FollowCounts,
		FollowerCount: user.FansCounts,
		IsFollow:      false,
	}
	q.ReturnOk()
}

func (q *ProxyUserinfo) ReturnOk() {
	q.c.JSON(http.StatusOK, UserInfoRespone{
		Response: commom.Response{StatusCode: 0},
		Userinfo: q.u,
	})
}
func (q *ProxyUserinfo) ReturnFail(msg string) {
	q.c.JSON(http.StatusOK, UserInfoRespone{
		Response: commom.Response{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}
