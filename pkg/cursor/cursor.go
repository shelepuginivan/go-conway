package cursor

import "github.com/nsf/termbox-go"

type Cursor struct {
	X int
	Y int
	W int
	H int
}

func New(x, y, w, h int) Cursor {
	return Cursor{x, y, w, h}
}

func (c Cursor) Remove() {
	termbox.SetBg(c.X, c.Y, termbox.ColorDefault)
}

func (c Cursor) Draw() {
	termbox.SetBg(c.X, c.Y, termbox.ColorRed)
}

func (c *Cursor) Up() {
	if c.Y > 1 {
		c.Y--
	}
}

func (c *Cursor) Down() {
	if c.Y < c.H-2 {
		c.Y++
	}
}

func (c *Cursor) Left() {
	if c.X > 1 {
		c.X--
	}
}

func (c *Cursor) Right() {
	if c.X < c.W-2 {
		c.X++
	}
}

func (c *Cursor) Goto(x, y int) {
	if x >= 1 && x <= c.W-2 && y >= 1 && y <= c.H-2 {
		c.X, c.Y = x, y
	}
}

func (c Cursor) Get() (int, int) {
	return c.X, c.Y
}
