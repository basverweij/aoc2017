package main

type pos struct {
	row, col int
}

func regionCount(g *grid) int {
	// intialize regions from hash bits
	regions := make([][]int, 128)
	for row := range regions {
		rowHash := g.rowHash(row)

		regions[row] = make([]int, 128)
		for i, b := range rowHash {
			for j := 0; j < 8; j++ {
				if (b>>uint(j))&1 == 1 {
					regions[row][i*8+7-j] = -1
				}
			}
		}
	}

	n := 0
	var todo []pos

	// map used bits to regions
	for row := 0; row < 128; row++ {
		for col := 0; col < 128; col++ {
			if regions[row][col] >= 0 {
				// skip free & already mapped bits
				continue
			}

			// we found the start of a new region
			n++
			todo = []pos{{row, col}}

			for len(todo) > 0 {
				p := todo[0]
				todo = todo[1:]

				// set current region
				regions[p.row][p.col] = n

				// scan above
				if p.row > 0 && regions[p.row-1][p.col] == -1 {
					todo = append(todo, pos{p.row - 1, p.col})
				}

				// scan below
				if p.row < 127 && regions[p.row+1][p.col] == -1 {
					todo = append(todo, pos{p.row + 1, p.col})
				}

				// scan left
				if p.col > 0 && regions[p.row][p.col-1] == -1 {
					todo = append(todo, pos{p.row, p.col - 1})
				}

				// scan right
				if p.col < 127 && regions[p.row][p.col+1] == -1 {
					todo = append(todo, pos{p.row, p.col + 1})
				}
			}
		}
	}

	return n
}
