package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/middleware"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
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
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (int64, string, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return -1, "", err
	}
	if resp.BaseResp.StatusCode != 0 {
		return -1, "", errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, resp.Token, nil
}

// CheckUser check user info
func CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (int64, string, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return -1, "", err
	}
	if resp.BaseResp.StatusCode != 0 {
		return -1, "", errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, resp.Token, nil
}

// UserInfo user info format
type UserInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// QuryUser check user info
func QueryUser(ctx context.Context, req *user.DouyinUserRequest) (*UserInfo, error) {
	resp, err := userClient.QueryCurUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	var userInfo UserInfo
	userInfo.ID = resp.User.Id
	userInfo.Name = resp.User.Name
	return &userInfo, nil
}
