package dao

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/model"
	"sync"
)

type CommentDao struct {
}

var (
	commentDao  *CommentDao
	commentOnce sync.Once
)

func NewCommentDao() *CommentDao {
	commentOnce.Do(func() {
		commentDao = new(CommentDao)
	})
	return commentDao
}

func (*CommentDao) QueryCommentsById(videoId int64) []model.Comments {
	var comments []model.Comments
	//从连接池里拿数据
	db := commom.GetDB()
	db.Table("comments").Find(&comments).Where("video_id = ?", videoId)
	if comments == nil {
		return nil
	}
	return comments
}
