package service

import (
	"context"

	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/user"
)

type QueryCurUserService struct {
	ctx context.Context
}

// NewQueryCurUserService new MGetUserService
func NewQueryCurUserService(ctx context.Context) *QueryCurUserService {
	return &QueryCurUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *QueryCurUserService) QueryCurUser(req *user.DouyinUserRequest) ([]*user.User, error) {
	modelUsers, err := db.QueryUserByID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
