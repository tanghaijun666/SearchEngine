package dao

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/model"
	"sync"
)

type CommentDaoByTHJ struct {
}

var (
	commentDaoByTHJ  *CommentDaoByTHJ
	commentOnceByTHJ sync.Once
)

func NewCommentDaoByTHJ() *CommentDaoByTHJ {
	commentOnceByTHJ.Do(func() {
		commentDaoByTHJ = new(CommentDaoByTHJ)
	})
	return commentDaoByTHJ
}

func (*CommentDaoByTHJ) QueryCommentsById(videoId int64) []model.Comments {
	var comments []model.Comments
	//从连接池里拿数据
	db := commom.GetDB()
	db.Table("comments").Find(&comments).Where("video_id = ?", videoId)
	if comments == nil {
		return nil
	}
	return comments
}
