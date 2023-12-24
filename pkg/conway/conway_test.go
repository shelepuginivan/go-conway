package conway

import (
	"reflect"
	"testing"
)

func TestConwayTick(t *testing.T) {
	initialGrid := [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}
	expectedGrid := [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}

	game := New(initialGrid)
	game.Tick()

	if !reflect.DeepEqual(game.Grid, expectedGrid) {
		t.Errorf("Expected grid after tick: %v, but got %v", expectedGrid, game.Grid)
	}
}

func TestConwaySumNeighbours(t *testing.T) {
	testGrid := [][]bool{
		{false, true, false, false, false},
		{true, true, true, false, false},
		{false, true, true, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}

	game := New(testGrid)
	sum := game.sumNeighbours(2, 2)

	if sum != 3 {
		t.Errorf("Expected sum of neighbors to be 3, but got %d", sum)
	}
}
