package logic

import (
	"context"

	"gozerodemo/userrpc/internal/svc"
	"gozerodemo/userrpc/sdk/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	logx.Infof("%v", *in)
	return &user.UserResponse{
		Id:     in.GetId(),
		Name:   "test",
		Gender: "N/A",
	}, nil
}
