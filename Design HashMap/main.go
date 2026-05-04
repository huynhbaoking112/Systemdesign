package main

type MyHashMap struct {
	counter int
	HashMap [][][]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		counter: 769,
		HashMap: make([][][]int, 769),
	}
}

func getIndex(counter int, val int) int {
	return val % counter
}

func isContain(hashMap *MyHashMap, key int, i int) (bool, int) {
	if len(hashMap.HashMap[i]) > 0 {
		for index, val := range hashMap.HashMap[i] {
			if val[0] == key {
				return true, index
			}
		}
	}
	return false, -1
}

func (this *MyHashMap) Put(key int, value int) {
	i := getIndex(this.counter, key)
	isContain, index := isContain(this, key, i)
	if !isContain {
		this.HashMap[i] = append(this.HashMap[i], []int{key, value})
		return
	}
	this.HashMap[i][index][1] = value
}

func (this *MyHashMap) Get(key int) int {
	i := getIndex(this.counter, key)
	isContain, index := isContain(this, key, i)
	if isContain {
		return this.HashMap[i][index][1]
	}
	return -1

}

func (this *MyHashMap) Remove(key int) {
	i := getIndex(this.counter, key)
	isContain, index := isContain(this, key, i)
	if isContain {
		this.HashMap[i][index] = this.HashMap[i][len(this.HashMap[i])-1]
		this.HashMap[i] = this.HashMap[i][:len(this.HashMap[i])-1]
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
