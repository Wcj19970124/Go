package main

import "fmt"

//使用channel实现goroutine之间同步
func learn2() {

	w1, w2 := make(chan bool), make(chan bool)
	go printLow(w1)
	go printHigh(w2)
	<-w1
	<-w2
}

func printLow(c chan bool) {
	for i := 'a'; i < 'a'+26; i++ {
		fmt.Printf("%c", i)
	}
	c <- true
}

func printHigh(c chan bool) {
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c", i)
	}
	c <- true
}
