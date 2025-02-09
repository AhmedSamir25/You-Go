package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("You Go")

	bgImage := canvas.NewImageFromFile("background.jpg")
	bgImage.FillMode = canvas.ImageFillStretch

	hello := widget.NewLabel("Welcome to You Go")
	button := widget.NewButton("Download Now", func() {
		hello.SetText("تم الضغط!")
	})

	content := container.NewMax(
		bgImage,
		container.NewVBox(
			container.NewCenter(hello),
			container.NewCenter(button),
		),
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 500))
	myWindow.ShowAndRun()
}
