package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(c *gin.Context) {
	var registerVar UserParam

	if err := c.BindQuery(&registerVar); err != nil {
		SendUserResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	log.Print(registerVar)

	if len(registerVar.Username) == 0 || len(registerVar.Password) == 0 {
		SendUserResponse(c, errno.ParamErr, -1, "")
		return
	}

	userID, token, err := rpc.CreateUser(context.Background(), &user.DouyinUserRegisterRequest{
		Username: registerVar.Username,
		Password: registerVar.Password,
	})

	if err != nil {
		SendUserResponse(c, errno.ConvertErr(err), -1, "")
		return
	}

	SendUserResponse(c, errno.Success, userID, token)
}
