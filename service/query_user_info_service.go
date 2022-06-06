package service

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/dao"
	"SimpleTikTok/model"
)

func querybyId(id int64) (*commom.Userinfo, error) {
	var user model.Users
	userLoginDao := dao.NewUserLoginDao()
	if err := userLoginDao.QueryUserbyId(id, &user); err != nil {
		return nil, err
	}
	comUser := &commom.Userinfo{
		Id:            user.ID,
		Name:          user.Username,
		FollowCount:   user.FollowCounts,
		FollowerCount: user.FansCounts,
		IsFollow:      false,
	}
	return comUser, nil
}
