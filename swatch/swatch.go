package swatch

import (
	"PixelGo/apptype"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}
	swatch.ExtendBaseWidget(swatch)
	return swatch
}

func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	rectangle := canvas.NewRectangle(s.Color)
	objects := []fyne.CanvasObject{rectangle}
	return &SwatchRenderer{
		square:  *rectangle,
		objects: objects,
		parent:  s,
	}
}
