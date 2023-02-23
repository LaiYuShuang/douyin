package main

import (
	"douyin/cmd/api/handles"
	"douyin/cmd/api/middleware"
	"douyin/cmd/api/rpc"
	"douyin/pkg/tracer"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	// TODO
	tracer.InitJaeger("api")
	rpc.InitRPC()
}

func main() {

	Init()
	r := gin.Default()
	r.Use(middleware.OpenTracing())
	r.Static("/resource", "./resource")
	douyin := r.Group("/douyin")

	userGroup := douyin.Group("/user")
	userGroup.POST("/login/", handlers.Login)
	userGroup.POST("/register/", handlers.Register)
	userGroup.GET("/", middleware.AuthMiddleware(), handlers.QueryCurUser)

	douyin.GET("/feed/", handlers.GetFeed)
	douyin.GET("/publish/list/", handlers.GetPublishList)
	douyin.POST("/publish/action/", handlers.PublishVideo)

	favoriteGroup := douyin.Group("/favorite")
	favoriteGroup.GET("/list/", handlers.GetFavoriteList)
	favoriteGroup.POST("/action/", handlers.FavoriteVideo)

	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		klog.Fatal(err)
	}

}
