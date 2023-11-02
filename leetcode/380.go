package leetcode

import "math/rand"

type RandomizedSet struct {
	keys map[int]int
	data []int
}

func ConstructRandomizedSet() RandomizedSet {
	return RandomizedSet{keys: map[int]int{}, data: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.keys[val]; !ok {
		this.data = append(this.data, val)
		this.keys[val] = len(this.data) - 1
		return true
	}

	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.keys[val]; !ok {
		return false
	}

	last := len(this.data) - 1
	this.data[this.keys[val]] = this.data[last]
	this.keys[this.data[last]] = this.keys[val]
	this.data = this.data[:last] //注意最后做删除操作
	delete(this.keys, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[rand.Intn(len(this.data))]
}
