package main

import "sort"

func longestCommonPrefix(strs []string) string {
	//找出最小的且相同的前缀
	minLen := len(strs[0])
	for i :=1 ;i<len(strs);i++{
		minLen = min(minLen, len(strs[i]))
	}
	prefix := strs[0]
	isEnd := false
	for i := 1; i<len(strs);i++{
		for j:=0;j<minLen;j++{
			if strs[i][j] != prefix[j]{
				prefix = prefix[0:j]
				isEnd
				break
			}
		}
	}
	return prefix
}

func main() {
	a := []string{"dog","racecar","car"}
	longestCommonPrefix(a)
}

func min(a,b int)int {
	if a <=b {
		return a
	}
	return b
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	n, SumClosest := len(nums), nums[0] + nums[1] + nums[2] // 初始化为前三元素的值，避免任何个

	for i:=0; i<n-2; i++{

		if i>0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, n-1
		for j<k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target{
				return target
			}
			if abs(sum-target) < abs(SumClosest-target){
				SumClosest = sum
			}
			if sum>target{
				k--
			}else{
				j++
			}
		}
	}

	return SumClosest
}

func abs(x int) int{
	if x<0{
		return -x
	}
	return x
}

