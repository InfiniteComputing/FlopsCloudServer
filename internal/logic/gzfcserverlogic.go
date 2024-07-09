package logic

import (
	"context"

	"gzfcserver/internal/svc"
	"gzfcserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GzfcserverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGzfcserverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GzfcserverLogic {
	return &GzfcserverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GzfcserverLogic) Gzfcserver(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.Response)
	resp.Message = req.Name

	return
}
