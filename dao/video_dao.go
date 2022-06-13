package dao

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/model"
	"errors"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

type VideoDao struct {
}

var videoDao *VideoDao
var publishOnce sync.Once

// NewVideoDaoInstance Singleton
func NewVideoDaoInstance() *VideoDao {
	publishOnce.Do(
		func() {
			videoDao = new(VideoDao)
		})

	return videoDao
}

func (*VideoDao) QueryVideoByUserId(userId int64) ([]*model.Videos, error) {
	var videos []*model.Videos
	db := commom.GetDB()
	err := db.Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		return nil, errors.New("无上传视屏")
	}
	return videos, nil
}

func (*VideoDao) CreateVideo(video *model.Videos) error {
	db := commom.GetDB()
	if err := db.Create(video).Error; err != nil {
		return errors.New("视频添加失败")
	}
	return nil
}

//akun后加的
// 如果没有找到用户，QueryVideoById 将返回 nil
func (*VideoDao) QueryVideoById(id int64) (*model.Video, error) {
	var video model.Video
	db := commom.GetDB()
	err := db.Where("id = ?", id).First(&video).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Fatal("find video by id err:" + err.Error())
		return nil, err
	}
	return &video, nil
}

// 如果未找到用户，QueryVideoBeforeTime 将返回空数组
func (*VideoDao) QueryVideoBeforeTime(time time.Time, limit int) ([]*model.Video, error) {
	var videos []*model.Video
	db := commom.GetDB()
	err := db.Where("create_at < ?", time).Order("create_at DESC").Limit(limit).Find(&videos).Error

	if err != nil {
		log.Fatal("batch find video before time err:" + err.Error())
		return nil, err
	}
	return videos, nil
}

func (*VideoDao) QueryVideoByAuthorId(authorId int64) ([]*model.Video, error) {
	var videos []*model.Video
	db := commom.GetDB()
	err := db.Where("author_id = ?", authorId).Find(&videos).Error
	if err != nil {
		log.Fatal("batch find video by author_id err:" + err.Error())
		return nil, err
	}
	return videos, nil
}
