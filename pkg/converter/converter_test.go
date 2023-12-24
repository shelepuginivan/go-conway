package converter

import (
	"reflect"
	"testing"
)

func TestIsVerticalBorder(t *testing.T) {
	cases := []struct {
		width    int
		x        int
		expected bool
	}{
		{x: 0, width: 5, expected: true},
		{x: 1, width: 5, expected: false},
		{x: 1, width: 2, expected: true},
		{x: 0, width: 2, expected: true},
		{x: 3, width: 10, expected: false},
		{x: 9, width: 10, expected: true},
	}

	for _, c := range cases {
		actual := isVerticalBorder(c.width, c.x)

		if actual != c.expected {
			t.Errorf("Expected %t, got %t", c.expected, actual)
		}
	}
}

func TestIsHorizontalBorder(t *testing.T) {
	cases := []struct {
		height   int
		y        int
		expected bool
	}{
		{y: 0, height: 5, expected: true},
		{y: 1, height: 5, expected: false},
		{y: 1, height: 2, expected: true},
		{y: 0, height: 2, expected: true},
		{y: 3, height: 10, expected: false},
		{y: 9, height: 10, expected: true},
	}

	for _, c := range cases {
		actual := isHorizontalBorder(c.height, c.y)

		if actual != c.expected {
			t.Errorf("Expected %t, got %t", c.expected, actual)
		}
	}
}

func TestIsCorner(t *testing.T) {
	cases := []struct {
		width    int
		height   int
		x        int
		y        int
		expected bool
	}{
		{x: 0, y: 0, width: 5, height: 4, expected: true},
		{x: 1, y: 0, width: 5, height: 4, expected: false},
		{x: 1, y: 0, width: 2, height: 4, expected: true},
		{x: 0, y: 3, width: 2, height: 4, expected: true},
		{x: 3, y: 2, width: 10, height: 15, expected: false},
		{x: 9, y: 14, width: 10, height: 15, expected: true},
	}

	for _, c := range cases {
		actual := isCorner(c.width, c.height, c.x, c.y)

		if actual != c.expected {
			t.Errorf("Expected %t, but got %t", c.expected, actual)
		}
	}
}

func TestConverterGameGridToCharGrid(t *testing.T) {
	c := New('@', '.', '|', '-', '+')

	gameGrid := [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}

	expectedCharGrid := []string{
		"+---+",
		"|.@.|",
		"|.@.|",
		"|.@.|",
		"+---+",
	}

	actualGrid := c.GameGridToCharGrid(gameGrid)

	if !reflect.DeepEqual(actualGrid, expectedCharGrid) {
		t.Errorf("Expected char grid: %v, but got %v", expectedCharGrid, actualGrid)
	}
}

func TestConverterCharGridToGameGrid(t *testing.T) {
	c := New('@', '.', '|', '-', '+')

	charGrid := []string{
		"+---+",
		"|.@.|",
		"|@..|",
		"|.@.|",
		"+---+",
	}

	expectedGameGrid := [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, true, false, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}

	actualGrid := c.CharGridToGameGrid(charGrid)

	if !reflect.DeepEqual(actualGrid, expectedGameGrid) {
		t.Errorf("Expected char grid: %v, but got %v", expectedGameGrid, actualGrid)
	}
}
