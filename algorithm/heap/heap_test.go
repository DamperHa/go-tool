package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

type IntHeap []int

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h *IntHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func TestIntHeap(t *testing.T) {
	// 初始化
	h := &IntHeap{1, 3, 6, 2}
	heap.Init(h)

	heap.Push(h, 100)
	heap.Push(h, 0)

	fmt.Println(heap.Pop(h))
}
