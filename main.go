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

	telaDolarLabel := widget.NewLabelWithStyle(
		"Tela 1",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)
	telaDolarContent := container.NewVBox(
		layout.NewSpacer(),
		telaDolarLabel,
		btnVoltar,
		layout.NewSpacer(),
	)
	telaAuthenticatorLabel := widget.NewLabelWithStyle(
		"Authenticator",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)
	telaAuthenticatorContent := container.NewVBox(
		layout.NewSpacer(),
		telaAuthenticatorLabel,
		btnVoltar,
		layout.NewSpacer(),
	)

	descansoBtn := widget.NewButton("Clique para voltar ao menu", nil)
	descansoContent := container.NewStack(descansoBtn)

	stack := container.NewStack(
		homeContent,
		telaDolarContent,
		telaAuthenticatorContent,
		descansoContent,
	)
	currentScreen := homeContent
	stack.Objects = []fyne.CanvasObject{currentScreen}

	inactivity := 10 * time.Second
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

	btnDolar.OnTapped = func() {
		resetTimer()
		cot, err := usdbrl.FetchDollar()
		var texto string
		if err != nil {
			texto = "Erro ao buscar cotação"
		} else {
			texto = fmt.Sprintf("USD = R$ %.2f", cot)
		}
		telaDolarLabel.SetText(texto)
		currentScreen = telaDolarContent
		stack.Objects = []fyne.CanvasObject{telaDolarContent}
		stack.Refresh()
	}

	btnAuthenticator.OnTapped = func() {
		resetTimer()
		texto := "Tela do Authenticator"
		telaDolarLabel.SetText(texto)
		currentScreen = telaAuthenticatorContent
		stack.Objects = []fyne.CanvasObject{telaAuthenticatorContent}
		stack.Refresh()
	}

	btnVoltar.OnTapped = func() {
		resetTimer()
		currentScreen = homeContent
		stack.Objects = []fyne.CanvasObject{homeContent}
		stack.Refresh()
	}

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
