package ui

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
)

// Types
type C = layout.Context
type D = layout.Dimensions

// Variables
var progress float32
var progressIncrementer chan float32
var installing bool

func Draw(window *app.Window) error {
	progressIncrementer = make(chan float32)
	// th := material.NewTheme(gofont.Collection())

	var ops op.Ops
	for {
		select {
		case e := <-window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				e.Frame(gtx.Ops)
			}

		case p := <-progressIncrementer:
			if installing && progress < 1 {
				progress += p
				window.Invalidate()
			}
		}
	}
	return nil
}
