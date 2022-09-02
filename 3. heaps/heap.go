package main

import (
	"container/heap"
	"fmt"
)

func main() {

	var intHeap = &IntegerHeap{1, 4, 5}

	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}

//integerHeap a type
type IntegerHeap []int

//integerHeap method - gets the length of integerHeap
func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

//integerHeap method - checks if element of i index is less than j index
func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

//integerHeap method - swaps the element of i to j index
func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

//integerHeap method - pushed the item
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

//integerHeap method - pops the item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}
