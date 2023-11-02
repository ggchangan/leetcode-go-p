package leetcode

func minMutation(startGene string, endGene string, bank []string) int {
	bankMap := make(map[string]bool, 0)
	for i := 0; i < len(bank); i++ {
		bankMap[bank[i]] = true
	}

	if _, ok := bankMap[endGene]; !ok {
		return -1
	}

	if startGene == endGene {
		return 0
	}

	next := func(gene string) (newGenes []string) {
		for i := 0; i < len(gene); i++ {
			for _, c := range []byte{'A', 'C', 'G', 'T'} {
				if c != gene[i] {
					newGene := gene[:i] + string(c) + gene[i+1:]
					if _, ok := bankMap[newGene]; ok {
						newGenes = append(newGenes, newGene)
					}
				}
			}
		}
		return
	}

	var geneChanges = make(map[string][]string, 0)
	for _, gene := range append(bank, startGene) {
		geneChanges[gene] = next(gene)
	}

	cur := []string{startGene}
	change := 0
	for len(cur) > 0 {
		nextLevel := make([]string, 0)
		for i := 0; i < len(cur); i++ {
			if cur[i] == endGene {
				return change
			}
			nextLevel = append(nextLevel, geneChanges[cur[i]]...)
			delete(geneChanges, cur[i])
		}
		change += 1
		cur = nextLevel
	}

	return -1
}
