package main

import (
	"context"
	"douyin/cmd/video/pack"
	"douyin/cmd/video/service"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	resp = new(video.DouyinPublishActionResponse)

	//check param
	if len(req.Data) == 0 || len(req.Token) == 0 || len(req.Title) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	//save video and create video
	err = service.NewSaveVideoService(ctx).SaveVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	//success
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.DouyinPublishListResponse)

	//check param
	if req.UserId == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		resp.VideoList = nil
		return resp, nil
	}

	//get publish list
	videos, err := service.NewPublishListService(ctx).GetPublishList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.VideoList = nil
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	resp = new(video.DouyinFeedResponse)

	videos, latestTime, err := service.NewFeedService(ctx).GetFeed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.VideoList = nil
		resp.NextTime = nil
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	resp.NextTime = &latestTime
	return resp, nil
}
