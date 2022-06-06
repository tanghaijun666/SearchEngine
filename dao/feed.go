package dao

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/model"
	"sync"
	"time"
)

type FeedDao struct {
}

var feedDao *FeedDao
var feedOnce sync.Once

func NewFeedDaoInstance() *FeedDao {
	feedOnce.Do(
		func() {
			feedDao = new(FeedDao)
		})
	return feedDao
}

func (*FeedDao) QueryVideoByTimeStamp(lastTime int64) ([]*model.VideoWithUser, error) {
	var videoswithUsers []*model.VideoWithUser
	tm := time.Unix(lastTime, 0)
	db := commom.GetDB()
	err := db.Table("videos").Joins("left join users on videos.user_id = users.id").
		Select(`videos.id ,videos.user_id, videos.video_title, videos.video_path, videos.cover_path, videos.like_counts,videos.create_time, users.username, users.fans_counts, users.follow_counts`).
		Where("create_time < ?", tm).Order("create_time desc").Limit(30).Find(&videoswithUsers).Error
	if err != nil {
		return nil, err
	}
	return videoswithUsers, nil
}
