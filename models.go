// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package framebuffer

import (
	"image/color"
)

const (
	mask8 = 1<<8 - 1
	mask5 = 1<<5 - 1
	mask6 = 1<<6 - 1
)

type RGBColor struct {
	R, G, B uint8
}

type RGB565Color struct {
	R, G, B uint8
}

type RGB555Color struct {
	R, G, B uint8
}

var (
	RGBModel    = color.ModelFunc(rgbModel)
	RGB555Model = color.ModelFunc(rgb555Model)
	RGB565Model = color.ModelFunc(rgb565Model)
)

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGBColor); ok {
		return c
	}

	r, g, b, _ := c.RGBA()
	return RGBColor{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
}

func rgb555Model(c color.Color) color.Color {
	if _, ok := c.(RGB555Color); ok {
		return c
	}

	r, g, b, _ := c.RGBA()
	return RGB555Color{
		uint8(r>>(8+8-5)) & mask5,
		uint8(g>>(8+8-5)) & mask5,
		uint8(b>>(8+8-5)) & mask5,
	}
}

func rgb565Model(c color.Color) color.Color {
	if _, ok := c.(RGB565Color); ok {
		return c
	}

	r, g, b, _ := c.RGBA()
	return RGB565Color{
		uint8(r>>(8+8-5)) & mask5,
		uint8(g>>(8+8-6)) & mask6,
		uint8(b>>(8+8-5)) & mask5,
	}
}

func (c RGBColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 0xffff
	return
}

func (c RGB555Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r = r<<11 | r<<6 | r<<1 | r>>4
	g = uint32(c.G)
	g = g<<11 | g<<6 | g<<1 | g>>4
	b = uint32(c.B)
	b = b<<11 | b<<6 | b<<1 | b>>4
	a = 0xffff
	return
}

func (c RGB565Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r = r<<11 | r<<6 | r<<1 | r>>4
	g = uint32(c.G)
	g = g<<10 | g<<4 | g>>2
	b = uint32(c.B)
	b = b<<11 | b<<6 | b<<1 | b>>4
	a = 0xffff
	return
}

// RGB image...split file?!
/*
func (i *RGB555) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(i.Rect)) {
		return RGBColor{}
	}

	clr := *(*uint16)(unsafe.Pointer(&i.Pix[i.PixOffset(x, y)]))

	return RGB555Color{
		uint8(clr>>10) & mask5,
		uint8(clr>>5) & mask5,
		uint8(clr) & mask5,
	}
}


type RGB struct {
	Pix    []byte
	Rect   image.Rectangle
	Stride int
}

func (i *RGB) Bounds() image.Rectangle { return i.Rect }
func (i *RGB) PixOffset(x, y int) int { return y * i.Stride + x * 4 }
func (i *RGB) ColorModel() color.Model { return RGBModel }
func (i *RGB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(i.Rect)) { return }
	o := i.PixOffset(x, y)
	c1 := color.NRGBAModel.Convert(c).(color.NRGBA)
	i.Pix[o+0] = c1.B
	i.Pix[o+1] = c1.G
	i.Pix[o+2] = c1.R
}

func (i *RGB) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(i.Rect)) {
		return RGBColor{}
	}
	o := i.PixOffset(x, y)
    return RGBColor {
        i.Pix[o+2] >>16,
        i.Pix[o+1] >>8,
        i.Pix[o+0],
    }
//	return color.NRGBA{i.Pix[o+2], i.Pix[o+1], i.Pix[o+0], 255}
}

*/
