package main

import (
	"fmt"
	"path"
)

//path包提供对于带有斜杠的文件路径的基本操作
//此外还有一个filepath包，此包提供的是更加通用的，兼容各操作系统的对于文件路径的操作方法
func main() {

	filePath := "Go/GoSourceCode/path/cf.cfg"

	fmt.Println(path.IsAbs(filePath)) //判断文件路径是不是绝对路径
	dir, file := path.Split(filePath) //从一个路径中获取文件名和当前文件所在路径
	fmt.Println(dir, file)
	lolpath := path.Join("Go", "GoSource", "bytes", "lol.cfg") //以斜杠组合提供的字符串，组合为一个文件路径，注意如果带有文件名，文件名需要放在最后一位
	fmt.Println(lolpath)
	fmt.Println(path.Dir(filePath))   //获取文件路径所在的目录
	fmt.Println(path.Base(filePath))  //获取文件路径的文件名
	fmt.Println(path.Ext(filePath))   //获取文件扩展名
	fmt.Println(path.Clean(filePath)) //这个应该比较少用
}
