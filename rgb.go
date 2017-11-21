// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package framebuffer

import (
	"image"
	"image/color"
	"unsafe"
)

type RGB struct {
	Pix    []byte
	Rect   image.Rectangle
	Stride int
}

func (i *RGB) Bounds() image.Rectangle { return i.Rect }
func (i *RGB) ColorModel() color.Model { return RGB555Model }

func (i *RGB) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(i.Rect)) {
		return RGBColor{}
	}

	clr := *(*uint32)(unsafe.Pointer(&i.Pix[i.PixOffset(x, y)]))

	return RGBColor{
		uint8(clr>>16) & mask8,
		uint8(clr>>8) & mask8,
		uint8(clr) & mask8,
	}
}

func (i *RGB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(i.Rect)) {
		return
	}

	cc := rgbModel(c).(RGBColor)

	p := i.PixOffset(x, y)
	i.Pix[p+0] = cc.R
	i.Pix[p+1] = cc.G
	i.Pix[p+2] = cc.B

	//	clr := uint16(cc.R)<<16 | uint16(cc.G)<<8 | uint16(cc.B)
	//	*(*uint16)(unsafe.Pointer(&i.Pix[i.PixOffset(x, y)])) = clr
}

func (i *RGB) PixOffset(x, y int) int {
	return (y-i.Rect.Min.Y)*i.Stride + (x-i.Rect.Min.X)*4
}
