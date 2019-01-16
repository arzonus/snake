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
}

// NewBoard returns new Board
// w and h - weight and height of tile board
func NewBoard(w, h int) *Board {
	b := &Board{
		w:     w,
		h:     h,
		tiles: make([][]*Tile, w),

		tileSize:   10,
		marginSize: 1,
	}

	for i := 0; i < w; i++ {
		b.tiles[i] = make([]*Tile, h)
		for j := 0; j < h; j++ {
			b.tiles[i][j] = newTile(i, j)
		}
	}
	b.initImages()

	return b
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
func (tb *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(colors.Frame)
	for x := 0; x < len(tb.tiles); x++ {
		for y := 0; y < len(tb.tiles[x]); y++ {
			tb.drawTile(tb.tiles[x][y], boardImage)
		}
	}
}

func (tb *Board) drawTile(t *Tile, boardImage *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tb.tileSize*t.x), float64(tb.tileSize*t.y))

	switch t.Type() {
	case Empty:
		boardImage.DrawImage(emptyImage, op)
		return
	case Food:
		boardImage.DrawImage(foodImage, op)
		return
	case SnakeHead:
		boardImage.DrawImage(snakeHeadImage, op)
		return
	case SnakeBody:
		boardImage.DrawImage(snakeBodyImage, op)
		return
	case SnakeBottom:
		boardImage.DrawImage(snakeBodyImage, op)
		return
	default:
		panic("no reach tile type")
	}
}
