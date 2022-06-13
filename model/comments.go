package model

import "time"

// 视频评论表
//comment数据表构造，因为之前的版本id都为string实际应该是int,现在修改过来了
type Comments struct {
	ID              int64     `gorm:"column:id;primary_key"`
	FatherCommentID int64     `gorm:"column:father_comment_id"`     // 父评论id
	ToUserID        int64     `gorm:"column:to_user_id;NOT NULL"`   // 被评论的用户id
	VideoID         int64     `gorm:"column:video_id;NOT NULL"`     // 视频id
	FromUserID      int64     `gorm:"column:from_user_id;NOT NULL"` // 留言者，评论的用户id
	Comment         int64     `gorm:"column:comment;NOT NULL"`      // 评论内容
	CreateTime      time.Time `gorm:"column:create_time;NOT NULL"`
}
