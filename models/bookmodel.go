package models

import (
	"book/models/system"
	"book/utils"
	"gorm.io/gorm"
)

type BookModel struct {
	db *gorm.DB
}

func NewBookModel() *BookModel {
	return &BookModel{
		db: GetDB(),
	}
}

//  UniqueQuery
//  @Description: 根据唯一条件查询，只会得到一条结果
//  @param condition：想要删除的条件，包括book_id,book_name
//  @return *system.Book
//  @return error
//
func (m *BookModel) UniqueQuery(condition map[string]string) (*system.Book, error) {
	var book system.Book
	sqlCondition := utils.Map2SqlCondition(condition)
	error := m.db.Table(book.TableName()).Where(sqlCondition).First(&book).Error
	if error != nil {
		return &book, error
	}
	return &book, nil
}

//
//  BookAdd
//  @Description: 新增一行
//  @param book：book中只有BookName有值，BookId是主键自增，无需设置
//  @return error
//
func (m *BookModel) BookAdd(book *system.Book) error {
	error := m.db.Table(book.TableName()).Create(book).Error
	if error != nil {
		return error
	}
	return nil
}

//
//  BookDel
//  @Description: 删除操作
//  @param condition 想要删除的条件，包括book_id,book_name
//  @return error
//
func (m *BookModel) BookDel(condition map[string]string) error {
	var book system.Book
	sqlCondition := utils.Map2SqlCondition(condition)
	error := m.db.Table(book.TableName()).Where(sqlCondition).Delete(&book).Error
	if error != nil {
		return error
	}
	return nil
}

func (m *BookModel) BookUpdate(book *system.Book) error {
	err := m.db.Table(book.TableName()).Updates(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *BookModel) BookGetAll() ([]system.Book, error) {
	var books []system.Book
	err := m.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
