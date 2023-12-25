package conway

import (
	"reflect"
	"testing"
)

func TestConwayGetCell(t *testing.T) {
	game := New([][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	})

	cases := []struct {
		x        int
		y        int
		expected bool
	}{
		{x: 0, y: 0, expected: false},
		{x: 4, y: 1, expected: false},
		{x: 2, y: 2, expected: true},
		{x: 3, y: 2, expected: false},
		{x: 2, y: 3, expected: true},
	}

	for _, c := range cases {
		actual := game.GetCell(c.x, c.y)

		if actual != c.expected {
			t.Errorf("Expected cell to be %t, but got %t", c.expected, actual)
		}
	}
}

func TestConwaySetCell(t *testing.T) {
	game := New([][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	})

	cases := []struct {
		x int
		y int
		v bool
	}{
		{x: 0, y: 0, v: true},
		{x: 4, y: 1, v: false},
		{x: 4, y: 4, v: false},
		{x: 3, y: 1, v: false},
		{x: 2, y: 4, v: false},
		{x: 4, y: 0, v: false},
		{x: 3, y: 2, v: true},
		{x: 2, y: 3, v: false},
	}

	for _, c := range cases {
		game.SetCell(c.x, c.y, c.v)
		newValue := game.GetCell(c.x, c.y)

		if newValue != c.v {
			t.Errorf("Expected cell to be %t, but got %t", c.v, newValue)
		}
	}

}

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

func TestConwayClear(t *testing.T) {
	grid := [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, true, false},
		{false, true, true, false, false},
		{false, false, false, false, false},
	}
	expectedGrid := [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}
	game := New(grid)
	game.Clear()

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
