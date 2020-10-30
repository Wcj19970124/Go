package main

import (
	"fmt"
	"log"
	"os"
)

//log日志包，主要用于日志输出，并且其是支持并发情况下的安全输出的：因为Logger结构体中使用了
//sync.Mutex保证输出操作的原子性
func learn1() {
	logFile, _ := os.Create("log.txt")
	fmt.Println(log.Flags()) //返回logger的标准选项
	// log.SetFlags(2)          //设置logger的标准选项
	log.SetPrefix("Warning:") //设置logger的标准输出前缀
	fmt.Println(log.Prefix()) //返回logger的标准输出前缀
	log.SetOutput(logFile)    //设置日志的输出地,可以os.Create()创建一个文件，然后让log输出到文件中
	log.Printf("%s", "牛皮啊")
	log.Print("世界很美,你想去看吗")
	log.Println("算了吧，生命这么贵，万一丢了可不好")
	log.Fatal("输出完就结束吧") //Fatal()这一系列的输出，相当于先输出然后os.Exit(1)
	log.Fatalf("%s", "想结束？没门")
	log.Fatalln("我就呵呵")
	log.Panic("我是错误，我来了") //Panic()这一类相当于输出之后，panic()抛错
	log.Panicf("%s", "你还是走吧")
	log.Panicln("你想上哪去？上天吗？")
}
