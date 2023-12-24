package converter

func isVerticalBorder(width, x int) bool {
	return x == 0 || x == width-1
}

func isHorizontalBorder(height, y int) bool {
	return y == 0 || y == height-1
}

func isCorner(width, height, x, y int) bool {
	return x == 0 && y == 0 ||
		x == width-1 && y == 0 ||
		x == width-1 && y == height-1 ||
		x == 0 && y == height-1
}
