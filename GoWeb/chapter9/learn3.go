package main

import (
	"fmt"
	"strconv"
	"time"
)

//使用channel进行goroutine之间的信息传递
//可以使用无缓冲和有缓冲通道分别试试
func learn3() {

	c := make(chan int)
	go throwNum(c)
	go printNum(c)
	time.Sleep(100 * time.Millisecond)
}

//生成数字
func throwNum(c chan int) {
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("throw:" + strconv.Itoa(i))
	}
}

//打印数字
func printNum(c chan int) {
	for n := range c {
		fmt.Println("print:" + strconv.Itoa(n))
	}
}
