package dao

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/model"
	"errors"
	"sync"
)

type UserLoginDao struct {
}

//生成一个新的单例，实现LoginDao的
var (
	userLoginDao *UserLoginDao
	once         sync.Once
)

func NewUserLoginDao() *UserLoginDao {
	//保证只执行一次
	once.Do(func() {
		userLoginDao = new(UserLoginDao)
	})
	return userLoginDao
}

func (*UserLoginDao) AddUser(user *model.Users) error {
	if user == nil {
		return errors.New("user 指针为空")
	}
	db := commom.GetDB()
	return db.Create(user).Error
}

//查询username和password 是否匹配
func (*UserLoginDao) CheckUser(username string, password string, user *model.Users) error {
	if user == nil {
		return errors.New("结构体为空")
	}
	db := commom.GetDB()
	db.Where("username=? and password=?", username, password).Take(user)
	if user.ID == 0 {
		return errors.New("用户不存在,或者账号密码出错")
	}
	return nil
}

//查询是否已经存在用户
func (*UserLoginDao) IsExisterUserbyname(username string) bool {
	var user model.Users
	//从连接池里拿数据
	db := commom.GetDB()
	db.Where("username=?", username).Take(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

//查询用户信息
func (*UserLoginDao) QueryUserbyId(userId int64, user *model.Users) error {
	if user == nil {
		return errors.New("QueryUserbyId user指针为空")
	}
	db := commom.GetDB()
	db.Where("id=?", userId).First(user)
	if user.ID == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}
