package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func GetPublishList(c *gin.Context) {
	var queryVar PublishListParam
	if err := c.BindQuery(&queryVar); err != nil {
		SendVideoResponse(c, errno.ConvertErr(err), nil)
	}

	req := &video.DouyinPublishListRequest{Token: queryVar.Token, UserId: queryVar.UserId}

	videoList, err := rpc.GetPublishList(context.Background(), req)

	if err != nil {
		SendVideoResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendVideoResponse(c, errno.Success, videoList)
}
