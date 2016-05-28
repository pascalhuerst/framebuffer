// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package framebuffer

import (
	"image"
	"image/color"
	"unsafe"
)

type BGR555 struct {
	Pix    []byte
	Rect   image.Rectangle
	Stride int
}

func (i *BGR555) Bounds() image.Rectangle { return i.Rect }
func (i *BGR555) ColorModel() color.Model { return RGB555Model }

func (i *BGR555) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(i.Rect)) {
		return RGBColor{}
	}

	clr := *(*uint16)(unsafe.Pointer(&i.Pix[i.PixOffset(x, y)]))

	return RGB555Color{
		uint8(clr) & mask5,
		uint8(clr>>5) & mask5,
		uint8(clr>>10) & mask5,
	}
}

func (i *BGR555) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(i.Rect)) {
		return
	}

	cc := rgb555Model(c).(RGB555Color)
	clr := uint16(cc.B)<<10 | uint16(cc.G)<<5 | uint16(cc.R)
	*(*uint16)(unsafe.Pointer(&i.Pix[i.PixOffset(x, y)])) = clr
}

func (i *BGR555) PixOffset(x, y int) int {
	return (y-i.Rect.Min.Y)*i.Stride + (x-i.Rect.Min.X)*2
}
