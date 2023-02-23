package service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
	"douyin/pkg/middleware"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

func (s *FavoriteActionService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) error {
	// update video
	err := db.SetVideo(s.ctx, req.VideoId, req.ActionType)
	if err != nil {
		return err
	}

	//check video's token
	_, claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		return errno.AuthorizationFailedErr
	}
	userId := claims.UserID

	//create favorite
	err = db.UpdateFavorite(s.ctx, userId, req.VideoId, req.ActionType)
	if err != nil {
		//update video
		err2 := db.SetVideo(s.ctx, req.VideoId, 2)
		if err2 != nil {
			return err2
		}
		return err
	}
	return nil

}
