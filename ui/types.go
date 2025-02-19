package ui

import (
	"example.com/apptype"
	"example.com/pxcanvas"
	"example.com/swatch"
	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
