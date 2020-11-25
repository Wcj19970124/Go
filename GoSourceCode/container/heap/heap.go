package main

import (
	"container/heap"
	"fmt"
)

//IntHeap 作为堆的整型数组
type IntHeap []int

//Len 实现接口的Len方法
func (h IntHeap) Len() int {
	return len(h)
}

//Less 实现接口的Less方法，决定是大顶堆还是小顶堆或者优先级队列
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

//Swap 实现接口的元素交换方法
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

//Push 实现接口的放入元素方法
func (h *IntHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
}

//Pop 实现堆的出堆操作
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//堆操作（小顶堆和大顶堆）
//当然也可以用其实现优先级队列，可在官网看到示例
//这里实现的是小顶堆，需要注意出堆和入堆的操作都需要使用heap.Push()和heap.Pop()
func main() {

	nums := &IntHeap{8, 6, 7, 3, 4, 1, 2, 9}
	heap.Init(nums)
	heap.Push(nums, 5)
	fmt.Println("最小值为:", (*nums)[0])
	for nums.Len() > 0 {
		fmt.Printf("%d", heap.Pop(nums))
	}
}
