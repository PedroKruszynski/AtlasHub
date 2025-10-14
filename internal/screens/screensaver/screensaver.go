package screensaver

import (
	homeBackground "atlasHub/static/home"
	"time"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type invisibleClickable struct {
	widget.BaseWidget
	onTap func()
}

func NewInvisibleClickable(onTap func()) *invisibleClickable {
	w := &invisibleClickable{onTap: onTap}
	w.ExtendBaseWidget(w)
	return w
}

func (w *invisibleClickable) Tapped(_ *fyne.PointEvent) {
	if w.onTap != nil {
		w.onTap()
	}
}

func (w *invisibleClickable) TappedSecondary(_ *fyne.PointEvent) {}

func (w *invisibleClickable) CreateRenderer() fyne.WidgetRenderer {
	return &invisibleRenderer{w: w}
}

type invisibleRenderer struct {
	w *invisibleClickable
}

func (r *invisibleRenderer) Layout(size fyne.Size) {}
func (r *invisibleRenderer) MinSize() fyne.Size    { return fyne.NewSize(0, 0) }
func (r *invisibleRenderer) Refresh()              {}

func (r *invisibleRenderer) Objects() []fyne.CanvasObject { return nil }
func (r *invisibleRenderer) Destroy()                     {}

func ScreensaverContainer(onClick func()) fyne.CanvasObject {
	img := canvas.NewImageFromResource(homeBackground.BackgroundPng)
	img.FillMode = canvas.ImageFillStretch

	clickable := NewInvisibleClickable(onClick)

	return container.NewStack(
		img,
		clickable,
	)
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
