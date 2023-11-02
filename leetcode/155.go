package leetcode

type MinStack struct {
	dataStack []int
	minStack  []int
}

func ConstructorMinStack() MinStack {
	return MinStack{
		dataStack: make([]int, 0),
		minStack:  make([]int, 0),
	}

}

func (this *MinStack) Push(val int) {
	this.dataStack = append(this.dataStack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		min := this.minStack[len(this.minStack)-1]
		if val > min {
			this.minStack = append(this.minStack, min)
		} else {
			this.minStack = append(this.minStack, val)
		}
	}

}

func (this *MinStack) Pop() {
	this.dataStack = this.dataStack[:len(this.dataStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.dataStack[len(this.dataStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
