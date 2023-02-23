package main

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
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

	c, err := userservice.NewClient(
		constants.UserServiceName,
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

	req := &user.DouyinUserRegisterRequest{Username: "111111", Password: "124"}
	resp, err := c.CreateUser(context.Background(), req)

	req1 := &user.DouyinUserLoginRequest{Username: "111111", Password: "124"}
	resp1, err := c.CheckUser(context.Background(), req1)

	req2 := &user.DouyinUserRequest{UserId: 1, Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjIyOTg3NjkyMjMsImlhdCI6MTY3NjY4OTIyMywiaXNzIjoiZG91eWluIn0.Yot75U-ZZlXwUfaXEEag6A2Vf-U4LL-wO3RKb07KQlA"}
	resp2, err := c.QueryCurUser(context.Background(), req2)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	log.Println(resp1)
	log.Println(resp2)
	time.Sleep(time.Second)
}
