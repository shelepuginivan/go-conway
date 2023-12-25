package conway

type Conway struct {
	Width  int
	Height int
	Grid   [][]bool
}

func New(grid [][]bool) Conway {
	width := len(grid[0])
	height := len(grid)

	return Conway{
		Width:  width,
		Height: height,
		Grid:   grid,
	}
}

func (c Conway) GetCell(x, y int) bool {
	return c.Grid[y][x]
}

func (c *Conway) SetCell(x, y int, value bool) *Conway {
	c.Grid[y][x] = value
	return c
}

func (c *Conway) Tick() *Conway {
	newGrid := EmptyGrid(c.Width, c.Height)

	for x := 1; x < c.Width-1; x++ {
		for y := 1; y < c.Height-1; y++ {
			sum := c.sumNeighbours(x, y)

			if c.GetCell(x, y) {
				newGrid[y][x] = sum == 2 || sum == 3
			} else {
				newGrid[y][x] = sum == 3
			}
		}
	}

	c.Grid = newGrid
	return c
}

func (c *Conway) Clear() *Conway {
	c.Grid = EmptyGrid(c.Width, c.Height)
	return c
}

func (c *Conway) FillRandom() *Conway {
	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			isBorder := x == 0 || y == 0 || x == c.Width-1 || y == c.Height-1
			c.SetCell(x, y, !isBorder && randomBool())
		}
	}

	return c
}

func (c Conway) sumNeighbours(x, y int) int {
	return boolToInt(c.GetCell(x-1, y-1)) +
		boolToInt(c.GetCell(x-1, y)) +
		boolToInt(c.GetCell(x-1, y+1)) +
		boolToInt(c.GetCell(x, y-1)) +
		boolToInt(c.GetCell(x, y+1)) +
		boolToInt(c.GetCell(x+1, y-1)) +
		boolToInt(c.GetCell(x+1, y)) +
		boolToInt(c.GetCell(x+1, y+1))
}
