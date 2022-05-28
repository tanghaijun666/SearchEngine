package model

import "time"

// 视频信息表
type Videos struct {
	ID         int64     `gorm:"column:id;primary_key"`
	UserID     int64     `gorm:"column:user_id;NOT NULL"`               // 发布者id
	VideoTitle string    `gorm:"column:video_title"`                    // 视频描述
	VideoPath  string    `gorm:"column:video_path;NOT NULL"`            // 视频存放的路径
	CoverPath  string    `gorm:"column:cover_path"`                     // 视频封面图
	LikeCounts int64     `gorm:"column:like_counts;default:0;NOT NULL"` // 喜欢/赞美的数量
	CreateTime time.Time `gorm:"column:create_time;NOT NULL"`           // 创建时间
}
