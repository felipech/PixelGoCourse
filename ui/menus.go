package ui

import (
	"PixelGo/util"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/png"
	"os"
	"strconv"
)

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, err error) {
		if uri == nil {
			return
		} else {
			err := png.Encode(uri, app.PixlCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PixilWindow)
				return
			}
			app.State.SetFilePath(uri.URI().Path())

		}
	}, app.PixilWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save As...", func() {
		if app.State.FilePath == "" {
			saveFileDialog(app)
		} else {
			tryClose := func(fh *os.File) {
				err := fh.Close()
				if err != nil {
					dialog.ShowError(err, app.PixilWindow)
				}
			}
			fh, err := os.Create(app.State.FilePath)
			defer tryClose(fh)
			if err != nil {
				return
			}
			err = png.Encode(fh, app.PixlCanvas.PixelData)
			if err != nil {
				return
			}

		}
	})
}

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("Tiene que ser un numero positivo")
			}
			if width <= 0 {
				return errors.New("debe ser > 0")
			}
			return nil
		}
		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator
		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)
		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}
		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				pixelWidth := 0
				pixelHeight := 0
				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("ancho invalido"), app.PixilWindow)
				} else {
					pixelWidth, _ = strconv.Atoi(widthEntry.Text)

				}
				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("alto invalido"), app.PixilWindow)
				} else {
					pixelWidth, _ = strconv.Atoi(heightEntry.Text)

				}
				app.PixlCanvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, app.PixilWindow)
	})
}

func BuildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu("File",
		BuildNewMenu(app),
		BuildOpenMenu(app),
		BuildSaveAsMenu(app))
}
func SetupMenus(app *AppInit) {
	menus := BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PixilWindow.SetMainMenu(mainMenu)
}

func BuildOpenMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Open...", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if uri != nil {
				return
			} else {
				images, _, err := image.Decode(uri)
				if err != nil {
					dialog.ShowError(err, app.PixilWindow)
					return
				}
				app.PixlCanvas.LoadImage(images)
				app.State.SetFilePath(uri.URI().Path())
				imgColors := util.GetImageColors(images)
				i := 0
				for c := range imgColors {
					if i == len(app.Swatches) {
						break
					}
					app.Swatches[i].SetColor(c)
				}
			}
		}, app.PixilWindow)
	})
}
