package main

import (
	"fmt"
	"sort"
)

//待排序切片的成员
type weight float64
type height float64
type plant struct {
	name string
	w    weight
	h    height
}

//用于指定字段排序的方法
type By func(p1, p2 plant) bool

func (by By) sort(plants []plant) {

	ps := &plantSort{
		plant: plants,
		by:    by,
	}
	sort.Sort(ps)
}

type plantSort struct {
	plant []plant
	by    func(p1, p2 plant) bool
}

func (p plantSort) Len() int {
	return len(p.plant)
}

func (p plantSort) Swap(i, j int) {
	p.plant[i], p.plant[j] = p.plant[j], p.plant[i]
}

func (p plantSort) Less(i, j int) bool {
	return p.by(p.plant[i], p.plant[j])
}

var plants = []plant{
	{"flower", 0.45, 0.8},
	{"tree", 125, 5.8},
	{"green", 0.3, 0.4},
	{"tian", 138.7, 4.8},
}

//针对一个结构体切片按照结构体的不同字段进行排序
func main() {

	name := func(p1, p2 plant) bool {
		return p1.name < p2.name
	}

	weight := func(p1, p2 plant) bool {
		return p1.w < p2.w
	}

	height := func(p1, p2 plant) bool {
		return p1.h < p2.h
	}

	unheight := func(p1, p2 plant) bool {
		return !(p1.h < p2.h)
	}

	By(name).sort(plants)
	fmt.Println(plants)
	By(weight).sort(plants)
	fmt.Println(plants)
	By(height).sort(plants)
	fmt.Println(plants)
	By(unheight).sort(plants)
	fmt.Println(plants)
}
