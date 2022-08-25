package main

import (
	"PixelGo/apptype"
	"PixelGo/pxcanvas"
	"PixelGo/swatch"
	"PixelGo/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"image/color"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")
	state := apptype.State{
		BrushColor: color.NRGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		SwatchSelected: 0,
	}

	pxCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       30,
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pxCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas:  pixlCanvas,
		PixilWindow: pixlWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	appInit.PixilWindow.ShowAndRun()
}
