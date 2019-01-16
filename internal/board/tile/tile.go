package tile

// Tile is part of board
// contains info about snake, food
type Tile struct {
	x, y int

	typ           Type
	snakeNextPart *Tile
}

func newTile(x, y int) *Tile {
	return &Tile{
		x: x,
		y: y,
	}
}

func (t *Tile) SnakeNextPart() *Tile {
	return t.snakeNextPart
}

func (t *Tile) SetSnakeNextPart(next *Tile) {
	t.snakeNextPart = next
}

func (t Tile) IsSnake() bool {
	return t.typ == SnakeHead || t.typ == SnakeBody || t.typ == SnakeBottom
}

func (t Tile) IsFood() bool {
	return t.typ == Food
}

func (t Tile) Type() Type {
	return t.typ
}

func (t *Tile) SetType(typ Type) {
	t.typ = typ
}

func (t *Tile) SetEmpty() {
	t.typ = Empty
	t.snakeNextPart = nil
}

func (t *Tile) Pos() (int, int) {
	return t.x, t.y
}

func (t *Tile) X() int {
	return t.x
}

func (t *Tile) Y() int {
	return t.y
}

func (t *Tile) Compare(c *Tile) bool {
	return t.x == c.x && t.y == c.y
}
