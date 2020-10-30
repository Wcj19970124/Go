package main

import (
	"fmt"
)

//Errorf():将参数按照指定format格式化，并且返回一个error
func main() {

	err := fmt.Errorf("%t", 1 == 2)
	if err != nil {
		fmt.Println(err.Error())
	}
}
