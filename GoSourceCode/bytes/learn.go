package main

import (
	"bytes"
	"fmt"
)

//bytes包主要提供的就是对于[]byte切片操作的诸多方法和函数
//和strings包中的函数方法基本一致
func main() {

	a := []byte("1234")
	b := []byte("5678")
	c := []byte("12")
	d := []byte("AB")
	fmt.Println(bytes.Compare(a, b)) //按照字典序对两个[]byte进行对比,返回-1,0,1
	fmt.Println(bytes.Equal(a, b))   //对比两个[]byte，返回bool
	run := bytes.Runes(a)            //将[]byte-》[]rune
	fmt.Println(run)
	fmt.Println(bytes.HasPrefix(a, b)) //前后缀是否为某个前缀切片/后缀切片
	fmt.Println(bytes.Contains(a, b))  //是否包含某个[]byte
	fmt.Println(bytes.Index(a, c))     //子切片在主切片中的第一次出现索引，不存在返回-1，还有LastIndex()
	fmt.Println(bytes.ToLower(d))  //大小写转换 
}
