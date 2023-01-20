package main

import "strings"

func main() {

}

//题目所述的滑动谜题，状态1变成状态2后，反向操作可从状态2变回状态1，故初始状态通过操作变成[[1,2,3],[4,5,0]]，与[[1,2,3],[4,5,0]]通过操作变成初始状态所需的最少移动次数是相同的。
//
//谜板只有6个位置，根据排列组合，实际可能出现的状态不会超过 6!=720，故把全部状态的所需的移动次数缓存起来是可行的。
//
//设置 [[1,2,3],[4,5,0]] 作为初始值，通过广度优先搜索（BFS），计算出初始值到达其他状态所需的最少移动次数，并用哈希表保存。
//
//搜索时只需从哈希表中取值，不存在则为无解，时间复杂度：O(1)。
//
//实际运行发现：
//
//有解的状态数量为360。
//如果有解，最多需要21步即可解开谜板。

var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

//和拨号盘的解题思路一致
func slidingPuzzle(board [][]int) int {
	const target = "123450"

	s := make([]byte, 0, 6)
	for _, r := range board {
		for _, v := range r {
			s = append(s, '0'+byte(v))
		}
	}
	start := string(s)
	if start == target {
		return 0
	}

	// 枚举 status 通过一次交换操作得到的状态
	get := func(status string) (ret []string) {
		s := []byte(status)
		x := strings.Index(status, "0")
		for _, y := range neighbors[x] {
			s[x], s[y] = s[y], s[x]
			ret = append(ret, string(s))
			s[x], s[y] = s[y], s[x]
		}
		return
	}

	type pair struct {
		status string
		step   int
	}
	q := []pair{{start, 0}}
	seen := map[string]bool{start: true}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			if !seen[nxt] {
				if nxt == target {
					return p.step + 1
				}
				seen[nxt] = true
				q = append(q, pair{nxt, p.step + 1})
			}
		}
	}
	return -1
}
