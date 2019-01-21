package ui

import (
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
	"image"
)

type Button struct {
	rect     image.Rectangle
	imageOpt ebiten.DrawImageOptions

	text string
	font font.Face
	bta  ButtonTextAlignment

	mouseDown bool
	onPressed func(b *Button)
}

func (b *Button) ()  {
	
}
//
//func (b *Button) Update() {
//	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
//		x, y := ebiten.CursorPosition()
//		if b.Rect.Min.X <= x && x < b.Rect.Max.X && b.Rect.Min.Y <= y && y < b.Rect.Max.Y {
//			b.mouseDown = true
//		} else {
//			b.mouseDown = false
//		}
//	} else {
//		if b.mouseDown {
//			if b.onPressed != nil {
//				b.onPressed(b)
//			}
//		}
//		b.mouseDown = false
//	}
//}
//
//func (b *Button) textDraw(image *ebiten.Image) {
//
//}
//
//func (b *Button) TextPos() (x int, y int) {
//	switch b.TextAlignment {
//	case ButtonTextAlignmentUp, ButtonTextAlignmentUpLeft, ButtonTextAlignmentUpRight:
//		return 0, 0
//	case ButtonTextAlignmentDown, ButtonTextAlignmentDownLeft, ButtonTextAlignmentDownRight:
//		return 0, 0
//	case ButtonTextAlignmentRight, ButtonTextAlignmentUpRight, ButtonTextAlignmentDownRight:
//		return 0, 0
//	case ButtonTextAlignmentLeft, ButtonTextAlignmentUpLeft, ButtonTextAlignmentDownLeft:
//		return 0, 0
//	default:
//		panic("not reach aligment")
//	}
//}
//
//func (b *Button) Draw(dst *ebiten.Image) {
//	t := imageTypeButton
//	if b.mouseDown {
//		t = imageTypeButtonPressed
//	}
//	drawNinePatches(dst, b.Rect, imageSrcRects[t])
//
//	bounds, _ := font.BoundString(uiFont, b.Text)
//	w := (bounds.Max.X - bounds.Min.X).Ceil()
//	x := b.Rect.Min.X + (b.Rect.Dx()-w)/2
//	y := b.Rect.Max.Y - (b.Rect.Dy()-uiFontMHeight)/2
//	text.Draw(dst, b.Text, uiFont, x, y, color.Black)
//	)
//
//
//	func(b *Button) SetOnPressed(f
//	func(b *Button)) {
//		b.onPressed = f
//	}
