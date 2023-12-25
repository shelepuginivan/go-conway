package cursor

import "testing"

func TestCursor(t *testing.T) {
	c := New(1, 1, 100, 100)
	x, y := c.Get()

	if x != 1 || y != 1 {
		t.Errorf("Expected x, y = 1, 1, but got %d, %d", x, y)
	}

	c.Up()
	c.Right()
	c.Right()
	x, y = c.Get()

	if x != 3 || y != 1 {
		t.Errorf("Expected x, y = 3, 1, but got %d, %d", x, y)
	}

	c.Down()
	c.Down()
	c.Left()
	x, y = c.Get()

	if x != 2 || y != 3 {
		t.Errorf("Expected x, y = 2, 1, but got %d, %d", x, y)
	}

	c.Goto(50, 50)
	x, y = c.Get()

	if x != 50 || y != 50 {
		t.Errorf("Expected x, y = 50, 50, but got %d, %d", x, y)
	}

	c.Goto(1000, 1000)
	x, y = c.Get()

	if x != 50 || y != 50 {
		t.Errorf("Expected x, y = 50, 50, but got %d, %d", x, y)
	}

	c.Goto(-1000, -1000)
	x, y = c.Get()

	if x != 50 || y != 50 {
		t.Errorf("Expected x, y = 50, 50, but got %d, %d", x, y)
	}

	for i := 0; i < 100; i++ {
		c.Left()
		c.Down()
	}

	x, y = c.Get()

	if x != 1 || y != 98 {
		t.Errorf("Expected x, y = 1, 98, but got %d, %d", x, y)
	}

	c.Draw()
	c.Remove()
}
