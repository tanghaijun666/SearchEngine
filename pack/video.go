package pack

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/model"
)

func AuthorIds(videoModels []*model.Video) []int64 {
	if videoModels != nil {
		var ids = make([]int64, 0, len(videoModels))
		for _, videoModel := range videoModels {
			ids = append(ids, videoModel.AuthorId)
		}
		return ids
	}
	return []int64{}
}

func Video(videoModel *model.Video) *commom.Video {
	if videoModel != nil {
		return &commom.Video{
			Id:            videoModel.Id,
			Author:        commom.Userinfo{},
			PlayUrl:       videoModel.PlayUrl,
			CoverUrl:      videoModel.CoverUrl,
			Title:         videoModel.Title,
			FavoriteCount: videoModel.FavoriteCount,
			CommentCount:  videoModel.FavoriteCount,
		}
	}
	return nil
}

func Videos(videoModels []*model.Video) []*commom.Video {
	if videoModels != nil {
		var videos = make([]*commom.Video, 0, len(videoModels))
		for _, model := range videoModels {
			videos = append(videos, Video(model))
		}
		return videos
	}
	return nil
}

func VideoPtrs(videoPtrs []*commom.Video) []commom.Video {
	if videoPtrs != nil {
		var videos = make([]commom.Video, len(videoPtrs))
		for i, ptr := range videoPtrs {
			videos[i] = *ptr
		}
		return videos
	}
	return []commom.Video{}
}
