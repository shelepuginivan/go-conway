package converter

type Converter struct {
	AliveChar            rune
	DeadChar             rune
	VerticalBorderChar   rune
	HorizontalBorderChar rune
	CornerChar           rune
}

func New(
	aliveChar,
	deadChar,
	verticalBorderChar,
	horizontalBorderChar,
	cornerChar rune,
) Converter {
	return Converter{
		AliveChar:            aliveChar,
		DeadChar:             deadChar,
		VerticalBorderChar:   verticalBorderChar,
		HorizontalBorderChar: horizontalBorderChar,
		CornerChar:           cornerChar,
	}
}

func (c Converter) GameGridToCharGrid(grid [][]bool) []string {
	width := len(grid[0])
	height := len(grid)

	charGrid := make([]string, height)

	for y, row := range grid {
		for x, cell := range row {
			if isCorner(width, height, x, y) {
				charGrid[y] += string(c.CornerChar)
			} else if isHorizontalBorder(height, y) {
				charGrid[y] += string(c.HorizontalBorderChar)
			} else if isVerticalBorder(width, x) {
				charGrid[y] += string(c.VerticalBorderChar)
			} else if cell {
				charGrid[y] += string(c.AliveChar)
			} else {
				charGrid[y] += string(c.DeadChar)
			}
		}
	}

	return charGrid
}

func (c Converter) CharGridToGameGrid(grid []string) [][]bool {
	gameGrid := make([][]bool, len(grid))

	for y, row := range grid {
		for _, char := range row {
			gameGrid[y] = append(gameGrid[y], char == c.AliveChar)
		}
	}

	return gameGrid
}
