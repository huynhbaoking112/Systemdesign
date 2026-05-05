package main

type OrderedStream struct {
	arr   []string
	index int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		arr:   make([]string, n),
		index: 0,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	result := []string{}
	idKey = idKey - 1
	this.arr[idKey] = value

	for idKey == this.index && this.arr[this.index] != "" {
		result = append(result, this.arr[this.index])
		isLast := this.index == len(this.arr)-1
		if !isLast {
			this.index++
			idKey++
		} else {
			break
		}
	}

	return result
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
