package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello Fyne")

	hello := widget.NewLabel("مرحبًا بك في Fyne!")
	myWindow.SetContent(container.NewVBox(
		hello,
		widget.NewButton("اضغط هنا", func() {
			hello.SetText("تم الضغط!")
		}),
	))

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}
