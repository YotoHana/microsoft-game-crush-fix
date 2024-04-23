package main

import (
	"os"
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
	btnOn := widget.NewButton("Запустить фикс", killtask)
	btnOff := widget.NewButton("Отключить фикс", addtask)
	content := container.NewVBox(btnOn, btnOff)
	w.SetContent(content)
	w.ShowAndRun()
}

func killtask() {
	c := exec.Command("taskkill", "/f", "/im", "explorer.exe")
	c.Stdout = os.Stdout
	c.Run()
}

func addtask() {
	c := exec.Command("explorer")
	c.Stdout = os.Stdout
	c.Run()
}
