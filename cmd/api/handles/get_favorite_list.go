package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func GetFavoriteList(c *gin.Context) {
	var queryVar FavoriteListParam
	if err := c.BindQuery(&queryVar); err != nil {
		SendFavoriteResponse(c, errno.ConvertErr(err), nil)
	}

	req := &favorite.DouyinFavoriteListRequest{Token: queryVar.Token, UserId: queryVar.UserId}

	videoList, err := rpc.GetFavoriteList(context.Background(), req)

	if err != nil {
		SendFavoriteResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendFavoriteResponse(c, errno.Success, videoList)
}
