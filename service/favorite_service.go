package service

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/dao"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type FavoriteService struct {
}

var favoriteService *FavoriteService
var favoriteOnce sync.Once

func NewFavoriteServiceInstance() *FavoriteService {
	favoriteOnce.Do(
		func() {
			favoriteService = &FavoriteService{}
		})
	return favoriteService
}

func (s *FavoriteService) FindUserByToken(token string) (*commom.Userinfo, error) {
	userId, err := dao.JwtAuth(token)
	if err != nil {
		//panic(err)
		return nil, errors.New("该token解析失败")
	}
	user, err := querybyId(userId)
	if err != nil {
		//panic(err)
		return nil, errors.New("查询用户失败")
	}
	fmt.Println(user)
	fmt.Println(userId)
	return user, nil
}

func (s *FavoriteService) FindVideosByToken(token string) ([]*commom.Video, error) {
	// invalid token
	if token == "" {
		return nil, nil
	}
	videoIds, err := dao.NewFavoriteDaoInstance().QueryVideoIdByToken(token)
	if err != nil {
		return nil, err
	}
	var videos []*commom.Video
	for _, id := range videoIds {
		video, _ := NewVideoServiceInstance().FindVideoById(id)
		//video.IsFavorite = true
		videos = append(videos, video)
	}
	return videos, nil
}

func (s *FavoriteService) TotalComment() (int64, error) {
	count, err := dao.NewFavoriteDaoInstance().Total()
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (s *FavoriteService) LastId() (int64, error) {
	count, err := dao.NewFavoriteDaoInstance().MaxId()
	if err != nil {
		return count, err
	}
	return count, nil
}

func (s *FavoriteService) Add(videoId int64, token string) error {
	// 点赞
	favoriteIdSequence, _ := favoriteService.LastId()
	atomic.AddInt64(&favoriteIdSequence, 1)
	newFavorite := &dao.Favorite{
		Id:        favoriteIdSequence,
		UserToken: token,
		VideoId:   videoId,
		CreateAt:  time.Now(),
	}
	err := dao.NewFavoriteDaoInstance().Save(newFavorite)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavoriteService) Withdraw(videoId int64, token string) error {
	// 删除喜欢
	err := dao.NewFavoriteDaoInstance().Delete(videoId, token)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavoriteService) FavoriteList(token string) ([]*commom.Video, error) {
	return s.FindVideosByToken(token)
}
