package snake

import (
	"fmt"
	"github.com/arzonus/snake/internal/board/tile"
	"github.com/arzonus/snake/internal/input"
	"log"
)

type Snake struct {
	head   *tile.Tile
	bottom *tile.Tile

	// TODO: softboard
	dir          input.Dir
	isChangedDir bool
}

func New(head *tile.Tile, bottom *tile.Tile, dir input.Dir) *Snake {
	s := &Snake{
		head:   head,
		bottom: bottom,
		dir:    dir,
	}

	s.head.SetType(tile.SnakeHead)
	s.bottom.SetType(tile.SnakeBottom)
	s.bottom.SetSnakeNextPart(s.head)

	return s
}

// ChangeDir - change snake direction
// don't allow multiple change dir between steps
func (s *Snake) ChangeDir(dir input.Dir) {
	if !s.isChangedDir && s.dir != dir &&
		((s.dir == input.DirUp && dir != input.DirDown) ||
			(s.dir == input.DirRight && dir != input.DirLeft) ||
			(s.dir == input.DirLeft && dir != input.DirRight) ||
			(s.dir == input.DirDown && dir != input.DirUp)) {
		log.Println("changed dir to", dir)
		s.isChangedDir = true
		s.dir = dir
	}

}

var BorderError = fmt.Errorf("snake: next tile is border")
var BodyError = fmt.Errorf("snake: next tile is snake")

// Move - move snake to next tile
func (s *Snake) Move(tb *tile.Board) error {
	log.Println(fmt.Sprintf("x: %d, y: %d", s.head.X(), s.head.Y()))

	t, ok := tb.Next(s.dir, s.head)
	if !ok {
		return BorderError
	}

	// if tile is not food, bottom move to next tile
	// if it is food, bottom stay on the tile
	if !t.IsFood() {
		t.SetEmpty()
		bottom := s.bottom.SnakeNextPart()
		bottom.SetType(tile.SnakeBottom)
		s.bottom.SetEmpty()
		s.bottom = bottom
	}

	if t.IsSnake() {
		return BodyError
	}

	s.head.SetType(tile.SnakeBody)
	s.head.SetSnakeNextPart(t)

	s.head = t
	s.head.SetType(tile.SnakeHead)

	s.isChangedDir = false
	return nil
}
