package service

import (
	"SimpleTikTok/dao"
	"SimpleTikTok/model"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	VideoTitle = 128
)

func PublishAction(data *multipart.FileHeader, title string, userId int64, c *gin.Context) error {
	if len(title) > VideoTitle {
		return errors.New("标题需小于128个字")
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		return err
	}
	video := &model.Videos{
		UserID:     userId,    // 发布者id
		VideoTitle: title,     // 视频标题
		VideoPath:  finalName, // 视频存放的路径
		// CoverPath:               ,      // 视频封面图
		CreateTime: time.Now(),
	}

	// update the database
	err := dao.NewVideoDaoInstance().CreateVideo(video)
	if err != nil {
		os.Remove(saveFile)
		return err
	}
	return nil
}

func PublishList(id int64) ([]*model.Videos, error) {
	videos, err := dao.NewVideoDaoInstance().QueryVideoByUserId(id)
	return videos, err
}
