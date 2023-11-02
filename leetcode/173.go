package leetcode

type BSTIterator struct {
	Stack []*TreeNode
	Cur   *TreeNode
}

func ConstructorBSTIterator(root *TreeNode) BSTIterator {
	return BSTIterator{
		Stack: []*TreeNode{},
		Cur:   root,
	}
}

func (this *BSTIterator) Next() int {
	for this.Cur != nil {
		this.Stack = append(this.Stack, this.Cur)
		this.Cur = this.Cur.Left
	}

	l := len(this.Stack)
	root := this.Stack[l-1]
	this.Stack = this.Stack[:l-1]
	this.Cur = root.Right

	return root.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) > 0 || this.Cur != nil
}
