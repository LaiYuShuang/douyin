package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func PublishVideo(c *gin.Context) {
	var queryVar PublishActionParam

	queryVar.Token = c.PostForm("token")
	queryVar.Title = c.PostForm("title")
	datafile, _ := c.FormFile("data")
	// if err := c.ShouldBind(&queryVar); err != nil {
	// 	SendResponseV(c, errno.ConvertErr(err), nil)
	// }

	data, err := datafile.Open()
	if err != nil {
		c.String(400, "文件格式错误")
		return
	}
	defer data.Close()
	queryVar.Data, err = ioutil.ReadAll(data)
	if err != nil {
		c.String(400, "文件错误")
		return
	}

	req := &video.DouyinPublishActionRequest{Token: queryVar.Token, Data: queryVar.Data, Title: queryVar.Title}

	err = rpc.PublishVideo(context.Background(), req)
	if err != nil {
		SendVideoResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendVideoResponse(c, errno.Success, nil)
}
