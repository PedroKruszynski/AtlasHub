package screensaver

import (
	"time"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ScreensaverContainer(onClick func()) fyne.CanvasObject {
	btn := widget.NewButton("Clique para voltar ao menu", func() {
		if onClick != nil {
			onClick()
		}
	})

	return container.NewStack(btn)
}

func StartScreensaverTimer(inactivity time.Duration, onTimeout func()) func() {
	timer := time.NewTimer(inactivity)

	resetTimer := func() {
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(inactivity)
	}

	go func() {
		for {
			<-timer.C
			if onTimeout != nil {
				onTimeout()
			}
			resetTimer()
		}
	}()

	return resetTimer
}
