package utils

import (
	"fmt"
	"strconv"
)

/**
由map直接生成对应的sql条件
*/
func Map2SqlCondition(m map[string]string) string {
	sqlcondition := ""
	for key, value := range m {
		sqlcondition = sqlcondition + "`" + key + "` = "
		sqlcondition = sqlcondition + "'" + value + "' and "
	}

	//裁剪掉最后一个多余的and
	sqlcondition = string([]byte(sqlcondition)[:len(sqlcondition)-4])
	return sqlcondition
}

//
//  Int2str
//  @Description: 将整型转成字符串
//  @param num
//  @return string
//
func Int2str(num int) string {
	return strconv.Itoa(num)
}

//
//  Str2Int
//  @Description: 字符串转整型
//  @param s
//  @return int
//
func Str2Int(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("转型失败")
		return 0
	}
	return result
}
