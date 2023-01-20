package main

import (
	"fmt"
	"strconv"
)
// 打开转盘所锁,最小的操作次数完成，那么采用穷举中的bfs

func openLock(deadends []string, target string) int {
	const start = "0000"
	if target == start {
		return 0
	}

	dead := map[string]bool{}
	for _, s := range deadends {
		dead[s] = true
	}
	if dead[start] {
		return -1
	}

	// 枚举 status 通过一次旋转得到的数字
	get := func(status string) (ret []string) {
		s := []byte(status)
		for i, b := range s {
			s[i] = b - 1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))
			s[i] = b + 1
			if s[i] > '9' {
				s[i] = '0'
			}
			ret = append(ret, string(s))
			s[i] = b
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
			if !seen[nxt] && !dead[nxt] {
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

func openLock2(deadends []string, target string) int {
	step := 0 // 旋转次数
	deadendsMap := make(map[string]bool)
	visitedMap := make(map[string]bool)

	for _, v := range deadends { // 记录所有“死亡点”
		deadendsMap[v] = true
	}

	q := []string{"0000"} // 队列q
	for len(q) > 0 {      // 循环直至队列为空
		size := len(q)              // 获取BFS当前level的节点个数
		for i := 0; i < size; i++ { // 遍历当前层的节点
			node := q[0]        // 获取出列的节点
			q = q[1:]           // 节点出列
			if node == target { // 如果出列的节点正好是目标节点
				return step // 返回当前所用的步数
			}
			if _, ok := visitedMap[node]; ok { // 之前访问过该节点，跳过
				continue
			}
			if _, ok := deadendsMap[node]; ok { // 遇到“死亡点”，跳过
				continue
			}
			visitedMap[node] = true // 将该点标记为访问过

			for j := 0; j < len(node); j++ { // 通过遍历当前字符串，找出它的所有子节点，安排入列
				num := int(node[j] - '0')                             // 获取当前的数字num
				up := (num + 1) % 10                                  // 往上拧所得的新数，比如1变成2
				down := (num + 9) % 10                                // 往下拧所得的新数，比如7变成6
				q = append(q, node[:j]+strconv.Itoa(up)+node[j+1:])   // 拼成新字符串，入列
				q = append(q, node[:j]+strconv.Itoa(down)+node[j+1:]) // 拼成新字符串 入列
			}
		}
		step++ // 当前层的所有节点遍历完毕，层次+1
	}
	return -1 // 无论如何都遇不到目标节点，返回-1
}


func main(){
	//a := []int{1,2,3,4,5}
	b := []int{6,7,8,9}
	//a = append(a, b...)
	//fmt.Println(a)
	//a[6]=100
	//fmt.Println(a)
	//fmt.Println(b)
	c := append([]int{}, b...)
	c[1]=100
	fmt.Println(b)
	fmt.Println(c)
}