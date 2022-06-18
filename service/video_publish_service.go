package service

import (
	"SimpleTikTok/commom"
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
	staticURL  = "http://localhost:8080/static/"
)

func ModeltoCommomStruct(mVideo *model.Videos, author commom.Userinfo) commom.Video {
	video := commom.Video{
		Id:     mVideo.ID,
		Author: author,

		PlayUrl:       staticURL + mVideo.VideoPath,
		CoverUrl:      staticURL + mVideo.CoverPath,
		FavoriteCount: mVideo.LikeCounts,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         mVideo.VideoTitle,
	}
	return video
}
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func getExt(fileName string) string {
	if idx := strings.LastIndex(fileName, "."); idx != -1 {
		fileName = fileName[idx+1:]
	}
	s := []string{"mp4", "3gp", "mov"}
	if contains(s, fileName) {
		return fileName
	}
	return "False"

}
func videoToJPEG(fileName string) string {
	if idx := strings.LastIndex(fileName, "."); idx != -1 {
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

	outputArray := []string{"public/photo/", videoToJPEG(fileName)}
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
	ext := getExt(filename)
	if ext == "false" {
		return errors.New("Invalid video type, please upload mp4 videos")
	}
	finalName := fmt.Sprintf("%d_%d.%s", userId, sec, ext)
	saveFile := filepath.Join("./public/video/", finalName)
	dbVideoFile := filepath.Join("./video/", finalName)
	dbPhotoFile := filepath.Join("./photo/", videoToJPEG(finalName))
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

func FindPublishList(id int64) ([]commom.Video, error) {
	videos, err := dao.NewVideoDaoInstance().QueryVideoByUserId(id)
	if err != nil {
		return nil, err
	}
	user, err := querybyId(id)
	if err != nil {
		return nil, err
	}
	videoList := make([]commom.Video, len(videos))
	for i, video := range videos {
		videoList[i] = ModeltoCommomStruct(video, *user)
	}
	return videoList, err
}
