package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/dal/minicache"
	"douyin/kitex_gen/video"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// GetPublishList get video list.
func (s *PublishListService) GetPublishList(req *video.DouyinPublishListRequest) ([]*video.Video, error) {
	videos, err := minicache.GetVideoList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	vios := make([]*video.Video, 0)
	for _, vv := range videos {
		//get author
		user, err := db.QueryUser(s.ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		u := video.User{
			Id:   user.Id,
			Name: user.Name,
		}

		//get IsFavorite
		isfavor, err := db.IsFavorite(s.ctx, req.UserId, vv.Id)

		v := video.Video{
			Id:            vv.Id,
			Author:        &u,
			PlayUrl:       vv.PlayUrl,
			CoverUrl:      vv.CoverUrl,
			FavoriteCount: vv.FavoriteCount,
			CommentCount:  vv.CommentCount,
			Title:         vv.Title,
			IsFavorite:    isfavor,
		}
		vios = append(vios, &v)

	}
	return vios, nil
}
