package util

import "user/kitex_gen/base"

func PackBaseRespWithError(err error) *base.BaseResp {
	return &base.BaseResp{
		StatusCode:    50000,
		StatusMessage: err.Error(),
	}
}
