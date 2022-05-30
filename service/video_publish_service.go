package service

import (
	"SimpleTikTok/dao"
	"SimpleTikTok/model"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	VideoTitle = 128
)

func mp4ToJPEG(fileName string) string {
	if idx := strings.Index(fileName, "."); idx != -1 {
		fileName = fileName[:idx]
	}
	jpegFileNameArray := []string{fileName, ".jpeg"}
	output := strings.Join(jpegFileNameArray, "")
	return output
}

// use ffmpeg to create screen shot
func ExtractImage(data *multipart.FileHeader, fileName string) error {
	fileContent, _ := data.Open()
	var fileBytes []byte
	fileBytes, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return errors.New("unable to read video into memory")
	}
	command := "ffmpeg"
	vframes := "1"
	vf := "select=eq(n\\,1)"

	outputArray := []string{"public/photo/", mp4ToJPEG(fileName)}
	output := strings.Join(outputArray, "")

	cmd := exec.Command(command,
		"-i", "-", // to read from stdin
		"-vf", vf,
		"-vframes", vframes,
		output)

	cmd.Stdin = bytes.NewBuffer(fileBytes)

	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	return err
}
func PublishAction(data *multipart.FileHeader, title string, userId int64, c *gin.Context) error {
	if len(title) > VideoTitle {
		return errors.New("标题需小于128个字")
	}
	now := time.Now() // current local time
	sec := now.Unix()
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%d_%s", userId, sec, filename)
	saveFile := filepath.Join("./public/video/", finalName)
	dbVideoFile := filepath.Join("./video/", finalName)
	dbPhotoFile := filepath.Join("./photo/", mp4ToJPEG(finalName))
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		return err
	}

	if err := ExtractImage(data, finalName); err != nil {
		return err
	}
	video := &model.Videos{
		UserID:     userId,      // 发布者id
		VideoTitle: title,       // 视频标题
		VideoPath:  dbVideoFile, // 视频存放的路径
		CoverPath:  dbPhotoFile, // 视频封面图
		CreateTime: now,
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
