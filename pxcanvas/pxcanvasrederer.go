package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
	canvasCursor []fyne.CanvasObject
}

func (p *PxCanvasRenderer) SetCursor(objects []fyne.CanvasObject) {
	p.canvasCursor = objects
}

func (p *PxCanvasRenderer) MinSize() fyne.Size {
	return p.pxCanvas.DrawingArea
}

func (p *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)

	for i := 0; i < len(p.canvasBorder); i++ {
		objects = append(objects, &p.canvasBorder[i])

	}
	objects = append(objects, p.canvasImage)
	objects = append(objects, p.canvasCursor...)
	return objects
}

func (r *PxCanvasRenderer) Destroy() {

}

func (p *PxCanvasRenderer) Layout(size fyne.Size) {
	p.LayoutCanvas(size)
	p.LayoutBorder(size)
}

func (p *PxCanvasRenderer) Refresh() {
	if p.pxCanvas.reloadImage {
		p.canvasImage = canvas.NewImageFromImage(p.pxCanvas.PixelData)
		p.canvasImage.ScaleMode = canvas.ImageScalePixels
		p.canvasImage.FillMode = canvas.ImageFillContain
		p.pxCanvas.reloadImage = false

	}
	p.Layout(p.pxCanvas.Size())
	canvas.Refresh(p.canvasImage)
}

func (p *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := p.pxCanvas.PxCols
	imgPxHeigth := p.pxCanvas.PxRows
	pxSize := p.pxCanvas.PxSize
	p.canvasImage.Move(fyne.NewPos(p.pxCanvas.CanvasOffset.X, p.pxCanvas.CanvasOffset.Y))
	p.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeigth*pxSize)))

}

func (p *PxCanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := p.pxCanvas.CanvasOffset
	imgPxWidth := p.canvasImage.Size().Width
	imgPxHeigth := p.canvasImage.Size().Height

	left := &p.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgPxHeigth)

	top := &p.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imgPxWidth, offset.Y)

	right := &p.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imgPxWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X+imgPxWidth, offset.Y+imgPxHeigth)

	bottom := &p.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgPxHeigth)
	bottom.Position2 = fyne.NewPos(offset.X+imgPxWidth, offset.Y+imgPxHeigth)

}
