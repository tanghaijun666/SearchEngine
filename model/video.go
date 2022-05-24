package model

import "time"

// 视频信息表
type Videos struct {
	ID           string    `gorm:"column:id;primary_key"`
	UserID       string    `gorm:"column:user_id;NOT NULL"`               // 发布者id
	VideoDesc    string    `gorm:"column:video_desc"`                     // 视频描述
	VideoPath    string    `gorm:"column:video_path;NOT NULL"`            // 视频存放的路径
	VideoSeconds float64   `gorm:"column:video_seconds"`                  // 视频秒数
	VideoWidth   int       `gorm:"column:video_width"`                    // 视频宽度
	VideoHeight  int       `gorm:"column:video_height"`                   // 视频高度
	CoverPath    string    `gorm:"column:cover_path"`                     // 视频封面图
	LikeCounts   int64     `gorm:"column:like_counts;default:0;NOT NULL"` // 喜欢/赞美的数量
	Status       int       `gorm:"column:status;NOT NULL"`                // 视频状态：1 发布成功 2禁止播放，管理员操作
	CreateTime   time.Time `gorm:"column:create_time;NOT NULL"`           // 创建时间
}
