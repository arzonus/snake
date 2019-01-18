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

	Rect image.Rectangle

	TextHeight int
	TextWeight int
}

func (bta ButtonTextAlignment) Vector() (x, y int) {

	w := bta.Rect.Dx() / 2
	h := bta.Rect.Dy() / 2

	switch bta.Weight {
	case ButtonTextAlignmentWeightLeft:
		x = -w
	case ButtonTextAlignmentWeightRight:
		x = w - bta.TextWeight
	case ButtonTextAlignmentWeightMiddle:
		x = -bta.TextWeight / 2
	}
	switch bta.Height {
	case ButtonTextAlignmentHeightUp:
		y = -h + bta.TextHeight
	case ButtonTextAlignmentHeightDown:
		y = h
	case ButtonTextAlignmentHeightMiddle:
		y = bta.TextHeight / 2
	}

	// add indent and shift left corner
	x += bta.IndentX + w
	y += bta.IndentY + h

	return
}
