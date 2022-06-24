package logic

import (
	"context"
	"fmt"
	"gozerodemo/testapi/internal/svc"
	"gozerodemo/testapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestapiLogic {
	return &TestapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestapiLogic) Testapi(req *types.Request) (resp *types.Response, err error) {
	//name := req.Name
	//logx.Info(name)
	//resp.Message = fmt.Sprintf("hello, %s", name)
	resp = new(types.Response)
	resp.Message = fmt.Sprintf("hello, %s", req.Name)
	return
}
