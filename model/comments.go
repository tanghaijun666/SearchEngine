package model

// 视频评论表
//comment数据表暂未构造，因为之前的版本id都为string实际应该是int
//type Comments struct {
//	ID             int    `gorm:"column:id;primary_key"`
//	FatherCommentID string    `gorm:"column:father_comment_id"`     // 父评论id
//	ToUserID        string    `gorm:"column:to_user_id;NOT NULL"`   // 被评论的用户id
//	VideoID         string    `gorm:"column:video_id;NOT NULL"`     // 视频id
//	FromUserID      string    `gorm:"column:from_user_id;NOT NULL"` // 留言者，评论的用户id
//	Comment         string    `gorm:"column:comment;NOT NULL"`      // 评论内容
//	CreateTime      time.Time `gorm:"column:create_time;NOT NULL"`
//}
