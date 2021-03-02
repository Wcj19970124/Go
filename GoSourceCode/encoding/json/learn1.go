package main

import (
	"encoding/hex"
	"fmt"
)

//测试将十六进制字符串转换为ASCII字符
func learn1() {

	str := "48656c6c6f20576f726c6421"
	signBin, err := hex.DecodeString(str)
	if err != nil {
		fmt.Println("hex decode error:" + err.Error())
	} else {
		fmt.Println(string(signBin))
	}

}
