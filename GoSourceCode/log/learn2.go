package main

import (
	"log"
	"os"
)

//内嵌结构体之后，内嵌结构体的属性和方法都会被自动提升到外部结构体
//可以利用log.Logger自定义日志级别，然后使用
//不过以下的这种定义并不好，因为通常来说都是直接定义为方法
type wcjlog struct {
	Error *log.Logger
	Warn  *log.Logger
	Debug *log.Logger
	Info  *log.Logger
	Trace *log.Logger
}

var cjlog wcjlog

func init() {
	cjlog.Error = log.New(os.Stdout, "error:", log.Ldate|log.Ltime|log.Llongfile|log.Lmsgprefix)
	cjlog.Warn = log.New(os.Stdout, "warn:", log.Ldate|log.Ltime|log.Llongfile|log.Lmsgprefix)
	cjlog.Debug = log.New(os.Stdout, "debug:", log.Ldate|log.Ltime|log.Llongfile|log.Lmsgprefix)
	cjlog.Info = log.New(os.Stdout, "info:", log.Ldate|log.Ltime|log.Llongfile|log.Lmsgprefix)
	cjlog.Trace = log.New(os.Stdout, "trace:", log.Ldate|log.Ltime|log.Llongfile|log.Lmsgprefix)
}

func main() {
	cjlog.Error.Println("我是错误")
	cjlog.Warn.Println("我是警告")
	cjlog.Debug.Println("我是debug")
	cjlog.Info.Println("我是提示")
	cjlog.Trace.Println("我是细节")
}
