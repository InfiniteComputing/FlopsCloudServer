package logic

import (
	"context"

	"gzfcserver/service/product/api/internal/svc"
	"gzfcserver/service/product/api/internal/types"
	"gzfcserver/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:     res.Id,
		Name:   res.Name,
		Desc:   res.Desc,
		Amount: res.Amount,
		Stock:  res.Stock,
		Status: res.Status,
	}, nil
}
