package leetcode

var FindWords = findWords

func findWords(board [][]byte, words []string) (ans []string) {
	m := len(board)
	n := len(board[0])
	var dfs func(i, j int, p *WordTrie)
	dfs = func(i, j int, p *WordTrie) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == '#' {
			return
		}

		c := board[i][j]

		p = p.Children[c-'a']
		if p == nil {
			return
		}

		board[i][j] = '#'
		if len(p.Word) > 0 {
			ans = append(ans, p.Word)
			p.Word = ""
		}

		dfs(i-1, j, p)
		dfs(i+1, j, p)
		dfs(i, j-1, p)
		dfs(i, j+1, p)

		board[i][j] = c

	}

	tree := buildTrie(words)

	//dfs(0, 0, tree)
	//dfs(1, 3, tree)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, tree)
		}
	}

	return
}

type WordTrie struct {
	Children [26]*WordTrie
	Word     string
}

func buildTrie(words []string) (trie *WordTrie) {
	trie = &WordTrie{}
	for i := 0; i < len(words); i++ {
		p := trie
		for j := 0; j < len(words[i]); j++ {
			c := words[i][j] - 'a'
			if p.Children[c] == nil {
				p.Children[c] = &WordTrie{}
			}
			p = p.Children[c]
		}
		p.Word = words[i]
	}

	return
}
