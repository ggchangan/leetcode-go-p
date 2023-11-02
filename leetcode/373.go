package leetcode

var KSmallestPairs = kSmallestPairs

func kSmallestPairs(nums1 []int, nums2 []int, k int) (ans [][]int) {
	fL := len(nums1)
	sL := len(nums2)

	//small nums with big k
	if fL*sL <= k {
		for i := 0; i < fL; i++ {
			for j := 0; j < sL; j++ {
				ans = append(ans, []int{nums1[i], nums2[j]})
			}
		}
		return
	}

	//初始化为min(k, fL)中所有元素和nums2[0]构成的heap
	init := func() (elems []Elem) {
		for i := 0; i < fL && i < k; i++ {
			e := Elem{
				Num1:   nums1[i],
				Num2:   nums2[0],
				Index2: 0,
			}
			elems = append(elems, e)
		}
		return
	}
	elems := init()
	buildMin(elems)

	for i := 0; i < k && len(elems) > 0; i++ {
		e := elems[0]
		ans = append(ans, []int{e.Num1, e.Num2})
		if e.Index2 == sL-1 {
			elems = elems[1:]
		} else {
			elems[0] = Elem{
				Num1:   e.Num1,
				Num2:   nums2[e.Index2+1],
				Index2: e.Index2 + 1,
			}
		}
		if len(elems) > 1 {
			heapifyElem(elems, len(elems), 0)
		}
	}

	return
}

type Elem struct {
	Num1   int
	Num2   int
	Index2 int
}

func heapifyElem(arr []Elem, n, i int) {
	smaller := func(e1, e2 Elem) bool {
		return e1.Num1+e1.Num2 < e2.Num1+e2.Num2
	}
	smallest := i
	lson, rson := 2*i+1, 2*i+2
	if lson < n && smaller(arr[lson], arr[smallest]) {
		smallest = lson
	}
	if rson < n && smaller(arr[rson], arr[smallest]) {
		smallest = rson
	}

	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
		heapifyElem(arr, n, smallest)
	}
}

func buildMin(arr []Elem) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyElem(arr, i, 0)
	}
}
