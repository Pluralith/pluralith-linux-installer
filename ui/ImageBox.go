package ui

import (
	"bytes"
	"fmt"
	"image"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
)

func ImageBox(gtx layout.Context, size image.Point, imageBytes []byte) layout.Dimensions {
	img, _, loadErr := image.Decode(bytes.NewReader(imageBytes))
	if loadErr != nil {
		fmt.Println(loadErr)
	}

	imageOp := paint.NewImageOp(img)
	dims := widget.Image{Src: imageOp, Scale: 0.5, Fit: widget.ScaleDown}.Layout(gtx)

	return dims
}
