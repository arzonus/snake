package menu

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"log"
	"unicode/utf8"
)

var (
	arcadeFont font.Face
)

const (
	fontSize = 32
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
}

type Menu struct {
	w, h int
}

func New(w, h int) *Menu {
	return &Menu{
		w: w,
		h: h,
	}
}

func (m *Menu) Draw(image *ebiten.Image) error {
	text.Draw(image, fmt.Sprintf("SNAKE"), arcadeFont, 0, 0, color.Black)
	size := utf8.RuneCountInString("SNAKE") * 32
	text.Draw(image, fmt.Sprintf("NEW"), arcadeFont, size, fontSize, color.Black)
	text.Draw(image, fmt.Sprintf("SNAKENEW"), arcadeFont, 0, fontSize*2, color.Black)

	return nil
}
