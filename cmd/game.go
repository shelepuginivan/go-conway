package cmd

import (
	"time"

	"github.com/nsf/termbox-go"
	"github.com/shelepuginivan/go-conway/pkg/conway"
	"github.com/shelepuginivan/go-conway/pkg/cursor"
)

func Game() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	width, height := termbox.Size()
	c := cursor.New(1, 1, width, height)
	engine := conway.New(conway.EmptyGrid(width, height))
	running := false

	c.Draw()
	drawBorder(termbox.ColorBlack)

	if err := termbox.Flush(); err != nil {
		return err
	}

	go func() {
		for {
			if running {
				drawGrid(engine.Tick())
				drawBorder(termbox.ColorGreen)

				if err := termbox.Flush(); err != nil {
					break
				}
			}

			time.Sleep(200 * time.Millisecond)
		}
	}()

	for {
		x, y := c.Get()
		c.Remove()

		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyArrowUp:
				c.Up()
			case termbox.KeyArrowDown:
				c.Down()
			case termbox.KeyArrowLeft:
				c.Left()
			case termbox.KeyArrowRight:
				c.Right()
			case termbox.KeyPgup:
				c.Goto(x, 1)
			case termbox.KeyPgdn:
				c.Goto(x, height-2)
			case termbox.KeyHome:
				c.Goto(1, y)
			case termbox.KeyEnd:
				c.Goto(width-2, y)
			case termbox.KeySpace:
				currentState := engine.GetCell(x, y)
				drawCell(x, y, !currentState)
				engine.SetCell(x, y, !currentState)
			case termbox.KeyDelete:
				engine.Clear()
			case termbox.KeyEnter:
				engine.Tick()
			case termbox.KeyEsc, termbox.KeyCtrlQ, termbox.KeyCtrlC:
				return nil
			}
			switch event.Ch {
			case 'q':
				return nil
			case 'r':
				engine.FillRandom()
			case 'x':
				running = !running
			case 'z':
				engine.Tick()
			}
		case termbox.EventMouse:
			c.Goto(event.MouseX, event.MouseY)
		case termbox.EventResize:
			if x > event.Width-2 {
				x = event.Width - 2
			}

			if y > event.Height-2 {
				y = event.Height - 2
			}

			c.Goto(x, y)
		}

		drawGrid(&engine)
		drawBorder(termbox.ColorBlack)
		c.Draw()

		if err := termbox.Flush(); err != nil {
			return err
		}
	}
}

func drawCell(x, y int, alive bool) {
	if alive {
		termbox.SetCell(x, y, '@', termbox.AttrBold, termbox.ColorDefault)
	} else {
		termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
}

func drawGrid(c *conway.Conway) {
	for x := 1; x < c.Width-1; x++ {
		for y := 1; y < c.Height-1; y++ {
			drawCell(x, y, c.GetCell(x, y))
		}
	}
}

func drawBorder(color termbox.Attribute) {
	width, height := termbox.Size()

	for x := 0; x < width; x++ {
		termbox.SetCell(x, 0, ' ', color, color)
		termbox.SetCell(x, height-1, ' ', color, color)
	}

	for y := 1; y < height-1; y++ {
		termbox.SetCell(0, y, ' ', color, color)
		termbox.SetCell(width-1, y, ' ', color, color)
	}
}
