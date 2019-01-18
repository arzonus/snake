package game

import (
	"github.com/arzonus/snake/internal/board"
	"github.com/arzonus/snake/internal/colors"
	"github.com/arzonus/snake/internal/input"
	"github.com/arzonus/snake/internal/menu"
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
		menu:  menu.New(ScreenWidth, ScreenHeight),
	}, nil
}

type Game struct {
	w, h int

	input *input.Input
	board *board.Board
	menu  *menu.Menu
}

func (g *Game) Update() error {
	g.input.Update()
	if err := g.board.Update(g.input); err != nil {
		return err
	}

	//if err := g.board.Step(); err != nil {
	//	return err
	//}

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
		return game.Draw(screen)
	}
}

func (g *Game) Draw(screen *ebiten.Image) error {
	screen.Fill(colors.Background)
	if err := g.menu.Draw(screen); err != nil {
		return err
	}

	return nil
	g.board.Draw(screen)

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.board.Image.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))

	return screen.DrawImage(g.board.Image, &ebiten.DrawImageOptions{})
}
