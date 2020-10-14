package m347_topKFrequent

import (
	"fmt"
	"sort"
	"testing"
)

func topKFrequent_v1(nums []int, k int) []int {
	r := []int{}
	m := map[int]int{}
	for _, v := range nums {
		if _, exist := m[v]; !exist {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	fmt.Printf("============= map %+v \n", m)

	for key, v := range m {
		fmt.Printf("========= key: %v , val: %v\n", key, v)
		if len(r) < k {
			r = append(r, key)
			fmt.Printf("===== %+v\n", r)
			continue
		}

		for i := range r {
			if m[r[i]] < v {
				r = append(r[:i], r[i+1:]...)
				fmt.Printf("------ %+v\n", r)
				r = append(r, key)
				fmt.Printf("++++++ %+v\n", r)
				break
			}
		}
	}
	return r
}

type Info struct {
	Key int
	Num int
}

func topKFrequent(nums []int, k int) []int {
	l := []*Info{}
	m := map[int]*Info{}
	for _, v := range nums {
		if _, exist := m[v]; !exist {
			m[v] = &Info{Key: v, Num: 1}
			l = append(l, m[v])
		} else {
			m[v].Num++
		}
	}
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Num > l[j].Num
	})
	r := []int{}

	for i := 0; i < k && i < len(l); i++ {
		r = append(r, l[i].Key)
	}
	return r
}

func Test_topKFrequent(t *testing.T) {
	// fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3, 3, 3, 3}, 2))
	// fmt.Println(topKFrequent([]int{1}, 1))
	// fmt.Println(topKFrequent([]int{3, 0, 1, 0}, 1))
	// fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	// fmt.Println(topKFrequent([]int{2, 3, 4, 1, 4, 0, 4, -1, -2, -1}, 2))
	// fmt.Println(topKFrequent([]int{-1, 1, 4, -4, 3, 5, 4, -2, 3, -1}, 3)) // -1,3,4

	fmt.Println(topKFrequent([]int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}, 3)) // 3,7,10
}
