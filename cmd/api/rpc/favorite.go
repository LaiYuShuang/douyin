package rpc

import (
	"context"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/favorite/favoriteservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/middleware"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

// CreateNote create note info
func FavoriteVideo(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) error {

	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseRasp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseRasp.StatusCode, resp.BaseRasp.StatusMessage)
	}
	return nil
}

// ThumbList query list of video
func GetFavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) ([]*favorite.Video, error) {
	resp, err := favoriteClient.FavoriteList(ctx, req)
	fmt.Println("客户端返回:ThumbList:", resp)
	if err != nil {
		return nil, err
	}
	if resp.BaseRape.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseRape.StatusCode, resp.BaseRape.StatusMessage)
	}
	return resp.VideoList, nil
}
