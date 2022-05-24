package model

type Users struct {
	ID                string `gorm:"column:id;primary_key"`
	Username          string `gorm:"column:username;NOT NULL"`             // 用户名
	Password          string `gorm:"column:password;NOT NULL"`             // 密码
	FaceImage         string `gorm:"column:face_image"`                    // 我的头像，如果没有默认给一张
	Nickname          string `gorm:"column:nickname"`                      // 昵称
	FansCounts        int    `gorm:"column:fans_counts;default:0"`         // 我的粉丝数量
	FollowCounts      int    `gorm:"column:follow_counts;default:0"`       // 我关注的人总数
	ReceiveLikeCounts int    `gorm:"column:receive_like_counts;default:0"` // 我接受到的赞美/收藏 的数量
}

// 用户喜欢的/赞过的视频
type UsersLikeVideos struct {
	ID      string `gorm:"column:id;primary_key"`
	UserID  string `gorm:"column:user_id;NOT NULL"`  // 用户
	VideoID string `gorm:"column:video_id;NOT NULL"` // 视频
}

// 用户粉丝关联关系表
type UsersFans struct {
	ID     string `gorm:"column:id;primary_key"`
	UserID string `gorm:"column:user_id;NOT NULL"` // 用户
	FanID  string `gorm:"column:fan_id;NOT NULL"`  // 粉丝
}
