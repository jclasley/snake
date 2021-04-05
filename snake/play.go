package snake

import (
	"github.com/nsf/termbox-go"
)

var (
	keyChan = make(chan keyPress)
)

type keyPress struct{
	ch rune
	direction int
}

func Start() {
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
				termbox.SetChar(0,0,'R')
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

