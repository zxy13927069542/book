package logic

import (
	"book/models/system"
	"book/utils"
	"context"
	"fmt"

	"book/api/internal/svc"
	"book/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BookUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBookUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookUpdateLogic {
	return BookUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookUpdateLogic) BookUpdate(req types.UpdateReq) (*types.UpdateResp, error) {
	// todo: add your logic here and delete this line
	//传递参数
	book := system.Book{
		BookId:   utils.Str2Int(req.Id),
		BookName: req.Name,
	}
	err := l.svcCtx.BookModel.BookUpdate(&book)
	if err != nil {
		fmt.Println("更新失败", err)
		return &types.UpdateResp{
			Message: "更新失败",
		}, err
	}
	return &types.UpdateResp{
		Message: "更新成功",
	}, nil
}
