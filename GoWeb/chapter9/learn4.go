package main

import (
	"fmt"
	"time"
)

//select 实现goroutine的多路复用，select会选择准备好的那个通道进行值的接收或者发送
//当有多个通道都准备好的时候，随机选择其中的一个通道进行操作，通常利用死循环，当default分支
//执行一定次数的时候进行退出，或者for condition
func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	go channel1(c1)
	go channel2(c2)
	ok1, ok2 := true, true

	for ok1 || ok2 {
		time.Sleep(100 * time.Millisecond)
		select {
		case str, ok1 := <-c1:
			if ok1 {
				fmt.Println(str)
			}
		case str, ok2 := <-c2:
			if ok2 {
				fmt.Println(str)
			}
		default:
			fmt.Println("end")
		}
	}
}

func channel1(c chan string) {
	defer close(c)
	c <- "forever god"
}

func channel2(c chan string) {
	defer close(c)
	c <- "UZI"
}
