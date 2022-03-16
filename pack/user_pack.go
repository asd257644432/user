package pack

import (
	"user/kitex_gen/common"
	"user/model"
)

func PackUserDBToRpc(user *model.User) *common.User {
	return &common.User{
		Id:       user.Id,
		Passport: user.Passport,
		Password: user.Password,
	}
}

func PackUserRpcToDB(user *common.User) *model.User {
	return &model.User{
		Id:       user.Id,
		Passport: user.Passport,
		Password: user.Password,
	}
}
