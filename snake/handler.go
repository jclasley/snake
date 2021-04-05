package snake

import (
	"time"
	"github.com/nsf/termbox-go"
)

const (
	RIGHT = 1 + iota
	DOWN
	LEFT
	UP
	ESC
)

func listen(ch chan keyPress) {

	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch e := termbox.PollEvent(); e.Type {
		case termbox.EventKey:
			switch e.Key {
			case termbox.KeyArrowRight:
				ch <- keyPress{direction: RIGHT}
			case termbox.KeyArrowDown:
				ch <- keyPress{direction: DOWN}
			case termbox.KeyArrowLeft:
				ch <- keyPress{direction: LEFT}
			case termbox.KeyArrowUp:
				ch <- keyPress{direction: UP}
			case termbox.KeyEsc:
				ch <- keyPress{direction: ESC}
			}
		case termbox.EventError:
			panic(e.Err)
		}
	}
}

func (s *Snake) HandlePresses() {
	
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if e := termbox.Flush(); e != nil {
		panic(e)
	}

run:

	for {
		select {
		case key := <-keyChan:
			switch key.direction {
			case 5:
				break run
			default:
				s.direction = key.direction
				s.Move()
			}
		default:
			s.Move()
			time.Sleep(100 * time.Millisecond)
		}
	}
}