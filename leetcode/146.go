package leetcode

type LRUNode struct {
	Key  int
	Val  int
	Next *LRUNode
}

type LRUCache struct {
	Capacity int
	Size     int
	Head     *LRUNode
	Tail     *LRUNode
	m        map[int]*LRUNode
}

func Constructor(capacity int) LRUCache {
	t := &LRUNode{}
	return LRUCache{
		Capacity: capacity,
		Head:     t,
		Tail:     t,
		m:        make(map[int]*LRUNode),
	}
}

func (this *LRUCache) Get(key int) int {
	preQ := this.m[key]
	if preQ == nil {
		return -1
	}
	//更新
	q := preQ.Next
	if q != this.Tail {
		this.m[q.Key] = this.Tail
		this.Tail.Next = q
		this.Tail = q
		preQ.Next = q.Next
		if q.Next != nil {
			this.m[q.Next.Key] = preQ
		}
	}

	return q.Val
}

func (this *LRUCache) Put(key int, value int) {
	preQ := this.m[key]
	//不存在
	if preQ == nil {
		t := &LRUNode{
			Key: key,
			Val: value,
		}
		this.m[key] = this.Tail
		this.Tail.Next = t
		this.Tail = this.Tail.Next
		if this.Size < this.Capacity {
			this.Size++
		} else {
			//删除
			q := this.Head.Next
			if q == this.Tail {
				this.Tail = this.Head
			} else {
				this.Head.Next = q.Next
				if q.Next != nil {
					this.m[q.Next.Key] = this.Head
				}
			}
			delete(this.m, q.Key)
		}
	} else {
		q := preQ.Next
		q.Val = value
		if q != this.Tail {
			this.m[q.Key] = this.Tail
			this.Tail.Next = q
			this.Tail = q
			preQ.Next = q.Next
			if q.Next != nil {
				this.m[q.Next.Key] = preQ
			}
		}
	}
}
