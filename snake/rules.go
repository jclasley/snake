package snake

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"errors"
)

var (
	w, h int
	area = w * h // used to determine if player has won the game -- length of snake == area
	block = '█'
	green = termbox.ColorGreen
	white = termbox.ColorWhite
	food *coord
	gameEvent = make(chan int)
)

const (
	MOVING = iota
	DEAD
)

type coord struct{
	x int
	y int
}

type Snake struct {
	length int
	segments []*coord
	direction int
}

func (s *Snake) GetHead() (int, int) {
	return s.segments[len(s.segments) - 1].x, s.segments[len(s.segments) - 1].y
}

func (s *Snake) Move() {
	hx, hy := s.GetHead()

	switch s.direction {
	case RIGHT:
		hx++
	case LEFT:
		hx--
	case UP:
		hy--
	case DOWN:
		hy++
	}
	if err := s.CheckDeath(hx, hy); err != nil {
		gameEvent <- DEAD
	}
	s.segments = append(s.segments, &coord{hx, hy})

	termbox.SetFg(s.segments[0].x, s.segments[0].y, termbox.AttrHidden)
	//drop "tail"
	s.segments = s.segments[1:]

	if checkFood(s.GetHead()) {
		s.segments = append(s.segments, food)
		newFood()
	}

	for _, v := range s.segments {
		termbox.SetCell(v.x, v.y, block, white, termbox.ColorDefault) // move head
	}
	termbox.Flush()
}

func checkFood(x,y int) bool {
	return x == food.x && y == food.y
}

func newFood() {
	x, y := randLocationInBounds()
	food = &coord{x, y}
	termbox.SetCell(x, y, block, green, termbox.ColorDefault)
	termbox.Flush()
}

func randLocationInBounds() (int, int) {
	rx := rand.Intn(w)
	ry := rand.Intn(h)
	return rx, ry
}

func (s *Snake) RenderInit() {
	x, y := randLocationInBounds()
	s.segments = append(s.segments, &coord{x, y})
}

func (s *Snake) CheckDeath(x,y int) error {
	c := termbox.GetCell(x, y)
	if c.Ch == block && c.Fg == white {
		return errors.New("You died.")
	}
	return nil
}