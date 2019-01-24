package board

import (
	"fmt"
	"github.com/arzonus/snake/internal/board/snake"
	"github.com/arzonus/snake/internal/board/tile"
	"github.com/arzonus/snake/internal/input"
	"github.com/hajimehoshi/ebiten"
	"log"
	"math/rand"
)

// Board response for game, snake behavior
type Board struct {
	// board sizes
	w, h int

	tileBoard *tile.Board
	snake     *snake.Snake
	foodTile  *tile.Tile

	count int
	speed int

	image *ebiten.Image
}

func (b *Board) Image() *ebiten.Image {
	return b.image
}

// New return new Board object
// weight and height - is game weight and height
// speed - game speed, more speed - slowly game
// game speed depends on ebiten TPS
func New(weight, height, speed int) (*Board, error) {
	var err error

	b := &Board{
		w:     weight,
		h:     height,
		speed: speed,
	}

	b.tileBoard, err = tile.NewBoard(weight, height)
	if err != nil {
		return nil, err
	}

	// init random snake position
	head := b.tileBoard.Tile(rand.Intn(b.w), rand.Intn(b.h))
	dir := input.DirUp
	bottom, ok := b.tileBoard.Next(input.DirDown, head)
	if !ok {
		bottom, _ = b.tileBoard.Next(input.DirUp, head)
		dir = input.DirDown
	}
	b.snake = snake.New(head, bottom, dir)

	b.generateFood()

	b.image, err = ebiten.NewImage(weight*10, height*10, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	log.Println(fmt.Sprintf("weight: %d, height: %d, speed: %d", weight, height, speed))
	return b, nil
}

// Update response for updating snake direction
func (b *Board) Update(input *input.Input) error {
	if dir, ok := input.Dir(); ok {
		b.snake.ChangeDir(dir)
	}
	return b.Step()
}

var GameOverError = fmt.Errorf("board: gameover")

// Step response for game process
func (b *Board) Step() error {
	b.count++
	if b.count%b.speed != 0 {
		return nil
	}

	if !b.foodTile.IsFood() {
		b.generateFood()
	}

	if err := b.snake.Move(b.tileBoard); err != nil {
		if err == snake.BodyError || err == snake.BorderError {
			log.Println(err)
			return GameOverError
		}
		return err
	}

	b.count = 0
	return nil
}

// generateFood generate food tile
func (b *Board) generateFood() {
	t := b.tileBoard.Tile(rand.Intn(b.w), rand.Intn(b.h))

	if t.IsSnake() {
		// TODO: refactor
		b.generateFood()
	}

	b.foodTile = t
	b.foodTile.SetEmpty()
	b.foodTile.SetType(tile.Food)
}

func (b *Board) Draw() error {
	if err := b.tileBoard.Draw(); err != nil {
		return err
	}

	return b.image.DrawImage(b.tileBoard.Image(), &ebiten.DrawImageOptions{})
}
