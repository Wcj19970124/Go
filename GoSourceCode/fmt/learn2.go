package main

import (
	"fmt"
	"io"
)

//fmt.Print()函数：
//Print(),Println(),Printf():将参数内容按照标准格式输出
//Sprint(),Sprintln(),Sprintf():将参数串联成字符串，返回一个字符串
//Fprint(),Fprintln(),Fprintf():将参数串联成字符串，写入io.Writer中
func learn2() {

	//Print()这种类型：按照默认格式输出
	//Println():输出之后进行换行
	//Printf():将输出内容按照指定参数进行格式化之后再输出
	str := "一只羊 两只羊\n 三只羊"
	s := "一头猪 两头猪"
	fmt.Print(str, s)
	fmt.Println(str, s)
	fmt.Printf("%v,%s", str, s)

	//Sprint():串连参数，如果参数都不是字符串，自动添加空格，返回一个字符串
	//Sprintln():串联参数，并且总是添加空格，最后添加换行，返回一个字符串
	//Sprintf():串联参数，按照指定format对参数进行格式化，返回字符串
	s1 := fmt.Sprint(str, s)
	fmt.Println(s1)
	s2 := fmt.Sprintln(str, s)
	fmt.Println(s2)
	s3 := fmt.Sprintf("%s %s", str, s)
	fmt.Println(s3)

	//Fprint():串联参数，都不是字符串，自动添加空格，最后写入io.Writer中，例如写入ResponseWriter
	//Fprintln():末尾添加换行
	//Fprintf():将参数按照指定format进行格式化，然后写入到io.Writer
	var w io.Writer
	fmt.Fprint(w, str, s)

}
