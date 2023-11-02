package leetcode

type WordDictionary struct {
	Children [26]*WordDictionary
	IsEnd    bool
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	p := this
	for _, c := range word {
		i := c - 'a'
		if p.Children[i] == nil {
			p.Children[i] = &WordDictionary{}
		}
		p = p.Children[i]
	}
	p.IsEnd = true
}

// dfs
func (this *WordDictionary) Search(word string) bool {
	var helper func(level int, p *WordDictionary) bool
	helper = func(level int, p *WordDictionary) bool {
		if p == nil {
			return false
		}

		if level == len(word) {
			return p.IsEnd
		}

		c := word[level]
		if c != '.' {
			p = p.Children[c-'a']
			return helper(level+1, p)
		}

		for _, child := range p.Children {
			if child != nil {
				if helper(level+1, child) {
					return true
				}
			}
		}
		return false
	}

	return helper(0, this)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
