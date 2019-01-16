package tile

type Type uint

const (
	Empty Type = iota
	Food
	SnakeHead
	SnakeBody
	SnakeBottom
)
