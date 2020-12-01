package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup sync.WaitGroup

//sync包下提供的一般的同步方法，例如互斥锁，读写互斥锁，Once,WaitGroup等
//除了Once/WaitGroup都是低水平线程的操作，更多的还是使用goroutine，不过这个还是有必要看一下的
func main() {
	runtime.GOMAXPROCS(1)

	waitGroup.Add(2)

	chan1 := make(chan int)

	go send(chan1)

	go receive(chan1)

	waitGroup.Wait()

	fmt.Println("finish.......")
}

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	defer waitGroup.Done()
}

func receive(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
	defer waitGroup.Done()
}
