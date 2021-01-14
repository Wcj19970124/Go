package main

import (
	"fmt"
	"os"
)

func learn2() {
	p, err := os.FindProcess(31982)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p)
	// p.Signal()  //向进程发送一个信号
	// p.Kill() //杀死该进程
	// p.Wait() //让进程处于等待状态,返回的是一个ProcessState结构体，该结构体实际上就是保存wait函数中退出的进程的信息
	// p.Release() //释放进程资源
	// os.StartProcess()  //启动一个进程

}
