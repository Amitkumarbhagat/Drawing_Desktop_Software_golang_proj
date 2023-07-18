package main

import (
	"PIXL/apptype"
	"PIXL/swatch"
	"PIXL/ui"
	"fyne.io/fyne/v2/app"
	"image/color"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{R: 255, G: 255, B: 255, A: 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixelWindow: pixlWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}
	ui.Setup(&appInit)
	appInit.PixelWindow.ShowAndRun()
}
