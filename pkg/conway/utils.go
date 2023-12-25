package conway

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func EmptyGrid(width, height int) [][]bool {
	mat := make([][]bool, height)

	for i := range mat {
		mat[i] = make([]bool, width)
	}

	return mat
}
