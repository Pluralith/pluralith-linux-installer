package ui

import (
	"bytes"
	"fmt"
	"image"

	"gioui.org/layout"
	"gioui.org/op/paint"
)

func ImageBox(gtx layout.Context, size image.Point, imageBytes []byte) layout.Dimensions {
	img, _, loadErr := image.Decode(bytes.NewReader(imageBytes))
	if loadErr != nil {
		fmt.Println(loadErr)
	}

	imageOp := paint.NewImageOp(img)
	imageOp.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return layout.Dimensions{Size: size}
}
