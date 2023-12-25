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

	go func() {
		for {
			if running {
				drawGrid(engine.Tick())

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
				drawCell(x, y, false)
				engine.SetCell(x, y, false)
			case termbox.KeyEnter:
				engine.Tick()
			case termbox.KeyCtrlS:
				running = !running
			case termbox.KeyEsc, termbox.KeyCtrlQ, termbox.KeyCtrlC:
				return nil
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
