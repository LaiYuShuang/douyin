package service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/favorite"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) GetFavoriteList(req *favorite.DouyinFavoriteListRequest) ([]*favorite.Video, error) {
	// get video'id list
	videos := make([]*favorite.Video, 0)
	vids, err := db.GetVideoIdList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	for _, vid := range vids {
		//get video
		video, err := db.GetVideoById(s.ctx, vid)
		if err != nil {
			return nil, err
		}

		//get author
		user, err := db.QueryUser(s.ctx, video.UserId)
		if err != nil {
			return nil, err
		}
		u := favorite.User{
			Id:   user.Id,
			Name: user.Name,
		}

		v := favorite.Video{
			Id:            video.Id,
			Author:        &u,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			IsFavorite:    true,
		}
		videos = append(videos, &v)
	}
	return videos, nil
}
