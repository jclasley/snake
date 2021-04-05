package main

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

const (
	RIGHT = 1 + iota
	DOWN 
	LEFT
	UP
	ESC
)

var (
	keyChan = make(chan keyPress)
)

type keyPress struct{
	ch rune
	direction int
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	
	go listen(keyChan)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if e := termbox.Flush(); e != nil {
		panic(e)
	}
run:

	for {
		select {
		case key := <-keyChan:
			if key.ch == 'q' {
				break run
			}
			switch key.direction {
			case 1:
				termbox.SetCell(0,0,'R', termbox.ColorBlue, termbox.ColorDefault)
			case 2:
				termbox.SetChar(0,0,'D')
			case 3:
				termbox.SetChar(0,0,'L')
			case 4:
				termbox.SetChar(0,0,'U')
			case 5:
				break run
			}
			
		}
		termbox.Flush()
	}
}

func createNextFood(x, y int) {
	rx := rand.Intn(x)
	ry := rand.Intn(y)
	termbox.SetCell(rx, ry, '█', termbox.ColorGreen, termbox.AttrBold)
}

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
			default:
				if e.Key == 'q' {
					ch <- keyPress{ch: 'q'}
				}
			}
		case termbox.EventError:
			panic(e.Err)
		}
	}
}