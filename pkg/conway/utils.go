package conway

import (
	"math/rand"
)

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func EmptyGrid(width, height int) [][]bool {
	mat := make([][]bool, height)

	for i := range mat {
		mat[i] = make([]bool, width)
	}

	return mat
}
