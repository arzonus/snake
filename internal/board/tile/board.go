package tile

import (
	"github.com/arzonus/snake/internal/colors"
	"github.com/arzonus/snake/internal/input"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

// Board store size, tiles and tile draw options
type Board struct {
	w, h  int
	tiles [][]*Tile

	screenWidth  int
	screenHeight int

	tileSize   int
	marginSize int

	image *ebiten.Image
}

func (tb *Board) Image() *ebiten.Image {
	return tb.image
}

// NewBoard returns new Board
// w and h - weight and height of tile board
func NewBoard(w, h int) (*Board, error) {
	b := &Board{
		w:     w,
		h:     h,
		tiles: make([][]*Tile, w),

		tileSize:   10,
		marginSize: 1,
	}

	var err error
	b.image, err = ebiten.NewImage(w*10, h*10, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	for i := 0; i < w; i++ {
		b.tiles[i] = make([]*Tile, h)
		for j := 0; j < h; j++ {
			b.tiles[i][j] = newTile(i, j)
		}
	}
	b.initImages()

	return b, nil
}

// Next returns next tile and true
// or nil and false, if next tile is border
func (tb *Board) Next(dir input.Dir, cur *Tile) (*Tile, bool) {
	nx, ny := dir.Pos(cur.x, cur.y)

	if nx >= tb.w || ny >= tb.h || nx < 0 || ny < 0 {
		return nil, false
	}

	return tb.Tile(nx, ny), true
}

// Tile returns pointer of Tile
func (tb *Board) Tile(x, y int) *Tile {
	return tb.tiles[x][y]
}

func (tb *Board) initImages() {

	foodImage, _ = ebiten.NewImage(tb.tileSize, tb.tileSize, ebiten.FilterDefault)
	foodImage.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})

	snakeBodyImage, _ = ebiten.NewImage(tb.tileSize, tb.tileSize, ebiten.FilterDefault)
	snakeBodyImage.Fill(color.Black)

	snakeHeadImage, _ = ebiten.NewImage(tb.tileSize, tb.tileSize, ebiten.FilterDefault)
	snakeHeadImage.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})

	emptyImage, _ = ebiten.NewImage(tb.tileSize, tb.tileSize, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
}

// Draw response for tile drawing
func (tb *Board) Draw() error {
	if err := tb.image.Fill(colors.Frame); err != nil {
		return err
	}
	for x := 0; x < len(tb.tiles); x++ {
		for y := 0; y < len(tb.tiles[x]); y++ {
			if err := tb.drawTile(tb.tiles[x][y], tb.image); err != nil {
				return err
			}
		}
	}
	return nil
}

func (tb *Board) drawTile(t *Tile, boardImage *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tb.tileSize*t.x), float64(tb.tileSize*t.y))

	switch t.Type() {
	case Empty:
		return boardImage.DrawImage(emptyImage, op)
	case Food:
		return boardImage.DrawImage(foodImage, op)
	case SnakeHead:
		return boardImage.DrawImage(snakeHeadImage, op)
	case SnakeBody:
		return boardImage.DrawImage(snakeBodyImage, op)
	case SnakeBottom:
		return boardImage.DrawImage(snakeBodyImage, op)
	default:
		panic("no reach tile type")
	}
}
