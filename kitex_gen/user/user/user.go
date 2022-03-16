// Code generated by Kitex v0.2.0. DO NOT EDIT.

package user

import (
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"user/kitex_gen/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceInfo
}

var userServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "User"
	handlerType := (*user.User)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetUser": kitex.NewMethodInfo(getUserHandler, newUserGetUserArgs, newUserGetUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.2.0",
		Extra:           extra,
	}
	return svcInfo
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserGetUserArgs)
	realResult := result.(*user.UserGetUserResult)
	success, err := handler.(user.User).GetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserGetUserArgs() interface{} {
	return user.NewUserGetUserArgs()
}

func newUserGetUserResult() interface{} {
	return user.NewUserGetUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetUser(ctx context.Context, req *user.GetUserRequest) (r *user.GetUserResponse, err error) {
	var _args user.UserGetUserArgs
	_args.Req = req
	var _result user.UserGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
