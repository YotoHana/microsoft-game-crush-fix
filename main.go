package main

import (
	"os/exec"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("MicroFix")
	w.Resize(fyne.NewSize(300, 80))
	btnOn := widget.NewButton("Enable fix", func ()  {
		exec.Command("taskkill", "/f", "/im", "explorer.exe").Run()
		w.Content().Refresh()
	})
	btnOff := widget.NewButton("Disable fix", func ()  {
		exec.Command("explorer").Start()
		w.Content().Refresh()
	})
	content := container.NewVBox(btnOn, btnOff)
	w.SetContent(content)
	w.SetFixedSize(true)
	w.ShowAndRun()
}