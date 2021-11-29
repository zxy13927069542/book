package utils_test

import (
	"book/utils"
	"fmt"
	"testing"
)

func TestMap2SqlCondition(t *testing.T) {
	//m := map[string]string {
	//	"id": "1",
	//}
	m := map[string]string{
		"name": "1234",
	}

	condition := utils.Map2SqlCondition(m)
	fmt.Println(condition)
}
