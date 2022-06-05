package service

import (
	"SimpleTikTok/dao"
	"SimpleTikTok/model"
	"errors"
)

type UserLoginResponse struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserLogininfo struct {
	username string
	password string
	data     *UserLoginResponse
}

const (
	MaxUsernmaeLength = 20
	MaxPasswordLength = 64
	MinPasswordLength = 8
)

//查询该用户是否已经存在,并且返回token和id
func QueryUserLogin(username string, password string) (*UserLoginResponse, error) {
	useinfo := UserLogininfo{username: username, password: password}
	if err := useinfo.checkLen(); err != nil {
		return nil, err
	}
	//获取数据
	if err := useinfo.getData(); err != nil {
		return nil, err
	}
	return useinfo.data, nil
}

func PostUserLogin(username, password string) (*UserLoginResponse, error) {
	useinfo := UserLogininfo{username: username, password: password}
	//校验参数长度
	if err := useinfo.checkLen(); err != nil {
		return nil, err
	}
	//更新数据到数据库
	if err := useinfo.updateData(); err != nil {
		return nil, err
	}
	return useinfo.data, nil

}

//更新用户的信息
func (u *UserLogininfo) updateData() error {
	userLoginDao := dao.NewUserLoginDao()
	user := model.Users{Username: u.username, Password: u.password}
	if userLoginDao.IsExisterUserbyname(u.username) {
		return errors.New("该用户已存在")
	}
	if err := userLoginDao.AddUser(&user); err != nil {
		return err
	}
	//颁发token
	token, err := dao.GenerateToken(user.ID)
	if err != nil {
		return err
	}
	u.data = &UserLoginResponse{UserId: int64(user.ID), Token: token}
	return nil
}

func (u *UserLogininfo) checkLen() error {
	if u.username == "" {
		return errors.New("用户名为空")
	}
	if u.password == "" {
		return errors.New("密码为空")
	}
	if len(u.username) > MaxUsernmaeLength {
		return errors.New("用户名过长")
	}
	if len(u.password) > MaxPasswordLength {
		return errors.New("密码过长")
	}
	if len(u.password) < MinPasswordLength {
		return errors.New("密码过短")
	}
	return nil
}

//得到需要的数据 UserId和token
func (u *UserLogininfo) getData() error {
	userLoginDao := dao.NewUserLoginDao()
	user := model.Users{Username: u.username, Password: u.password}
	if err := userLoginDao.CheckUser(u.username, u.password, &user); err != nil {
		return err
	}

	//颁发token
	token, err := dao.GenerateToken(user.ID)
	if err != nil {
		return err
	}

	u.data = &UserLoginResponse{UserId: int64(user.ID), Token: token}
	return nil

}
