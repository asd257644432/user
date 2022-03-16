package handler

import (
	"context"
	"user/kitex_gen/base"
	"user/kitex_gen/user"
	"user/logic"
	"user/util"
)

type GetUserHandler struct {
	ctx context.Context
}

func NewGetUserHandler(ctx context.Context) *GetUserHandler {
	return &GetUserHandler{ctx: ctx}
}

func (h *GetUserHandler) Handle(req *user.GetUserRequest) *user.GetUserResponse {
	resp := user.NewGetUserResponse()
	resp.BaseResp = base.NewBaseResp()

	user, err := logic.NewUserLogic(h.ctx).GetUser(req.GetUserId())
	if err != nil {
		resp.BaseResp = util.PackBaseRespWithError(err)
		return resp
	}
	resp.User = user
	return resp
}
