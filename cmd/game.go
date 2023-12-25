package cmd

import (
	"time"

	"github.com/nsf/termbox-go"
	"github.com/shelepuginivan/go-conway/pkg/conway"
)

func Game() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	width, height := termbox.Size()
	engine := conway.New(conway.EmptyGrid(width, height))
	running := false

	x, y := 1, 1
	drawCursor(x, y)

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
		removeCursor(x, y)

		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:

			switch event.Key {
			case termbox.KeyArrowUp:
				if y > 1 {
					y--
				}
			case termbox.KeyArrowDown:
				if y < height-2 {
					y++
				}
			case termbox.KeyArrowLeft:
				if x > 1 {
					x--
				}
			case termbox.KeyArrowRight:
				if x < width-2 {
					x++
				}
			case termbox.KeyPgup:
				y = 1
			case termbox.KeyPgdn:
				y = height - 2
			case termbox.KeyHome:
				x = 1
			case termbox.KeyEnd:
				x = width - 2
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
			if event.MouseX > 1 && event.MouseX < width-2 && event.MouseY > 1 && event.MouseY < height-2 {
				x, y = event.MouseX, event.MouseY
			}
		case termbox.EventResize:
			if x > event.Width-2 {
				x = event.Width - 2
			}

			if y > event.Height-2 {
				y = event.Height - 2
			}
		}

		drawGrid(&engine)
		drawCursor(x, y)

		if err := termbox.Flush(); err != nil {
			return err
		}
	}
}

func drawCursor(x, y int) {
	termbox.SetBg(x, y, termbox.ColorRed)
}

func drawCell(x, y int, alive bool) {
	if alive {
		termbox.SetCell(x, y, '@', termbox.AttrBold, termbox.ColorDefault)
	} else {
		termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
}

func removeCursor(x, y int) {
	termbox.SetBg(x, y, termbox.ColorDefault)
}

func drawGrid(c *conway.Conway) {
	for x := 1; x < c.Width-1; x++ {
		for y := 1; y < c.Height-1; y++ {
			drawCell(x, y, c.GetCell(x, y))
		}
	}
}
