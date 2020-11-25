package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

//crypto包内提供的是一些加密算法相关的操作，其下的md5提供的md5加密算法的操作
func main() {

	fmt.Println(MD5("520wcj"))
}

//MD5 使用md5包实现MD5加密
func MD5(pwd string) string {
	h := md5.New()                        //返回一个hash.Hash接口
	h.Write([]byte(pwd))                  //写入原始密码
	return hex.EncodeToString(h.Sum(nil)) //生成校验和
}
