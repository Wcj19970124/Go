package main

import (
	"fmt"
	"sort"
)

//Person 结构体
type Person struct {
	Name string
	Age  int
}

//实现String()接口
func (p Person) String() string {
	return fmt.Sprintf("%s:%d", p.Name, p.Age)
}

//ByAge 按年龄排序的Person数组
type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

//此包提供了对于切片和自定义数据集的排序功能
//示例一：自定义切片，并且用内部的字段进行排序
func learn1() {

	b := ByAge{
		{"wcj", 24},
		{"hhy", 23},
		{"ysj", 22},
		{"zjq", 27},
	}

	fmt.Println(b)
	sort.Sort(b)
	fmt.Println(b)
}
