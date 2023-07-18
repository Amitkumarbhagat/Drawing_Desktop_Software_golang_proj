package ui

import (
	"PIXL/apptype"
	"PIXL/swatch"
	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixelWindow fyne.Window
	state       *apptype.State
	Swatches    []*swatch.Swatch
}
