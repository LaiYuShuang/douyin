package pack

import (
	"bytes"
	"douyin/cmd/video/dal/db"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io/ioutil"
	"os"
	"time"
)

//Video return a db.video
func Video(usrId int64, playUrl, coverUrl, title string) db.Video {
	return db.Video{
		UserId:        usrId,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		Title:         title,
		FavoriteCount: 0,
		CommentCount:  0,
		CreateTime:    time.Now().Unix(),
	}
}

// SaveVideoToResource save video into api/resource/videos
func SaveVideoToResource(videoPath string, video []byte) error {
	//save video
	err := ioutil.WriteFile(videoPath, video, 0666)
	if err != nil {
		return err
	}
	return nil
}

// SaveCoverToResource save video into api/resource/covers
func SaveCoverToResource(coverPath string, videoPath string) error {

	//make cover from video
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return err
	}

	//save cover
	img, err := imaging.Decode(buf)
	if err != nil {
		return err
	}
	err = imaging.Save(img, coverPath)
	if err != nil {
		return err
	}

	return nil
}
