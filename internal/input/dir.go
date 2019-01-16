package input

// Dir represents a direction.
type Dir int

const (
	DirUp Dir = iota
	DirRight
	DirDown
	DirLeft
)

// String returns a string representing the direction.
func (d Dir) String() string {
	switch d {
	case DirUp:
		return "Up"
	case DirRight:
		return "Right"
	case DirDown:
		return "Down"
	case DirLeft:
		return "Left"
	}
	panic("not reach")
}

// Vector returns a [-1, 1] value for each axis.
func (d Dir) Vector() (x, y int) {
	switch d {
	case DirUp:
		return 0, -1
	case DirRight:
		return 1, 0
	case DirDown:
		return 0, 1
	case DirLeft:
		return -1, 0
	}
	panic("not reach")
}

// Pos returns a [x-1, y+1] values
func (d Dir) Pos(x, y int) (posx int, posy int) {
	switch d {
	case DirUp:
		return x, y - 1
	case DirDown:
		return x, y + 1
	case DirRight:
		return x + 1, y
	case DirLeft:
		return x - 1, y
	}
	panic("not reach")
}
