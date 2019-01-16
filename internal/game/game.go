package game

import (
	"github.com/arzonus/snake/internal/board"
	"github.com/arzonus/snake/internal/colors"
	"github.com/arzonus/snake/internal/input"
	"github.com/hajimehoshi/ebiten"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 400
	ScreenHeight = 400
)

func New() (*Game, error) {
	return &Game{
		input: input.New(),
		board: board.New(40, 40, 15),
	}, nil
}

type Game struct {
	input *input.Input
	board *board.Board
}

func (g *Game) Update() error {
	g.input.Update()
	if err := g.board.Update(g.input); err != nil {
		return err
	}

	if err := g.board.Step(); err != nil {
		return err
	}

	return nil
}

func Update(game *Game) func(screen *ebiten.Image) error {
	return func(screen *ebiten.Image) error {
		if err := game.Update(); err != nil {
			return err
		}
		if ebiten.IsDrawingSkipped() {
			return nil
		}
		game.Draw(screen)
		return nil
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colors.Background)
	g.board.Draw(screen)

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.board.Image.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))

	screen.DrawImage(g.board.Image, &ebiten.DrawImageOptions{})
}
