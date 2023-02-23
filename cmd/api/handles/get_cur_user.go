package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func QueryCurUser(c *gin.Context) {
	var queryVar QueryUserParam
	if err := c.BindQuery(&queryVar); err != nil {
		SendQueryUserResponse(c, errno.ConvertErr(err), &UserInfo{
			ID:   -1,
			Name: "",
		})
		return
	}
	userInfo, err := rpc.QueryUser(context.Background(), &user.DouyinUserRequest{
		UserId: queryVar.UserID,
		Token:  queryVar.Token,
	})
	if err != nil {
		SendQueryUserResponse(c, errno.ConvertErr(err), &UserInfo{
			ID:   -1,
			Name: "",
		})
		return
	}
	SendQueryUserResponse(c, errno.Success, &UserInfo{
		ID:   userInfo.ID,
		Name: userInfo.Name,
	})
}
