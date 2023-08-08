package main

import "fmt"

// Do not edit the class below except for the buildHeap,
// siftDown, siftUp, peek, remove, and insert methods.
// Feel free to add new properties and methods to the class.
type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	// Do not edit the lines below.
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
	lastParent := (len(array) - 2) / 2
	for i := lastParent + 1; i >= 0; i-- {
		h.siftDown(i, len(array)-1)
	}
}

func (h *MinHeap) siftDown(currentIndex, endIndex int) {
	leftChildIdx := (currentIndex * 2) + 1
	for leftChildIdx <= endIndex {
		rightChildIdx := -1
		swapIdx := leftChildIdx
		if leftChildIdx+1 <= endIndex {
			rightChildIdx = leftChildIdx + 1
			if (*h)[rightChildIdx] < (*h)[leftChildIdx] {
				swapIdx = rightChildIdx
			}
		}
		if (*h)[currentIndex] > (*h)[swapIdx] {
			(*h)[currentIndex], (*h)[swapIdx] = (*h)[swapIdx], (*h)[currentIndex]
			currentIndex = swapIdx
			leftChildIdx = (currentIndex * 2) + 1
		} else {
			return
		}
	}
}

func (h *MinHeap) siftUp() {
	currentIdx := len(*h) - 1
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && (*h)[currentIdx] < (*h)[parentIdx] {
		(*h)[currentIdx], (*h)[parentIdx] = (*h)[parentIdx], (*h)[currentIdx]
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}

func (h MinHeap) Peek() int {
	if len(h) == 0 {
		return -1
	}
	return h[0]
}

func (h *MinHeap) Remove() int {
	lastIdx := len(*h) - 1
	(*h)[0], (*h)[lastIdx] = (*h)[lastIdx], (*h)[0]
	poppedVal := (*h)[lastIdx]
	(*h) = (*h)[:lastIdx]
	h.siftDown(0, lastIdx-1)
	return poppedVal
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}

func main() {
	var minHeap = NewMinHeap([]int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41})
	fmt.Println(minHeap)
	minHeap.Insert(76)
	fmt.Println(minHeap)
	minHeap.Insert(87)
	fmt.Println(minHeap)
}
