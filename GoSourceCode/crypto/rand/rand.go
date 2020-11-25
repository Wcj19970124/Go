package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

//与math/rand包不同,前者提供的是伪随机数的生成，常用于验证码，随机密码，随机数等
//而crypto/rand包提供的用于加解密的安全随机数，其下三个方法:
func main() {

	n, err := rand.Int(rand.Reader, big.NewInt(128)) //返回一个位于[0,max]的随机数
	if err == nil {
		fmt.Println("获得的随机数为：", n)
	}

	m, err := rand.Prime(rand.Reader, 5) //返回一个指定位数的随机数质数
	if err == nil {
		fmt.Println("获得指定位数的随机数，大概率为质数：", m)
	}

	x, err := rand.Read([]byte("1234")) //辅助函数，只有当err == nil 时，x == len([]byte)
	if err == nil {
		fmt.Println("字节切片的长度为:", x)
	}

}
