package main

import (
	"context"
	"douyin/cmd/favorite/pack"
	"douyin/cmd/favorite/service"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.DouyinFavoriteActionResponse)
	//check param
	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp.BaseRasp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp.BaseRasp = pack.BuildBaseResp(err)
		return resp, err
	}

	resp.BaseRasp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.DouyinFavoriteListResponse)
	//check param
	if len(req.Token) == 0 || req.UserId == 0 {
		resp.BaseRape = pack.BuildBaseResp(errno.ParamErr)
		resp.VideoList = nil
		return resp, nil
	}

	videos, err := service.NewFavoriteListService(ctx).GetFavoriteList(req)
	if err != nil {
		resp.BaseRape = pack.BuildBaseResp(err)
		resp.VideoList = nil
		return resp, err
	}

	resp.BaseRape = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}
