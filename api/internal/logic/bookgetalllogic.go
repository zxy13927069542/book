package logic

import (
	"book/utils"
	"context"
	"fmt"

	"book/api/internal/svc"
	"book/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type BookGetAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBookGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookGetAllLogic {
	return BookGetAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
//  BookGetAll
//  @Description: 获取book表所有记录
//  @receiver l
//  @param req
//  @return *types.GetAllResp
//  @return error
//
func (l *BookGetAllLogic) BookGetAll(req types.GetAllReq) (*types.GetAllResp, error) {
	// todo: add your logic here and delete this line

	var books []types.Book
	all, err := l.svcCtx.BookModel.BookGetAll()
	if err != nil {
		fmt.Println("获取所有失败", err)
		return nil, err
	}
	for _, book := range all {
		books = append(books, types.Book{
			Id:   utils.Int2str(book.BookId),
			Name: book.BookName,
		})
	}
	return &types.GetAllResp{
		Books: books,
	}, nil
}
