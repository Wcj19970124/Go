package main

import (
	"fmt"
	"strconv"
)

//Quote(),UnQuote()函数：主要就是返回字符串的双引号字面值形式和
//Parse()函数：这个见名知意，就是将字符串形式解析成为某种类型
//Format()函数：这个就是将某种类u行格式化为字符串形式  ---这里不做演示了
func main() {

	var a rune = 32

	if strconv.IsPrint(a) { //判断参数是否可以打印
		fmt.Println("可以打印")
	}

	v := "正青春"
	fmt.Println(strconv.Quote(v)) //Quote()返回字符串的双引号字面值形式,控制字符，非ASCII形式进行转义

}
