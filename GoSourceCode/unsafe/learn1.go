package main

import (
	"fmt"
	"unsafe"
)

type tour struct {
	name   string
	gender string
}

//unsafe包主要是提供了几个跳过go安全类型检查的方法
//一般不会使用到这个包，而且使用这个包是不遵循一系列协议的，所以慎用
func main() {

	var t tour
	fmt.Println(unsafe.Sizeof(t))          //看类型所占的内存大小
	fmt.Println(unsafe.Offsetof(t.gender)) //看字段的偏移量
	fmt.Println(unsafe.Alignof(t))         //查看类型在内存中的对齐方式
}
