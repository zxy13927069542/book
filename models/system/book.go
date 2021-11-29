package system

type Book struct {
	BookId   int    `gorm:"primaryKey;autoIncrement;comment:书籍编号"`
	BookName string `gorm:"unique"`
}

func (b Book) TableName() string {
	return "book"
}
