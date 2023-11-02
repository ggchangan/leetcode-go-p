package leetcode

type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

func ConstructorTrie() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	p := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if p.Children[c] == nil {
			p.Children[c] = &Trie{}
		}
		p = p.Children[c]
	}
	p.IsEnd = true
}
func (this *Trie) SearchPrefix(word string) *Trie {
	p := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		p = p.Children[c]
		if p == nil {
			return nil
		}
	}
	return p
}

func (this *Trie) Search(word string) bool {
	p := this.SearchPrefix(word)
	return p != nil && p.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}
