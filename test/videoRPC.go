package main

import (
	"context"
	"douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"
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

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
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

	req := &video.DouyinPublishActionRequest{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjIyOTg3NjkyMjMsImlhdCI6MTY3NjY4OTIyMywiaXNzIjoiZG91eWluIn0.Yot75U-ZZlXwUfaXEEag6A2Vf-U4LL-wO3RKb07KQlA", Data: []byte{0}, Title: "mytest"}
	resp, err := c.PublishAction(context.Background(), req)

	req1 := &video.DouyinPublishListRequest{UserId: 1, Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjIyOTg3NjkyMjMsImlhdCI6MTY3NjY4OTIyMywiaXNzIjoiZG91eWluIn0.Yot75U-ZZlXwUfaXEEag6A2Vf-U4LL-wO3RKb07KQlA"}
	resp1, err := c.PublishList(context.Background(), req1)

	var t int64 = 0
	var token string = ""
	req2 := &video.DouyinFeedRequest{LatestTime: &t, Token: &token}
	resp2, err := c.Feed(context.Background(), req2)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	log.Println(resp1)
	log.Println(resp2)
	time.Sleep(time.Second)
}
