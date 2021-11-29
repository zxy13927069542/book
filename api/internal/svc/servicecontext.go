package svc

import (
	"book/config"
	"book/models"
)

type ServiceContext struct {
	Config    config.Config
	BookModel *models.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		BookModel: models.NewBookModel(),
	}
}
