package main

import "github.com/gonutz/prototype/draw"

func main() {

	title := "Capture mouse input"
	height := 300
	width := 300

	rectXCord := 0
	rectYCord := 0
	rectHeight := 30
	rectWidth := 60

	draw.RunWindow(title, height, width, func(window draw.Window) {

		rectXCord, rectYCord = window.MousePosition()

		rectXCord = rectXCord - 30
		rectYCord = rectYCord - 15

		window.FillRect(rectXCord, rectYCord, rectWidth, rectHeight, draw.Cyan)

	})
}
