package commom

//这里存的都是要返回的json

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64    `json:"id"`
	Author        Userinfo `json:"author"`
	PlayUrl       string   `json:"play_url" json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}

type Comment struct {
	Id         int64    `json:"id,omitempty"`
	User       Userinfo `json:"user"`
	Content    string   `json:"content,omitempty"`
	CreateDate string   `json:"create_date,omitempty"`
}

type Userinfo struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
