package snake

import (
	"github.com/nsf/termbox-go"
	"math"
)

func draw() {
	min := math.Min(float64(w), float64(h))
	top := int(min / 10)
	bot := int(min - float64(top))
	size := bot - top
	
	// sides
	for i := top; i <= size; i++ {
		termbox.SetCell(bot, i, '|', white, termbox.ColorDefault)
		termbox.SetCell(i, bot, '-', white, termbox.ColorDefault)
	}
	for i := top; i <= size; i++ {
		termbox.SetCell(top, i, '|', white, termbox.ColorDefault)
		termbox.SetCell(i, top, '-', white, termbox.ColorDefault)
	}
}