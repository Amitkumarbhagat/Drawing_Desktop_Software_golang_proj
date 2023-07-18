package ui

import (
	"PIXL/swatch"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		s := swatch.NewSwatch(app.state, initialColor, i, func(s *swatch.Swatch) {
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			app.state.SwatchSelected = s.SwatchIndex
			app.state.BrushColor = s.Color
		})
		if i == 0 {
			s.Selected = true
			app.state.SwatchSelected = 0
			s.Refresh()
		}
		app.Swatches = append(app.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}
	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
