package logic

import (
	"context"
	"user/dal"
	"user/kitex_gen/common"
	"user/pack"
)

type UserLogic struct {
	ctx context.Context
}

func NewUserLogic(ctx context.Context) *UserLogic {
	return &UserLogic{ctx: ctx}
}

func (l *UserLogic) GetUser(id int64) (*common.User, error) {
	user, err := dal.NewUserDal(l.ctx).GetUser(id)
	if err != nil {
		return nil, err
	}
	return pack.PackUserDBToRpc(user), nil
}
