package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//并发和并行:Go语言这两个概念和Java这种依托于线程的略微有点区别
//类似于Java这类语言，用于并发和并行的是线程，而线程实际上是直接映射到OS上
//和cpu进行绑定的，但是Go中的goroutine和线程并不一样，goroutine是和线程进行绑定的
//或者说的更准确点是和逻辑处理器进行绑定的，而每个线程和一个逻辑处理器进行绑定的
//所以并行的概念只存在于多个逻辑处理器的情况下，但是使用的并不多，使用的较多的就是并发
//Go1.5之前，默认程序是使用一个逻辑处理器；1.5之后默认逻辑处理器数和cpu数一致
//使用WaitGroup等待组来保证goroutine之间的同步
func learn1() {
	runtime.GOMAXPROCS(1)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()
		for i := 'a'; i < 'a'+26; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%c", i)
		}
	}()

	go func() {
		defer waitGroup.Done()
		for i := 'A'; i < 'A'+26; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%c", i)
		}
	}()

	waitGroup.Wait()

	fmt.Println("main goroutine exit")
}
