// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package framebuffer

import (
	"image"
	"image/color"
	"unsafe"
)

type BGR565 struct {
	Pix    []byte
	Rect   image.Rectangle
	Stride int
}

func (i *BGR565) Bounds() image.Rectangle { return i.Rect }
func (i *BGR565) ColorModel() color.Model { return RGB565Model }

func (i *BGR565) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(i.Rect)) {
		return RGBColor{}
	}

	clr := *(*uint16)(unsafe.Pointer(&i.Pix[i.PixOffset(x, y)]))

	return RGB565Color{
		uint8(clr) & mask5,
		uint8(clr>>5) & mask6,
		uint8(clr>>11) & mask5,
	}
}

func (i *BGR565) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(i.Rect)) {
		return
	}

	cc := rgb565Model(c).(RGB565Color)
	clr := uint16(cc.G)<<11 | uint16(cc.G)<<6 | uint16(cc.R)
	*(*uint16)(unsafe.Pointer(&i.Pix[i.PixOffset(x, y)])) = clr
}

func (i *BGR565) PixOffset(x, y int) int {
	return (y-i.Rect.Min.Y)*i.Stride + (x-i.Rect.Min.X)*2
}
