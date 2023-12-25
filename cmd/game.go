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
				termbox.Flush()
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()

	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyArrowUp:
				if y > 1 {
					removeCursor(x, y)
					y--
				}
			case termbox.KeyArrowDown:
				if y < height-2 {
					removeCursor(x, y)
					y++
				}
			case termbox.KeyArrowLeft:
				if x > 1 {
					removeCursor(x, y)
					x--
				}
			case termbox.KeyArrowRight:
				if x < width-2 {
					removeCursor(x, y)
					x++
				}
			case termbox.KeyPgup:
				removeCursor(x, y)
				y = 1
			case termbox.KeyPgdn:
				removeCursor(x, y)
				y = height - 2
			case termbox.KeyHome:
				removeCursor(x, y)
				x = 1
			case termbox.KeyEnd:
				removeCursor(x, y)
				x = width - 2
			case termbox.KeySpace:
				if engine.GetCell(x, y) {
					setDead(x, y)
					engine.SetCell(x, y, false)
				} else {
					setAlive(x, y)
					engine.SetCell(x, y, true)
				}
			case termbox.KeyDelete:
				setDead(x, y)
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
				removeCursor(x, y)
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
		termbox.Flush()
	}
}

func drawCursor(x, y int) {
	termbox.SetBg(x, y, termbox.ColorRed)
}

func setAlive(x, y int) {
	termbox.SetCell(x, y, '@', termbox.AttrBold, termbox.ColorDefault)
}

func setDead(x, y int) {
	termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
}

func removeCursor(x, y int) {
	termbox.SetBg(x, y, termbox.ColorDefault)
}

func drawGrid(c *conway.Conway) {
	for x := 1; x < c.Width-1; x++ {
		for y := 1; y < c.Height-1; y++ {
			if c.GetCell(x, y) {
				setAlive(x, y)
			} else {
				setDead(x, y)
			}
		}
	}
}
