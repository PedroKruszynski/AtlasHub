package layouts

import "fyne.io/fyne/v2"

type MarginLayout struct {
	MarginPercent float32
}

func NewMarginLayout(percent float32) *MarginLayout {
	return &MarginLayout{MarginPercent: percent}
}

func (m *MarginLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) == 0 {
		return
	}

	obj := objects[0]

	marginX := size.Width * m.MarginPercent
	marginY := size.Height * m.MarginPercent

	newSize := fyne.NewSize(
		size.Width-(marginX*2),
		size.Height-(marginY*2),
	)

	obj.Resize(newSize)
	obj.Move(fyne.NewPos(marginX, marginY))
}

func (m *MarginLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) == 0 {
		return fyne.NewSize(0, 0)
	}

	min := objects[0].MinSize()
	return fyne.NewSize(
		min.Width/(1-(2*m.MarginPercent)),
		min.Height/(1-(2*m.MarginPercent)),
	)
}
