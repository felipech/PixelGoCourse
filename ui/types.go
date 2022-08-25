package ui

import (
	"PixelGo/apptype"
	"PixelGo/pxcanvas"
	"PixelGo/swatch"
	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas  *pxcanvas.PxCanvas
	PixilWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
