package service

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/dao"
	"SimpleTikTok/pack"
	"sync"
)

type VideoService struct {
}

var videoService *VideoService
var serviceOnce sync.Once

func NewVideoServiceInstance() *VideoService {
	serviceOnce.Do(
		func() {
			videoService = &VideoService{}
		})
	return videoService
}
func (s *VideoService) FindVideoById(id int64) (*commom.Video, error) {
	videoModel, err := dao.NewVideoDaoInstance().QueryVideoById(id)
	if err != nil {
		return nil, err
	}

	if videoModel == nil {
		return nil, nil
	}

	userModel, err := dao.NewUserDaoInstance().QueryUserById(videoModel.AuthorId)
	if err != nil {
		return nil, err
	}

	user := pack.User(userModel)
	video := pack.Video(videoModel)

	video.Author = *user
	return video, nil
}
