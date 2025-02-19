package ui

import "fyne.io/fyne/v2/container"

func SetUp(app *AppInit) {
	SetUpMenus(app)
	swachesContainer := BuildSwatches(app)
	colorPicker := setUpColorPicker(app)

	appLayout := container.NewBorder(nil, swachesContainer, nil, colorPicker, app.PixlCanvas)

	app.PixlWindow.SetContent(appLayout)
}
