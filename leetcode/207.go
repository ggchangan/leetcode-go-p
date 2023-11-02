package leetcode

var CanFinish = canFinish

func canFinish(numCourses int, prerequisites [][]int) bool {
	coursePrerequisites := make([]int, numCourses)
	edges := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		course1 := prerequisites[i][0]
		course2 := prerequisites[i][1]
		coursePrerequisites[course1]++
		edges[course2] = append(edges[course2], course1)
	}

	can := []int{}
	for i := 0; i < numCourses; i++ {
		if coursePrerequisites[i] == 0 {
			can = append(can, i)
		}
	}

	for len(can) > 0 {
		course := can[0]
		for _, c := range edges[course] {
			coursePrerequisites[c]--
			if coursePrerequisites[c] == 0 {
				can = append(can, c)
			}
		}
		can = can[1:]
	}

	for i := 0; i < numCourses; i++ {
		if coursePrerequisites[i] != 0 {
			return false
		}
	}

	return true
}
