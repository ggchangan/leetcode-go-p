package heap

import (
	"container/heap"
)

type minHp struct{ data []int }

func (h minHp) Len() int           { return len(h.data) }
func (h minHp) Less(i, j int) bool { return h.data[i] < h.data[j] }
func (h minHp) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *minHp) Push(x any) {
	h.data = append(h.data, x.(int))
}

func (h *minHp) Pop() any {
	a := h.data
	v := a[len(a)-1]
	h.data = a[:len(a)-1]
	return v
}

func (h minHp) Top() any { return h.data[0] }

type maxHp struct{ minHp }

// 重写排序
func (h maxHp) Less(i, j int) bool { return h.data[i] > h.data[j] }

type MedianFinder struct {
	leftMax  maxHp //<= median, 最大堆, 左堆
	rightMin minHp //>median, 最小堆，右堆
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) { //放入适当的堆，并维持两个堆数据量平衡
	if this.leftMax.Len() == 0 || num <= this.leftMax.Top().(int) { //放入左堆
		heap.Push(&this.leftMax, num)
		if this.leftMax.Len() > this.rightMin.Len()+1 { //发生倾斜
			heap.Push(&this.rightMin, heap.Pop(&this.leftMax))
		}
	} else { //放入右堆
		heap.Push(&this.rightMin, num)
		if this.rightMin.Len() > this.leftMax.Len() { //发生倾斜
			heap.Push(&this.leftMax, heap.Pop(&this.rightMin))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	leftL := this.leftMax.Len()
	rightL := this.rightMin.Len()
	if leftL > rightL {
		return float64(this.leftMax.Top().(int))
	}

	return float64(this.leftMax.Top().(int)+this.rightMin.Top().(int)) / 2
}
