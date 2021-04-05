package snake

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

var (
	keyChan = make(chan keyPress)
	s *Snake
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
	InitVars()

	NewGame()
	go listen(keyChan)
	s.HandlePresses()
}

func NewGame() {
	s = &Snake{
		length:    1,
		direction: 0,
	}
	s.RenderInit()
	newFood()
	fmt.Println(food.x, food.y)
	x, y := s.GetHead()
	fmt.Println(x, y)
}

func InitVars() {
		w, h = termbox.Size()
}

