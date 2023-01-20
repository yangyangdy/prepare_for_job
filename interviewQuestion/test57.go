package main

import (
	"fmt"
	"sort"
)

func trailingZeroes(n int) int {
	res :=0
	divisor := 5
	for divisor <= n {
		res += n/divisor
		divisor += 5
	}
	return res
}

type testA struct {

}
func main() {
	mp := make(map[int]*testA)
	mp[1] = &testA{}

	//mp := make(map[int]int,0)
	//if v, ok := mp[1]; ok {
	//	fmt.Println(v)
	//}else {
	//	fmt.Println("yes")
	//}
	//arr := []int{7,2,5,3,1}
	//sort.Slice(arr, func(i, j int) bool {
	//	return arr[i] > arr[j]
	//})
	//sort.Ints()
	//fmt.Println(arr)
	//s := []string{"2022-10-2", "2022-10-7"}
	//err := fmt.Errorf("str is arr=%v", s)
	//fmt.Println(err.Error())
	//arr := []int{1,2,3,4}
	//arr = arr[:len(arr)-1]
	//fmt.Println(arr)
	//res := trailingZeroes(125)
	//fmt.Println(res)
	//nums1 := []int{3,2,6,4}
	//nums2 := []int{2,1,3,7}
	//res := advantageCount(nums1, nums2)
	//fmt.Println(res)
}


func advantageCount(nums1, nums2 []int) []int {
	n := len(nums1)
	ans := make([]int, n)
	sort.Ints(nums1)
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	fmt.Println(ids)
	//升序排序,nums2的升序序号在ids中
	sort.Slice(ids, func(i, j int) bool { return nums2[ids[i]] < nums2[ids[j]] })
	fmt.Println(ids)
	left, right := 0, n-1
	for _, x := range nums1 {
		if x > nums2[ids[left]] {
			ans[ids[left]] = x // 用下等马比下等马
			left++
		} else {
			ans[ids[right]] = x // 用下等马比上等马
			right--
		}
	}
	return ans
}


