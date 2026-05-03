package MinHeap

import (
	"errors"
)

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: []int{},
	}
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func (h *MinHeap) Push(x int) {
	h.data = append(h.data, x)
	index := h.Len() - 1
	for index > 0 {
		parent_index := (index - 1) / 2
		if h.data[parent_index] <= h.data[index] {
			break
		}
		h.data[parent_index], h.data[index] = h.data[index], h.data[parent_index]
		index = parent_index
	}
}

func (h *MinHeap) Peek() (int, error) {
	if h.Len() == 0 {
		return 0, errors.New("Heap is empty")
	}
	return h.data[0], nil
}

func (h *MinHeap) Pop() (int, error) {
	if h.Len() == 0 {
		return 0, errors.New("Heap is empty")
	}
	result := h.data[0]
	h.data[0] = h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	index := 0
	if h.Len() > 0 {
		for {
			right := index*2 + 2
			left := index*2 + 1
			min_result := index
			if right < h.Len() && h.data[right] < h.data[min_result] {
				min_result = right
			}

			if left < h.Len() && h.data[left] < h.data[min_result] {
				min_result = left
			}

			if min_result == index {
				break
			}
			h.data[min_result], h.data[index] = h.data[index], h.data[min_result]
			index = min_result
		}
	}

	return result, nil
}
