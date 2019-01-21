package ui

import "image"

type ButtonTextAlignmentHeight int

const (
	ButtonTextAlignmentHeightMiddle ButtonTextAlignmentHeight = iota
	ButtonTextAlignmentHeightUp
	ButtonTextAlignmentHeightDown
)

type ButtonTextAlignmentWeight int

const (
	ButtonTextAlignmentWeightMiddle ButtonTextAlignmentWeight = iota
	ButtonTextAlignmentWeightLeft
	ButtonTextAlignmentWeightRight
)

type ButtonTextAlignment struct {
	Weight ButtonTextAlignmentWeight
	Height ButtonTextAlignmentHeight

	// Indent is calculated relative text center
	IndentX int
	IndentY int

	rect image.Rectangle

	textHeight int
	textWeight int
}

func (bta ButtonTextAlignment) Vector() (x, y int) {

	w := bta.rect.Dx() / 2
	h := bta.rect.Dy() / 2

	switch bta.Weight {
	case ButtonTextAlignmentWeightLeft:
		x = -w
	case ButtonTextAlignmentWeightRight:
		x = w - bta.textWeight
	case ButtonTextAlignmentWeightMiddle:
		x = -bta.textWeight / 2
	}
	switch bta.Height {
	case ButtonTextAlignmentHeightUp:
		y = -h + bta.textHeight
	case ButtonTextAlignmentHeightDown:
		y = h
	case ButtonTextAlignmentHeightMiddle:
		y = bta.textHeight / 2
	}

	// add indent and shift left corner
	x += bta.IndentX + w
	y += bta.IndentY + h

	return
}

func (bta ButtonTextAlignment) Pos(x, y int) (int, int) {
	vx, vy := bta.Vector()
	return x + vx, y + vy
}
