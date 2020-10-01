package Ruleset

type ConwaysRuleset struct {

}

func (ruleset *ConwaysRuleset) IsAlive(field *[][]bool, x int, y int) bool {
		aliveNeighbor := getClosestAliveCount(field, x, y)
		result := (*field)[x][y]
		if result {
			if aliveNeighbor != 2 && aliveNeighbor != 3 {
				return false
			}
		} else {
			if aliveNeighbor == 3 {
				return true
			}
		}

		return result
	}

	func check(field *[][]bool, x,y int, counter *int8) {
		if (*field)[x][y] {
			*counter += 1
		}
	}

	func getClosestAliveCount(field *[][]bool, x int, y int) int8 {
		width := len(*field)
		height := len((*field)[0])

		//    1_2_3
		// 1| 1 2 3
		// 2| 4 5 6
		// 3| 7 8 9
		var alive int8

		var i1 = (x + width - 1) % width
		var i3 = (x + width + 1) % width
		var j1 = (y + height - 1) % height
		var j3 = (y + height + 1) % height
		check(field, i1, j1, &alive)
		check(field, i1, y, &alive)
		check(field, i1, j3, &alive)
		check(field, x, j1, &alive)
		check(field, x, j3, &alive)
		check(field, i3, j1, &alive)
		check(field, i3, y, &alive)
		check(field, i3, j3, &alive)

		return alive
	}
