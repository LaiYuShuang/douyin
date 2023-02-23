package main

import (
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/favorite/favoriteservice"

	"context"
	"douyin/pkg/constants"
	"douyin/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"time"
)

func main() {
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

	//is favorite
	req := &favorite.DouyinFavoriteActionRequest{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjIyOTg3NjkyMjMsImlhdCI6MTY3NjY4OTIyMywiaXNzIjoiZG91eWluIn0.Yot75U-ZZlXwUfaXEEag6A2Vf-U4LL-wO3RKb07KQlA", ActionType: 1, VideoId: 1}
	resp, err := c.FavoriteAction(context.Background(), req)

	//get list
	req2 := &favorite.DouyinFavoriteListRequest{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjIyOTg3NjkyMjMsImlhdCI6MTY3NjY4OTIyMywiaXNzIjoiZG91eWluIn0.Yot75U-ZZlXwUfaXEEag6A2Vf-U4LL-wO3RKb07KQlA", UserId: 1}
	resp2, err := c.FavoriteList(context.Background(), req2)

	//un favorite
	req1 := &favorite.DouyinFavoriteActionRequest{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjIyOTg3NjkyMjMsImlhdCI6MTY3NjY4OTIyMywiaXNzIjoiZG91eWluIn0.Yot75U-ZZlXwUfaXEEag6A2Vf-U4LL-wO3RKb07KQlA", ActionType: 2, VideoId: 1}
	resp1, err := c.FavoriteAction(context.Background(), req1)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	log.Println(resp1)
	log.Println(resp2)
	time.Sleep(time.Second)
}
