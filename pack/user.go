package pack

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/dao"
)

// 用户如果参数为 nil，则返回 nil
func User(userModel *dao.User) *commom.Userinfo {
	if userModel != nil {
		return &commom.Userinfo{
			Id:            userModel.Id,
			Name:          userModel.Name,
			FollowCount:   userModel.FollowCount,
			FollowerCount: userModel.FollowerCount,
		}
	}
	return nil
}

func Users(userModels []*dao.User) []*commom.Userinfo {
	if userModels != nil {
		var users = make([]*commom.Userinfo, 0, len(userModels))
		for _, model := range userModels {
			users = append(users, User(model))
		}
		return users
	}
	return nil
}

// MUser 如果参数为 nil 则返回空映射
func MUser(userModels map[int64]dao.User) map[int64]*commom.Userinfo {
	if userModels != nil {
		var users = make(map[int64]*commom.Userinfo, len(userModels))
		for id, userModel := range userModels {
			users[id] = User(&userModel)
		}
		return users
	}
	return nil
}

func MUserByName(userModels map[string]dao.User) map[string]commom.Userinfo {
	if userModels != nil {
		var users = make(map[string]commom.Userinfo, len(userModels))
		for name, userModel := range userModels {
			users[name] = *User(&userModel)
		}
		return users
	}
	return nil
}
