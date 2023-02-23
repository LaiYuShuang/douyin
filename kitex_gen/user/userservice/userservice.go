// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "douyin/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser":   kitex.NewMethodInfo(createUserHandler, newUserServiceCreateUserArgs, newUserServiceCreateUserResult, false),
		"CheckUser":    kitex.NewMethodInfo(checkUserHandler, newUserServiceCheckUserArgs, newUserServiceCheckUserResult, false),
		"QueryCurUser": kitex.NewMethodInfo(queryCurUserHandler, newUserServiceQueryCurUserArgs, newUserServiceQueryCurUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCreateUserArgs)
	realResult := result.(*user.UserServiceCreateUserResult)
	success, err := handler.(user.UserService).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserArgs() interface{} {
	return user.NewUserServiceCreateUserArgs()
}

func newUserServiceCreateUserResult() interface{} {
	return user.NewUserServiceCreateUserResult()
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCheckUserArgs)
	realResult := result.(*user.UserServiceCheckUserResult)
	success, err := handler.(user.UserService).CheckUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCheckUserArgs() interface{} {
	return user.NewUserServiceCheckUserArgs()
}

func newUserServiceCheckUserResult() interface{} {
	return user.NewUserServiceCheckUserResult()
}

func queryCurUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceQueryCurUserArgs)
	realResult := result.(*user.UserServiceQueryCurUserResult)
	success, err := handler.(user.UserService).QueryCurUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceQueryCurUserArgs() interface{} {
	return user.NewUserServiceQueryCurUserArgs()
}

func newUserServiceQueryCurUserResult() interface{} {
	return user.NewUserServiceQueryCurUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (r *user.DouyinUserRegisterResponse, err error) {
	var _args user.UserServiceCreateUserArgs
	_args.Req = req
	var _result user.UserServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (r *user.DouyinUserLoginResponse, err error) {
	var _args user.UserServiceCheckUserArgs
	_args.Req = req
	var _result user.UserServiceCheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryCurUser(ctx context.Context, req *user.DouyinUserRequest) (r *user.DouyinUserResponse, err error) {
	var _args user.UserServiceQueryCurUserArgs
	_args.Req = req
	var _result user.UserServiceQueryCurUserResult
	if err = p.c.Call(ctx, "QueryCurUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
