package main

type MyQueue struct {
	q1 []int
	q2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		q1: []int{},
		q2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.q1 = append(this.q1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.q2) == 0 {
		for i := len(this.q1) - 1; i >= 0; i-- {
			this.q2 = append(this.q2, this.q1[i])
		}
		this.q1 = []int{}
	}
	a := this.q2[len(this.q2)-1]
	this.q2 = this.q2[:len(this.q2)-1]
	return a

}

func (this *MyQueue) Peek() int {
	if len(this.q2) == 0 {
		for i := len(this.q1) - 1; i >= 0; i-- {
			this.q2 = append(this.q2, this.q1[i])
		}
		this.q1 = []int{}
	}
	a := this.q2[len(this.q2)-1]
	return a
}

func (this *MyQueue) Empty() bool {
	return len(this.q1)+len(this.q2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
