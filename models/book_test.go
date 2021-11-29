package models_test

import (
	"book/config"
	"book/models"
	"book/models/system"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"reflect"
	"testing"
)

var configFile = flag.String("a", "../config/book-api.yaml", "the config file")

func LoadConfig() {
	flag.Parse()
	c := config.SetConfig()
	//如果已有配置，则不再覆盖
	equal := reflect.DeepEqual(*c, config.Config{})
	if equal {
		//从运行指令获取参数
		conf.MustLoad(*configFile, c) //从默认配置文件或者指令的配置文件读取配置
	}

}

func TestUniqueQuery(t *testing.T) {
	//加载配置
	LoadConfig()
	models.SetupDataBase()

	condition := map[string]string{
		"book_name": "大学英语4",
	}
	query, err := models.NewBookModel().UniqueQuery(condition)
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	fmt.Println(query)

}

func TestBookAdd(t *testing.T) {
	//加载配置
	LoadConfig()
	models.SetupDataBase()

	book := system.Book{
		BookName: "大学物理3",
	}
	err := models.NewBookModel().BookAdd(&book)
	if err != nil {
		fmt.Println("插入失败", err)
	}
}

func TestBookDel(t *testing.T) {
	//加载配置
	LoadConfig()
	models.SetupDataBase()

	condition := map[string]string{
		"book_name": "大学物理1",
	}
	err := models.NewBookModel().BookDel(condition)
	if err != nil {
		fmt.Println("删除失败", err)
	}
}

func TestBookUpdate(t *testing.T) {
	//加载配置
	LoadConfig()
	models.SetupDataBase()

	book := &system.Book{
		BookId:   33,
		BookName: "小学拼音",
	}
	err := models.NewBookModel().BookUpdate(book)
	if err != nil {
		fmt.Println("删除失败", err)
	}
}

func TestBookGetAll(t *testing.T) {
	//加载配置
	LoadConfig()
	models.SetupDataBase()

	books, err := models.NewBookModel().BookGetAll()
	if err != nil {
		fmt.Println("查询所有失败", err)
		return
	}
	for _, book := range books {
		fmt.Println(book)
	}
}
