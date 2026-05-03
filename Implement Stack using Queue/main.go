package main

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{
		q: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
	l := len(this.q)
	if l == 1 {
		return
	}
	for i := 0; i < l-1; i++ {
		a := this.Pop()
		this.q = append(this.q, a)
	}
	return
}

func (this *MyStack) Pop() int {
	a := this.q[0]
	this.q = this.q[1:]
	return a
}

func (this *MyStack) Top() int {
	a := this.q[0]
	return a
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {

}
