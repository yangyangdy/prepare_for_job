package main

import "debug/elf"

//前缀树的结构
type Trie struct {
	children [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func(t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

//拓扑排序，有序数组
func findOrder(numCourse int, prerequisites [][]int) []int {
	var (
		edges = make([][]int, numCourse)
		visited = make([]int, numCourse)
		result []int
		valid bool = true //判断是否有环
		dfs func(u int)
	)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0{
				dfs(v)
				if !valid{
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	//得到每个课程的指向的下一个课程列表
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i<numCourse && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	for i:=0;i<len(result)/2;i++{
		result[i], result[numCourse-i-1] = result[numCourse-i-1], result[i]
	}
	return result
}


//拓扑排序，用广度优先的思路去做
//1. 需要一个入度表， 记录每个节点的有几个前驱
//2. 需要一个队列，入度为0的节点入队，队列里的节点就是可以排课的课程
//3. 需要一个出边表，处理完该节点后处理它的出边，出边的处理就是给出边的入度减1.之后等于0就加入队列处理
//4. 课程的依赖关系可能出现环状，需要加入检查
func findOrder2(numCourse int, prerequisites [][]int) []int{
	//入度表，记录本课程先修课程有几门
	inDegree := make([]int, numCourse)
	//出边表，记录本课程是哪些课程的先修
	outEdge := make([][]int, numCourse)
	//结果集
	res := []int{}

	for i:=0;i<len(prerequisites);i++ {
		//更新入度
		inDegree[prerequisites[i][0]]++
		//更新出边
		outEdge[prerequisites[i][1]] = append(outEdge[prerequisites[i][1]], prerequisites[i][0])
	}
	//队列用于记录入度为0的课程
	queue := []int{}

	for i:=0;i<numCourse;i++{
		if inDegree[i] == 0{
			queue = append(queue, i)
		}
	}
	//没有不依赖其他课程的课程，返回空
	if len(queue) == 0 {
		return res
	}

	for len(queue) > 0{
		//取出队头，加入结果集
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)

		//处理队头的出边，把入度为0的入队
		for i:=0;i<len(outEdge[node]);i++{
			out := outEdge[node][i]
			inDegree[out]-- //减去
			if inDegree[out] == 0{
				queue = append(queue, out)
			}
		}
	}
	//检查是否所有的课程都被安排，如果没有，说明有环，无法安排，返回空
	if len(res) == numCourse {
		return res
	}
	return []int{}
}

class Solution {
public int rob(int[] nums) {
if (nums == null || nums.length == 0) {
return 0;
}
int length = nums.length;
if (length == 1) {
return nums[0];
}
int[] dp = new int[length];
dp[0] = nums[0];
dp[1] = Math.max(nums[0], nums[1]);
for (int i = 2; i < length; i++) {
dp[i] = Math.max(dp[i - 2] + nums[i], dp[i - 1]);
}
return dp[length - 1];
}
}
