package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
//螺旋遍历矩阵
func spiralOrder(matrix [][]int) []int{
	m := len(matrix) //行数
	if m == 0 {
		return nil
	}
	n := len(matrix[0]) //列数
	if n == 0 {
		return nil
	}
	// 上下左右的下标
	top := 0
	left := 0
	bottom := m-1
	right := n-1
	count := 0
	sum := m*n
	res := []int{}
	//外层循环每次遍历一圈
	for count < sum {
		i, j := top, left
		for j<=right && count < sum {
			res = append(res, matrix[i][j])
			count++
			j++
		}
		i, j = top + 1, right
		for i<=bottom && count < sum {
			res = append(res, matrix[i][j])
			count++
			i++
		}
		i,j = bottom, right-1
		for j>=left && count < sum {
			res = append(res, matrix[i][j])
			count++
			j--
		}
		i,j=bottom-1, left
		for i>top && count < sum {
			res = append(res, matrix[i][j])
			count++
			i--
		}
		//进入下一层
		top,left,bottom,right = top+1,left+1,bottom-1,right-1
	}
	return res
}

var BCR_TO_DISCOUNT_MAP = map[float64]float64{
	1000:0.7,
	1001:0.75,
	1002:0.8,
	1003:0.85,
	1004:0.9,
	1005:0.95,
	1:1.0,
}

func main(){
	fx := BCR_TO_DISCOUNT_MAP
	fmt.Println(fx)
	var ff float64 = 4.294967295e+09
	fmt.Println(ff)
	r := math.Ceil(1.05)
	fmt.Println(r)
	path := "/ade/denfe/../defe"
	arr := strings.Split(path, "/")
	for _,v := range arr{
		println("val:", v)
	}
	a := [...][]int{
		{1,2},
		{3,4},
		{5,6},
		{7,8},
	}
	b := a[2]
	b[0] = 30
	b[1] = 40
	fmt.Println(b)
	fmt.Println(a)
	var s = []string{"11", "22", "33"}
	ret := strings.Join(s, "|")
	fmt.Println(ret)
}



func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}



func addBinary(a string, b string) string{
	ans :=""
	carry :=0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i:=0;i<n;i++{
		if i<lenA{
			carry += int(a[lenA-i-1]-'0')
		}
		if i<lenB{
			carry+=int(b[lenB-i-1]-'0')
		}
		ans = strconv.Itoa(carry%2)+ans
		carry /= 2
	}
	if carry > 0{
		ans = "1" + ans
	}
	return ans
}

func max(a,b int) int {
	if a>b{
		return a
	}
	return b
}