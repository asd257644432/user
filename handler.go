package main

import (
	"context"
	"user/handler"
	"user/kitex_gen/user"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// GetUser implements the UserImpl interface.
func (s *UserImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	resp = handler.NewGetUserHandler(ctx).Handle(req)
	return
}
