package main

import (
	"container/list"
	"fmt"
)

//此包提供的双向链表的实现
//使用起来比较简单，另外还有一个ring包提供环形链表
//api也比较简单，这里就不举例了
func main() {

	l := list.New()                              //新建一个双向链表
	l.PushBack(2)                                //从链表末尾插入元素
	l.PushFront(6)                               //从链表头处插入元素
	for e := l.Front(); e != nil; e = e.Next() { //双向链表的遍历
		fmt.Println(e.Value) //每个元素都是一个Element的结构体，内部的Value代表元素节点的值
	}
}
