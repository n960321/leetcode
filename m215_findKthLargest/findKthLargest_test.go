package m215_findKthLargest

import (
	"container/heap"
	"sort"
	"testing"
)

/*
https://leetcode.com/problems/kth-largest-element-in-an-array/

215. Kth Largest Element in an Array
Medium

4374

288

Add to List

Share
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

// PriorityQueue reference: https://www.cnblogs.com/huxianglin/p/6925119.html

func findKthLargest(nums []int, k int) int {
	minHeap := MinHeap(nums)
	heap.Init(&minHeap)
	for minHeap.Len() > k {
		heap.Pop(&minHeap)
	}
	return heap.Pop(&minHeap).(int)
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// T: O(log n)
// S: O(log n)
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// T: O(log n) - because when we remove the root (min) we need to reorder.
// S: O(log n)
func (h *MinHeap) Pop() interface{} {
	n := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return n
}

func findKthLargest_v1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
