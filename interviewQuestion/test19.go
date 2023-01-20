package main

import "fmt"

type aa struct {
	a int
	b string
}
func main() {
	mp := map[int]*aa{
		1:&aa{1,"1"},
		2:&aa{2,"2"},
		3:&aa{3,"3"},
	}
	fmt.Println(mp[1])
	fmt.Println(mp[10])
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
	mytest1()
	go mytest2()
	mytest3()
	for {
		select{

		}
	}
}

func mytest1() {
	fmt.Println("this is test 1")
}

func mytest2() {
	//defer func(){
	//	if err := recover();err!=nil{
	//		fmt.Println("this is test 2")
	//	}
	//}()
	fmt.Println("this is test 2")
	panic("test 2 is panic")
}

func mytest3() {
	fmt.Println("this is test 3")
}

type Trie struct {
	children [26]*Trie
	isEnd bool
}

func (t *Trie) Insert(s string) {
	node := t
	for _, ch := range s{
		if node.children[ch-'a'] == nil {
			node.children[ch-'a'] = &Trie{}
		}
		node = node.children[ch-'a']
	}
	node.isEnd = true
}

//查找前缀
func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		if node.children[ch-'a'] == nil {
			return nil
		}
		node = node.children[ch-'a']
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func(t *Trie) StartWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

