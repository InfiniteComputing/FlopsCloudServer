package logic

import (
	"context"

	"gzfcserver/service/product/model"
	"gzfcserver/service/product/rpc/internal/svc"
	"gzfcserver/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	// todo: add your logic here and delete this line
	newProduct := &model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
		Status: in.Status,
	}

	res, err := l.svcCtx.ProductModel.Insert(l.ctx, newProduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newProductId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.CreateResponse{
		Id: uint64(newProductId),
	}, nil
}
