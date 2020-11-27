package main

import (
	"flag"
	"fmt"
)

var ip = flag.Int("flagname", 1234, "help message") //method1:此方法返回的是一个命令行参数指针

var flagType string

func init() {
	flag.StringVar(&flagType, "flagPe", "help me", "sos") //method2：将一个命令行参数和一个变量绑定到一块，返回的就是值类型
}

//此包提供命令行参数的解析功能
//实际上就是对命令行中的参数,os.Args进行处理
func learn() {

	//第二步：解析命令行参数，在碰到第一个非命令行参数的时候停止
	flag.Parse()

	//第三步：直接使用命令行参数
	fmt.Println("help message:", *ip)
	fmt.Println("sos:", flagType)
}
