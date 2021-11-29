package logic

import (
	"book/utils"
	"context"

	"book/api/internal/svc"
	"book/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BookGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBookGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookGetLogic {
	return BookGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookGetLogic) BookGet(req types.QueryReq) (*types.QueryResp, error) {
	// todo: add your logic here and delete this line
	condition := make(map[string]string)
	//对应数据库中的字段
	if req.Id != "" {
		condition["book_id"] = req.Id
	}
	if req.Name != "" {
		condition["book_name"] = req.Name
	}
	book, error := l.svcCtx.BookModel.UniqueQuery(condition)
	if error != nil {
		return &types.QueryResp{
			Message: "查询失败",
		}, error
	}
	return &types.QueryResp{
		Book: types.Book{
			Id:   utils.Int2str(book.BookId),
			Name: book.BookName,
		},
		Message: "查询成功",
	}, nil

}
