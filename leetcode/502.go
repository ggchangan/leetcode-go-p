package leetcode

import (
	"container/heap"
	"sort"
)

var FindMaximizedCapital = findMaximizedCapital

type pair struct{ p, c int }

type pairHeap []pair

func (h pairHeap) Len() int { return len(h) }
func (h pairHeap) Less(i, j int) bool {
	return h[i].p > h[j].p
}
func (h pairHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(x any) {
	*h = append(*h, x.(pair))
}
func (h *pairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	cL := len(capital)
	data := make([]pair, cL)
	for i := 0; i < len(capital); i++ {
		data[i] = pair{p: profits[i], c: capital[i]}
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].c < data[j].c
	})
	if w < data[0].c {
		return w
	}

	h := &pairHeap{}
	heap.Init(h)
	i := 0 //current min capital
	dL := len(data)
	cnt := 0
	for cnt < k {
		for i < dL && data[i].c <= w {
			heap.Push(h, data[i])
			i++
		}
		if h.Len() == 0 {
			break
		}
		p := heap.Pop(h).(pair)
		w += p.p
		cnt++
	}

	return w
}
