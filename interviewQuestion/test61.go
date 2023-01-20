package main

func openLock3(deadends []string, target string) int {
	step := 0
	deadendsMap := make(map[string]bool)
	visitedMap := make(map[string]bool)

	for _, v := range deadends {
		deadendsMap[v] = true
	}

	get := func(status string) (ret []string) {
		s := []byte(status)
		for i, b:=range s {
			s[i] =b-1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))
			s[i] = b+1
			if s[i]>'9'{
				s[i]='0'
			}
			ret = append(ret, string(s))
			s[i] = b
		}
		return
	}

	q := []string{"0000"}
	//循环直至队列为空
	for len(q) > 0 {
		size := len(q)
		for i:=0;i<size;i++{
			//遍历当前层的节点
			node := q[0]//获取出列的节点
			q = q[1:]//节点出列
			if node == target {
				return step
			}
			if _, ok := visitedMap[node]; ok {
				//已经访问过的节点就跳过
				continue
			}
			if _, ok := deadendsMap[node]; ok {
				continue
			}
			visitedMap[node] = true
			for _, nxt := range get(node) {
				if nxt == target {
					return step
				}
				q= append(q, nxt)
			}
			step++
		}
	}
	return 0
}

func main() {

}
