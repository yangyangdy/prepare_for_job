package my_obsatcle_question

type trie struct {
	children [26]*trie
	isEnd bool
}

func (t *trie) Insert(s string) {
	node := t
	for _, ch := range s {
		if node.children[ch-'a'] == nil {
			node.children[ch-'a'] = &trie{}
		}
		node = node.children[ch-'a']
	}
	node.isEnd = true
}

func (t *trie) SearchPrefix(prefix string) *trie {
	node := t
	for _, ch := range prefix { // 进行一层一层地查找
		if node.children[ch-'a'] == nil {
			return nil
		}
		node = node.children[ch-'a']
	}
	return node //返回匹配的最后一个节点
}

func (t *trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

//是否有prefix的前缀
func (t *trie) StartWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}