package ui

import (
	"image"
	"pluralith-linux-installer/assets"
	"pluralith-linux-installer/backend"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// Types
type C = layout.Context
type D = layout.Dimensions

func Draw(window *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	backend.StateStore.DownloadProgressIncrementer = make(chan float32)

	go backend.DownloadRelease()

	var ops op.Ops
	for {
		select {
		case e := <-window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				// Define layout
				layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceEnd,
				}.Layout(gtx,
					layout.Rigid(
						func(gtx C) D {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								layout.Rigid(func(gtx C) D {
									return layout.Inset{
										Top:    unit.Dp(100),
										Bottom: unit.Dp(90),
										Left:   unit.Dp(130),
										Right:  unit.Dp(130),
									}.Layout(gtx, func(gtx C) D {
										return ImageBox(gtx, image.Pt(240, 240), assets.ImageStore.PluralithIcon)
									})
								}),
								layout.Rigid(func(gtx C) D {
									return layout.Inset{
										Top:    unit.Dp(30),
										Bottom: unit.Dp(20),
										Left:   unit.Dp(60),
										Right:  unit.Dp(60),
									}.Layout(gtx, func(gtx C) D {
										bar := material.ProgressBar(th, backend.StateStore.DownloadProgress)
										bar.Color = RGB(0x5E84FC)
										bar.TrackColor = RGB(0xDCE4ED)
										return bar.Layout(gtx)
									})
								}),
								layout.Rigid(func(gtx C) D {
									return layout.Inset{
										Top:    unit.Dp(0),
										Bottom: unit.Dp(30),
										Left:   unit.Dp(60),
										Right:  unit.Dp(60),
									}.Layout(gtx, func(gtx C) D {
										return layout.Inset{
											Top:    unit.Dp(0),
											Bottom: unit.Dp(0),
											Left:   unit.Dp(55),
											Right:  unit.Dp(0),
										}.Layout(gtx, func(gtx C) D {
											if backend.StateStore.DownloadProgress >= 1 {
												return ImageBox(gtx, image.Pt(290, 60), assets.ImageStore.CompleteBadge)
											}
											return ImageBox(gtx, image.Pt(290, 60), assets.ImageStore.DownloadBadge)
										})
									})
								}),
							)
						},
					),
				)

				e.Frame(gtx.Ops)
			}

		case p := <-backend.StateStore.DownloadProgressIncrementer:
			if backend.StateStore.DownloadProgress < 1 {
				backend.StateStore.DownloadProgress += p
				window.Invalidate()
			} else {
				backend.LaunchPluralith()
				window.Close()
			}
		}
	}
}
