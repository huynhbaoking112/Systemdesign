package main

type NeighborSum struct {
	indices []indexes
	grid    [][]int
}

type indexes struct {
	row int
	col int
}

func Constructor(grid [][]int) NeighborSum {
	indices := make([]indexes, len(grid)*len(grid), len(grid)*len(grid))

	for i := range grid {
		for j := range grid[i] {
			indices[grid[i][j]] = indexes{
				row: i,
				col: j,
			}
		}
	}

	return NeighborSum{indices: indices, grid: grid}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	row, col := this.indices[value].row, this.indices[value].col

	sum := 0

	if row > 0 {
		sum += this.grid[row-1][col]
	}

	if row < len(this.grid)-1 {
		sum += this.grid[row+1][col]
	}

	if col > 0 {
		sum += this.grid[row][col-1]
	}

	if col < len(this.grid)-1 {
		sum += this.grid[row][col+1]
	}

	return sum
}

func (this *NeighborSum) DiagonalSum(value int) int {
	row, col := this.indices[value].row, this.indices[value].col

	sum := 0

	if row > 0 && col > 0 {
		sum += this.grid[row-1][col-1]
	}

	if row > 0 && col < len(this.grid)-1 {
		sum += this.grid[row-1][col+1]
	}

	if row < len(this.grid)-1 && col > 0 {
		sum += this.grid[row+1][col-1]
	}

	if row < len(this.grid)-1 && col < len(this.grid)-1 {
		sum += this.grid[row+1][col+1]
	}

	return sum
}
