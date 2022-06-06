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

type VideoWithUser struct {
	ID           int64
	UserID       int64     // 发布者id
	VideoTitle   string    // 视频描述
	VideoPath    string    // 视频存放的路径
	CoverPath    string    // 视频封面图
	LikeCounts   int64     // 喜欢/赞美的数量
	CreateTime   time.Time // 创建时间
	Username     string    // 用户名
	FansCounts   int64     // 我的粉丝数量
	FollowCounts int64     // 我关注的人总数

}
