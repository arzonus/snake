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

	Image *ebiten.Image
}

// New return new Board object
// weight and height - is game weight and height
// speed - game speed, more speed - slowly game
// game speed depends on ebiten TPS
func New(weight, height, speed int) *Board {
	b := &Board{
		w:         weight,
		h:         height,
		speed:     speed,
		tileBoard: tile.NewBoard(weight, height),
	}

	// init random snake position

	head := b.tileBoard.Tile(rand.Intn(b.w), rand.Intn(b.h))
	// TODO: panic NPE
	bottom, _ := b.tileBoard.Next(input.DirDown, head)
	b.snake = snake.New(head, bottom, input.DirUp)

	b.generateFood()

	b.Image, _ = ebiten.NewImage(weight, height, ebiten.FilterDefault)

	log.Println(fmt.Sprintf("weight: %d, height: %d, speed: %d", weight, height, speed))
	return b
}

// Update response for updating snake direction
func (b *Board) Update(input *input.Input) error {
	if dir, ok := input.Dir(); ok {
		b.snake.ChangeDir(dir)
	}
	return nil
}

// Step response for game process
func (b *Board) Step() error {
	b.count++
	if b.count%b.speed != 0 {
		return nil
	}
	x, y := ebiten.ScreenSizeInFullscreen()

	log.Println("full screen size:", x, y)

	if !b.foodTile.IsFood() {
		b.generateFood()
	}

	if err := b.snake.Move(b.tileBoard); err != nil {
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

func (b *Board) Draw(screen *ebiten.Image) {
	b.tileBoard.Draw(screen)
}
