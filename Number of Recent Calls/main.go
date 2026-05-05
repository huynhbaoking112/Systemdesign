package main

type RecentCounter struct {
	arr   []int
	index int
}

func Constructor() RecentCounter {
	return RecentCounter{
		arr:   []int{},
		index: 0,
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.arr = append(this.arr, t)
	for this.arr[this.index] < t-3000 {
		this.index++
	}
	return len(this.arr) - this.index
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
