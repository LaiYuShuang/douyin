package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func GetFeed(c *gin.Context) {

	var queryVar FeedParam

	if err := c.BindQuery(&queryVar); err != nil {
		SendFeedResponse(c, errno.ConvertErr(err), nil, 0)
	}

	if queryVar.LatestTime <= 0 {
		queryVar.LatestTime = 0
	}

	req := &video.DouyinFeedRequest{Token: &queryVar.Token, LatestTime: &queryVar.LatestTime}

	videoList, nextTime, err := rpc.GetFeed(context.Background(), req)

	if err != nil {
		SendFeedResponse(c, errno.ConvertErr(err), nil, 0)
		return
	}
	SendFeedResponse(c, errno.Success, videoList, nextTime)
}
