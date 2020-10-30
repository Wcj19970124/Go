package main

import (
	"fmt"
	"time"
)

//Duration对象表示一个时间段
//time是专为一个时间点定义的，但是Duration是专为一个时间段定义
//重要的方法：ParseDuration(),Since(),以及获取时分秒等方法
func learn2() {

	// var d time.Duration

	d, _ := time.ParseDuration("1h45m") //将一个字符串解析为Duration对象
	fmt.Println(d)
	today, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-30 17:54:35", time.Local) //注意时区的问题
	fmt.Println(time.Since(today))                                                             //Since()函数，用于获取从t到现在经过的时间，返回值为对象
	fmt.Println(d.Hours(), d.Minutes(), d.Seconds(), d.Milliseconds())                         //获取Duration对象的属性
}
