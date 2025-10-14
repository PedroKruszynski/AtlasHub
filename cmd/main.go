package main

import (
	screensaver "atlasHub/internal/screens/screensaver"
	usdbrl "atlasHub/internal/screens/usdbrl"
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

	// --------------------------
	// Tela Home
	// --------------------------
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

	// --------------------------
	// Tela Dolar
	// --------------------------
	telaDolarLabel := widget.NewLabelWithStyle(
		"Tela Dólar",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)
	telaDolarContent := container.NewVBox(
		layout.NewSpacer(),
		telaDolarLabel,
		btnVoltar,
		layout.NewSpacer(),
	)

	// --------------------------
	// Tela Authenticator
	// --------------------------
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

	// --------------------------
	// Tela de Descanso (extraída)
	// --------------------------
	var stack *fyne.Container
	var currentScreen fyne.CanvasObject

	descansoContent := screensaver.ScreensaverContainer(func() {
		currentScreen = homeContent
		stack.Objects = []fyne.CanvasObject{homeContent}
		stack.Refresh()
	})

	resetTimer := screensaver.StartScreensaverTimer(2*time.Second, func() {
		stack.Objects = []fyne.CanvasObject{descansoContent}
		stack.Refresh()
	})

	// --------------------------
	// Stack principal
	// --------------------------
	stack = container.NewStack(
		homeContent,
		telaDolarContent,
		telaAuthenticatorContent,
		descansoContent,
	)
	currentScreen = homeContent
	stack.Objects = []fyne.CanvasObject{currentScreen}

	// ===========================
	// Botões de navegação
	// ===========================
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

	myWindow.SetContent(stack)
	myWindow.ShowAndRun()
}
