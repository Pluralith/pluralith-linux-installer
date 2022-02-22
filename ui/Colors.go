package ui

import "image/color"

func RGB(c uint32) color.NRGBA {
	return transformHex(0xff000000 | c)
}

func transformHex(c uint32) color.NRGBA {
	return color.NRGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}
