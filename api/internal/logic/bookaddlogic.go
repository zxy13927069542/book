package logic

import (
	"book/models/system"
	"context"

	"book/api/internal/svc"
	"book/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BookAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBookAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookAddLogic {
	return BookAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/**
插入一条记录并查询该记录
*/
func (l *BookAddLogic) BookAdd(req types.AddReq) (*types.AddResp, error) {
	// todo: add your logic here and delete this line
	book := system.Book{
		BookName: req.Name,
	}
	err := l.svcCtx.BookModel.BookAdd(&book)
	if err != nil {
		return &types.AddResp{
			Message: "插入失败",
		}, err
	}
	return &types.AddResp{
		Message: "插入成功",
	}, nil
}
