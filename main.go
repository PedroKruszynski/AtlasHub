package main

import (
	"atlasHub/internal/usdbrl"
	homeBackground "atlasHub/static/home"
	"fmt"
	"time"

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

	btnDolar := widget.NewButton("Cotação Dolar", nil)
	btnAuthenticator := widget.NewButton("Authenticator", func() {})
	btn3 := widget.NewButton("Tap me", func() {})
	btnVoltar := widget.NewButton("Voltar para Home", nil)

	homeContent := container.New(
		layout.NewStackLayout(),
		img,
		container.NewVBox(
			homeLabel,
			container.New(
				layout.NewGridLayout(2),
				btnDolar,
				btnAuthenticator,
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

	descansoBtn := widget.NewButton("Clique para voltar ao menu", nil)
	descansoContent := container.NewStack(descansoBtn)

	stack := container.NewStack(homeContent, tela1Content, descansoContent)
	currentScreen := homeContent
	stack.Objects = []fyne.CanvasObject{currentScreen}

	inactivity := 3 * time.Second
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

	btn1.OnTapped = func() {
		resetTimer()
		cot, err := usdbrl.FetchDollar()
		var texto string
		if err != nil {
			texto = "Erro ao buscar cotação"
		} else {
			texto = fmt.Sprintf("USD = R$ %.2f", cot)
		}
		tela1Label.SetText(texto)
		currentScreen = tela1Content
		stack.Objects = []fyne.CanvasObject{tela1Content}
		stack.Refresh()
	}

	btnVoltar.OnTapped = func() {
		resetTimer()
		currentScreen = homeContent
		stack.Objects = []fyne.CanvasObject{homeContent}
		stack.Refresh()
	}

	btn2.OnTapped = func() { resetTimer() }
	btn3.OnTapped = func() { resetTimer() }

	descansoBtn.OnTapped = func() {
		resetTimer()
		currentScreen = homeContent
		stack.Objects = []fyne.CanvasObject{homeContent}
		stack.Refresh()
	}

	go func() {
		for {
			<-timer.C
			stack.Objects = []fyne.CanvasObject{descansoContent}
			stack.Refresh()
		}
	}()

	myWindow.SetContent(stack)

	myWindow.ShowAndRun()
}
