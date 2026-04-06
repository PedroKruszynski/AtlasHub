package main

import (
	"atlasHub/internal/screens/layouts"
	screensaver "atlasHub/internal/screens/screensaver"
	"fmt"
	"image/color"
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

	// img := canvas.NewImageFromResource(homeBackground.BackgroundPng)

	// --------------------------
	// Tela Home
	// --------------------------
	homeLabel := widget.NewLabelWithStyle(
		"\nAtlasHub Menu",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	btnAuthenticator := widget.NewButton("Authenticator", func() {})
	btnTasks := widget.NewButton("Tasks", func() {})
	btnVoltar := widget.NewButton("Voltar para Home", nil)
	bg := canvas.NewRectangle(color.NRGBA{R: 242, G: 242, B: 242, A: 255})

	homeContent := container.New(
		layout.NewStackLayout(),
		bg,
		container.NewCenter(
			container.NewPadded(homeLabel),
			container.New(
				layout.NewGridLayout(2),
				btnAuthenticator,
				btnTasks,
			),
		),
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
	// Tela Tasks
	// --------------------------
	telaTasksLabel := widget.NewLabelWithStyle(
		"Tasks",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)
	bgBlack := canvas.NewRectangle(color.Black)

	card := container.NewStack(
		canvas.NewRectangle(color.White),
		container.NewVBox(
			layout.NewSpacer(),
			container.NewCenter(
				telaTasksLabel,
				btnVoltar,
			),
			layout.NewSpacer(),
		),
	)

	telaTasksContent := container.New(
		layouts.NewMarginLayout(0.1),
		container.NewStack(
			bgBlack,
			card,
		),
	)

	// --------------------------
	// Tela de Descanso (extraída)
	// --------------------------
	var stack *fyne.Container
	var currentScreen fyne.CanvasObject
	var resetTimer func()

	descansoContent := screensaver.ScreensaverContainer(func() {
		if resetTimer != nil {
			resetTimer()
		}

		currentScreen = homeContent
		stack.Objects = []fyne.CanvasObject{homeContent}
		stack.Refresh()
	})

	resetTimer = screensaver.StartScreensaverTimer(8*time.Second, func() {
		fmt.Println("TIMEOUT DISPAROU:", time.Now())

		fyne.Do(func() {
			stack.Objects = []fyne.CanvasObject{descansoContent}
			stack.Refresh()
		})
	})

	resetTimer()

	// --------------------------
	// Stack principal
	// --------------------------
	stack = container.NewStack(
		homeContent,
		telaAuthenticatorContent,
		telaTasksContent,
		descansoContent,
	)
	currentScreen = homeContent
	stack.Objects = []fyne.CanvasObject{currentScreen}

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

	btnTasks.OnTapped = func() {
		resetTimer()
		currentScreen = telaTasksContent
		stack.Objects = []fyne.CanvasObject{telaTasksContent}
		stack.Refresh()
	}

	myWindow.SetContent(stack)
	myWindow.ShowAndRun()
}
