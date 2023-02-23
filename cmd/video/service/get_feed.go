package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/dal/minicache"
	"douyin/kitex_gen/video"
	"douyin/pkg/constants"
	"fmt"
)

type FeedService struct {
	ctx context.Context
}

// NewFeedService new FeedService
func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) GetFeed(req *video.DouyinFeedRequest) ([]*video.Video, int64, error) {
	videos, err := minicache.GetFeedList(s.ctx, *req.LatestTime, constants.VideoLimitNum)
	if err != nil {
		return nil, 0, err
	}
	vios := make([]*video.Video, 0)
	for _, vv := range videos {
		//get author

		fmt.Println(vv.UserId)
		user, err := db.QueryUser(s.ctx, vv.UserId)
		if err != nil {
			return nil, 0, err
		}
		u := video.User{
			Id:   user.Id,
			Name: user.Name,
		}

		//get IsFavorite by rpc, to be done?
		v := video.Video{
			Id:            vv.Id,
			Author:        &u,
			PlayUrl:       vv.PlayUrl,
			CoverUrl:      vv.CoverUrl,
			FavoriteCount: vv.FavoriteCount,
			CommentCount:  vv.CommentCount,
			Title:         vv.Title,
			IsFavorite:    false,
		}
		vios = append(vios, &v)
	}
	time := videos[len(videos)-1].CreateTime
	return vios, time, nil
}
