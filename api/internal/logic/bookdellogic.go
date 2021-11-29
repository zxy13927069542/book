package logic

import (
	"context"
	"fmt"

	"book/api/internal/svc"
	"book/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BookDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBookDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookDelLogic {
	return BookDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//删除
func (l *BookDelLogic) BookDel(req types.DelReq) (*types.DelResp, error) {
	// todo: add your logic here and delete this line
	condition := make(map[string]string)
	//对应数据库中的字段
	if req.Id != "" {
		condition["book_id"] = req.Id
	}
	if req.Name != "" {
		condition["book_name"] = req.Name
	}
	err := l.svcCtx.BookModel.BookDel(condition)
	if err != nil {
		fmt.Println("删除失败", err)
		return &types.DelResp{
			Message: "删除失败",
		}, err
	}
	return &types.DelResp{
		Message: "删除成功",
	}, nil
}
