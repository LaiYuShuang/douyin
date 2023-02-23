package service

import (
	"context"
	"douyin/cmd/video/dal/minicache"
	"douyin/cmd/video/pack"
	"douyin/kitex_gen/video"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/middleware"
	"fmt"
	"strconv"
	"time"
)

type SaveVideoService struct {
	ctx context.Context
}

// NewSaveVideoService new SaveVideoService
func NewSaveVideoService(ctx context.Context) *SaveVideoService {
	return &SaveVideoService{ctx: ctx}
}

// SaveVideo create video info.
func (s *SaveVideoService) SaveVideo(req *video.DouyinPublishActionRequest) error {

	//check video's token and get the video's id
	_, claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		return errno.AuthorizationFailedErr
	}
	userId := claims.UserID
	videoName := fmt.Sprintf("%s%s", strconv.FormatInt(userId, 16), strconv.FormatInt(time.Now().UnixNano()/1e6, 16))

	//VideoResourceAddr should be changed to your ServerHost
	playUrl := fmt.Sprintf("%s/%s/%s.mp4", constants.VideoResourceAddr, constants.VideoUrlPath, videoName)
	coverUrl := fmt.Sprintf("%s/%s/%s.png", constants.VideoResourceAddr, constants.VideoCoverUrlPath, videoName)

	//save video
	videoSavePath := fmt.Sprintf("%s/%s.mp4", constants.VideoSavePath, videoName)
	err = pack.SaveVideoToResource(videoSavePath, req.Data)
	if err != nil {
		return err
	}

	//save corver
	coverSavePath := fmt.Sprintf("%s/%s.png", constants.VideoCoverSavePath, videoName)
	err = pack.SaveCoverToResource(coverSavePath, videoSavePath)
	if err != nil {
		return err
	}

	//video write into database
	vio := pack.Video(userId, playUrl, coverUrl, req.Title)
	err = minicache.SaveVideo(s.ctx, vio)
	if err != nil {
		return err
	}
	return nil
}
