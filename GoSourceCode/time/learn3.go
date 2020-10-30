package main

import (
	"fmt"
	"time"
)

//Sleep()函数，Tick()函数
func main() {
	tick := time.Tick(1 * time.Second) //Tick()返回一个通道，通常可以当作定时器使用
	for range tick {
		time.Sleep(100 * time.Millisecond) //让goroutine睡眠
		fmt.Println("我是不是很像定时器")
	}
}
