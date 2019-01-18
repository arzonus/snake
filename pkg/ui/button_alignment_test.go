package ui

import (
	"github.com/stretchr/testify/assert"
	"image"
	"testing"
)

func TestButtonTextAlignment_Vector(t *testing.T) {
	var cases = []struct {
		Name string

		BTA ButtonTextAlignment
		X   int
		Y   int
	}{
		{
			Name: "default",
			X:    0,
			Y:    0,
		},
		{
			Name: "height middle, weight middle, zero rectangle",
			BTA: ButtonTextAlignment{
				Weight: ButtonTextAlignmentWeightMiddle,
				Height: ButtonTextAlignmentHeightMiddle,

				TextHeight: 10,
				TextWeight: 20,
			},
			X: -10,
			Y: 5,
		},
		{
			Name: "height middle, weight middle",
			BTA: ButtonTextAlignment{
				Weight: ButtonTextAlignmentWeightMiddle,
				Height: ButtonTextAlignmentHeightMiddle,

				TextHeight: 10,
				TextWeight: 20,
				Rect:       image.Rect(0, 0, 30, 20),
			},
			X: 5,
			Y: 15,
		},
		{
			Name: "height middle, weight right",
			BTA: ButtonTextAlignment{
				Weight: ButtonTextAlignmentWeightRight,
				Height: ButtonTextAlignmentHeightMiddle,

				TextHeight: 10,
				TextWeight: 20,
				Rect:       image.Rect(0, 0, 30, 20),
			},
			X: 10,
			Y: 15,
		},
		{
			Name: "height middle, weight left",
			BTA: ButtonTextAlignment{
				Weight: ButtonTextAlignmentWeightLeft,
				Height: ButtonTextAlignmentHeightMiddle,

				TextHeight: 10,
				TextWeight: 20,
				Rect:       image.Rect(0, 0, 30, 20),
			},
			X: 0,
			Y: 15,
		},
		{
			Name: "height up, weight middle",
			BTA: ButtonTextAlignment{
				Height: ButtonTextAlignmentHeightUp,
				Weight: ButtonTextAlignmentWeightMiddle,

				TextHeight: 10,
				TextWeight: 20,
				Rect:       image.Rect(0, 0, 30, 20),
			},
			X: 5,
			Y: 10,
		},
		{
			Name: "height up, weight left",
			BTA: ButtonTextAlignment{
				Height: ButtonTextAlignmentHeightUp,
				Weight: ButtonTextAlignmentWeightLeft,

				TextHeight: 10,
				TextWeight: 20,
				Rect:       image.Rect(0, 0, 30, 20),
			},
			X: 0,
			Y: 10,
		},
		{
			Name: "height up, weight right",
			BTA: ButtonTextAlignment{
				Height: ButtonTextAlignmentHeightUp,
				Weight: ButtonTextAlignmentWeightRight,

				TextHeight: 10,
				TextWeight: 20,
				Rect:       image.Rect(0, 0, 30, 20),
			},
			X: 10,
			Y: 10,
		},
		{
			Name: "height down, weight middle",
			BTA: ButtonTextAlignment{
				Height: ButtonTextAlignmentHeightDown,
				Weight: ButtonTextAlignmentWeightMiddle,

				TextHeight: 10,
				TextWeight: 20,
				Rect:       image.Rect(0, 0, 30, 20),
			},
			X: 5,
			Y: 20,
		},
		{
			Name: "height down, weight left",
			BTA: ButtonTextAlignment{
				Height: ButtonTextAlignmentHeightDown,
				Weight: ButtonTextAlignmentWeightLeft,

				TextHeight: 10,
				TextWeight: 20,
				Rect:       image.Rect(0, 0, 30, 20),
			},
			X: 0,
			Y: 20,
		},
		{
			Name: "height down, weight right",
			BTA: ButtonTextAlignment{
				Height: ButtonTextAlignmentHeightDown,
				Weight: ButtonTextAlignmentWeightRight,

				TextHeight: 10,
				TextWeight: 20,
				Rect:       image.Rect(0, 0, 30, 20),
			},
			X: 10,
			Y: 20,
		},
		{
			Name: "height middle, weight middle, x, y ident",
			BTA: ButtonTextAlignment{
				Weight: ButtonTextAlignmentWeightMiddle,
				Height: ButtonTextAlignmentHeightMiddle,

				TextHeight: 10,
				TextWeight: 20,
				Rect:       image.Rect(0, 0, 30, 20),
				IndentX:    5,
				IndentY:    5,
			},
			X: 10,
			Y: 20,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			x, y := c.BTA.Vector()
			assert.Equal(t, c.X, x, "x")
			assert.Equal(t, c.Y, y, "y")
		})
	}
}
