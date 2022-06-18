package service

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/dao"
	"SimpleTikTok/model"
)

func VideoWUserToVideo(VideoWithUser *model.VideoWithUser) commom.Video {
	author := commom.Userinfo{
		Id:            VideoWithUser.UserID,
		Name:          VideoWithUser.Username,
		FollowCount:   VideoWithUser.FollowCounts,
		FollowerCount: VideoWithUser.FansCounts,
		IsFollow:      false,
	}
	video := commom.Video{
		Id:     VideoWithUser.ID,
		Author: author,

		PlayUrl:       staticURL + VideoWithUser.VideoPath,
		CoverUrl:      staticURL + VideoWithUser.CoverPath,
		FavoriteCount: VideoWithUser.LikeCounts,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         VideoWithUser.VideoTitle,
	}
	return video

}

func GetGuestAccountFeed(lastTime int64) (int64, []commom.Video, error) {
	videos, err := dao.NewFeedDaoInstance().QueryVideoByTimeStamp(lastTime)
	if err != nil {
		return 0, nil, err
	}
	videoList := make([]commom.Video, len(videos))
	for i, video := range videos {
		videoList[i] = VideoWUserToVideo(video)
	}

	return videos[len(videos)-1].CreateTime.Unix(), videoList, err
}
func GetLoginAccountFeed(id int64, lastTime int64) (int64, []commom.Video, error) {
	timeStamp, videos, err := GetGuestAccountFeed(lastTime)
	if err != nil {
		return 0, nil, err
	}
	return timeStamp, videos, err
}
