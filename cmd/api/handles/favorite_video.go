package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func FavoriteVideo(c *gin.Context) {
	var queryVar FavoriteActionParam
	if err := c.BindQuery(&queryVar); err != nil {
		SendFavoriteResponse(c, errno.ConvertErr(err), nil)
		return
	}
	req := &favorite.DouyinFavoriteActionRequest{VideoId: queryVar.VideoId, Token: queryVar.Token, ActionType: queryVar.ActionType}
	err := rpc.FavoriteVideo(context.Background(), req)
	if err != nil {
		SendFavoriteResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendFavoriteResponse(c, errno.Success, nil)
}
