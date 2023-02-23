package main

import (
	"context"
	"douyin/cmd/user/pack"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/middleware"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		resp.UserId = constants.EmptyUserId
		resp.Token = constants.EmptyToken
		return resp, nil
	}

	uid, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UserId = constants.EmptyUserId
		resp.Token = constants.EmptyToken
		return resp, nil
	}

	// token, _ := global.Jwt.CreateToken(userID, global.JWTSetting.AppKey, global.JWTSetting.AppSecret)
	token, err := middleware.CreateToken(uid)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UserId = constants.EmptyUserId
		resp.Token = constants.EmptyToken
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = uid
	resp.Token = token
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		resp.UserId = constants.EmptyUserId
		resp.Token = constants.EmptyToken
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UserId = constants.EmptyUserId
		resp.Token = constants.EmptyToken
		return resp, nil
	}

	// token, _ := global.Jwt.CreateToken(userID, global.JWTSetting.AppKey, global.JWTSetting.AppSecret)
	token, err := middleware.CreateToken(uid)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UserId = constants.EmptyUserId
		resp.Token = constants.EmptyToken
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = uid
	resp.Token = token
	return resp, nil
}

// QueryCurUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryCurUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserResponse)

	if req.UserId == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		resp.User = nil
		return resp, nil
	}

	users, err := service.NewQueryCurUserService(ctx).QueryCurUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.User = nil
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.User = users[0]
	return resp, nil
}
