package main

import (
	"atlasHub/internal/usdbrl"
	homeBackground "atlasHub/static/home"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("AtlasHub")
	myWindow.Resize(fyne.NewSize(1024, 600))

	img := canvas.NewImageFromResource(homeBackground.BackgroundPng)

	homeLabel := widget.NewLabelWithStyle(
		"\nAtlasHub Menu",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	btn1 := widget.NewButton("Ir para Tela 1", nil)
	btn2 := widget.NewButton("Tap me", func() {})
	btn3 := widget.NewButton("Tap me", func() {})

	homeContent := container.New(
		layout.NewStackLayout(),
		img,
		container.NewVBox(
			homeLabel,
			container.New(
				layout.NewGridLayout(2),
				btn1,
				btn2,
				btn3,
			),
		),
	)

	tela1Label := widget.NewLabelWithStyle(
		"Tela 1",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)
	btnVoltar := widget.NewButton("Voltar para Home", nil)
	tela1Content := container.NewVBox(
		layout.NewSpacer(),
		tela1Label,
		btnVoltar,
		layout.NewSpacer(),
	)

	stack := container.NewStack()
	homeLayer := homeContent
	tela1Layer := tela1Content

	stack.Add(homeLayer)

	btn1.OnTapped = func() {
		// busca cotação
		cot, err := usdbrl.FetchDollar()
		var texto string
		if err != nil {
			texto = "Erro ao buscar cotação"
		} else {
			texto = fmt.Sprintf("USD = R$ %.2f", cot)
		}
		tela1Label.SetText(texto)
		stack.Objects = []fyne.CanvasObject{tela1Layer}
		stack.Refresh()
	}
	btnVoltar.OnTapped = func() {
		stack.Objects = []fyne.CanvasObject{homeLayer}
		stack.Refresh()
	}

	myWindow.SetContent(stack)

	myWindow.ShowAndRun()
}
