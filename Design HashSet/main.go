package main

type MyHashSet struct {
	bucketCount int
	bucketArray [][]int
}

func Constructor() MyHashSet {
	return MyHashSet{
		bucketCount: 769,
		bucketArray: make([][]int, 769),
	}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	bucket := key % this.bucketCount
	this.bucketArray[bucket] = append(this.bucketArray[bucket], key)
}

func (this *MyHashSet) Remove(key int) {
	bucket := key % this.bucketCount
	for i, v := range this.bucketArray[bucket] {
		if v == key {
			this.bucketArray[bucket][i] = this.bucketArray[bucket][len(this.bucketArray[bucket])-1]
			this.bucketArray[bucket] = this.bucketArray[bucket][:len(this.bucketArray[bucket])-1]
			// this.bucketArray[bucket] = append(this.bucketArray[bucket][:i], this.bucketArray[bucket][i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	bucket := key % this.bucketCount
	for _, v := range this.bucketArray[bucket] {
		if v == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
