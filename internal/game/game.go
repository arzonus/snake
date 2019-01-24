package game

import (
	"fmt"
	"github.com/arzonus/snake/internal/board"
	"github.com/arzonus/snake/internal/colors"
	"github.com/arzonus/snake/internal/input"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth   = 400
	ScreenHeight  = 400
	fontSize      = 32
	smallFontSize = fontSize / 2
)

var (
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	smallArcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func New() (*Game, error) {
	var g = new(Game)
	g.input = input.New()

	return g, nil
}

type gameMode int

const (
	menuMode gameMode = iota
	boardMode
	gameoverMode
)

type Game struct {
	w, h int

	input *input.Input
	board *board.Board

	mode gameMode
}

func (g *Game) Update() error {
	g.input.Update()
	if err := g.board.Update(g.input); err != nil {
		return err
	}

	return nil
}

func Update(game *Game) func(screen *ebiten.Image) error {
	return func(screen *ebiten.Image) error {

		if err := game.update(); err != nil {
			return err
		}

		if ebiten.IsDrawingSkipped() {
			return nil
		}
		return game.Draw(screen)
	}
}

func (g *Game) update() error {
	switch g.mode {
	case boardMode:
		if err := g.Update(); err != nil {
			if err == board.GameOverError {
				g.mode = gameoverMode
				return nil
			}
			return err
		}
		return nil
	case gameoverMode:
		if ebiten.IsKeyPressed(ebiten.KeySpace) ||
			ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) ||
			ebiten.TouchIDs() != nil {
			g.mode = menuMode
			return nil
		}
	case menuMode:
		if ebiten.IsKeyPressed(ebiten.KeySpace) ||
			ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) ||
			ebiten.TouchIDs() != nil {
			g.mode = boardMode

			var err error
			g.board, err = board.New(40, 40, 15)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) error {
	if err := screen.Fill(colors.Background); err != nil {
		return err
	}

	switch g.mode {
	case menuMode:
		return g.drawMenu(screen)
	case gameoverMode:
		return g.drawGameOver(screen)
	case boardMode:
		return g.drawBoard(screen)
	}

	return nil
}

func drawCentredTxt(txt string, image *ebiten.Image, fnt font.Face, color color.Color, row int) {
	x := ScreenWidth/2 - font.MeasureString(fnt, txt).Ceil()/2
	y := ScreenHeight/2 + (2*row+1)*arcadeFont.Metrics().Height.Ceil()/2
	text.Draw(image, txt, fnt, x, y, color)
}

func (g *Game) drawMenu(screen *ebiten.Image) error {
	drawCentredTxt("START", screen, arcadeFont, color.Black, 0)
	drawCentredTxt("PRESS TOUCH OR PRESS KEY", screen, smallArcadeFont, color.Black, 1)
	return nil
}

func (g *Game) drawGameOver(screen *ebiten.Image) error {
	drawCentredTxt("GAMEOVER", screen, arcadeFont, color.Black, -2)
	drawCentredTxt(fmt.Sprintf("YOUR RESULT: %d", g.board.Result()), screen, smallArcadeFont, color.Black, 0)
	drawCentredTxt("PRESS TOUCH OR PRESS KEY", screen, smallArcadeFont, color.Black, 1)
	drawCentredTxt("TO START NEW GAME", screen, smallArcadeFont, color.Black, 2)
	return nil
}

func (g *Game) drawBoard(screen *ebiten.Image) error {
	if err := g.board.Draw(); err != nil {
		return err
	}
	return screen.DrawImage(g.board.Image(), &ebiten.DrawImageOptions{})
}
