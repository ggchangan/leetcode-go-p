package leetcode

var Construct = construct

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var helper func(r1, r2, c1, c2 int) *Node

	helper = func(r1, r2, c1, c2 int) *Node {
		for i := r1; i <= r2; i++ {
			for j := c1; j <= c2; j++ {
				if grid[i][j] != grid[r1][c1] {
					rMid := (r1 + r2) / 2
					cMid := (c1 + c2) / 2
					return &Node{
						Val:         false,
						IsLeaf:      false,
						TopLeft:     helper(r1, rMid, c1, cMid),
						TopRight:    helper(r1, rMid, cMid+1, c2),
						BottomLeft:  helper(rMid+1, r2, c1, cMid),
						BottomRight: helper(rMid+1, r2, cMid+1, c2),
					}
				}
			}
		}

		return &Node{
			IsLeaf: true,
			Val:    grid[r1][c1] == 1,
		}
	}

	return helper(0, len(grid)-1, 0, len(grid)-1)
}
