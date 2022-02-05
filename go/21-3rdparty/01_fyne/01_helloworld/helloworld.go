package main

import (
  "fyne.io/fyne/app"
  "fyne.io/fyne/widget"
)

func main() {
  a := app.New()
  window := a.NewWindow("Hello World")
  window.SetContent(widget.NewLabel("Hello World!"))
  window.ShowAndRun()
}

