package screensaver

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DescansoScreen(onClick func()) fyne.CanvasObject {
	btn := widget.NewButton("Clique para voltar ao menu", func() {
		if onClick != nil {
			onClick()
		}
	})

	return container.NewStack(btn)
}
