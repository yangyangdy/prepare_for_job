package main

type unionFind struct {
	parent []int
	rank []int
	count int
}

func (uf *unionFind) init(n int) {
	uf.count = n
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
	}
}

func (uf *unionFind) find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root] //找根节点
	}
	//compress path
	for
}
func main() {

}
