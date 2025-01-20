package main

import "container/heap"

func main() {

}

/*
*
215. 数组中的第K个最大元素 https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4
*/

func findKthLargest(nums []int, k int) int {
	// 小顶堆，堆顶是最小元素
	pq := &MinHeap{}
	heap.Init(pq)
	for _, e := range nums {
		// 每个元素都要过一遍二叉堆
		heap.Push(pq, e)
		// 堆中元素多于 k 个时，删除堆顶元素
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	// pq 中剩下的是 nums 中 k 个最大元素，
	// 堆顶是最小的那个，即第 k 个最大元素
	return heap.Pop(pq).(int)
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
