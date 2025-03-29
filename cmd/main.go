package main

import (
	"fmt"

	"colorer/internal/colorer"
)

func main() {
	// Colors
	fmt.Println("Je suis le meilleur " + colorer.Magenta() + "Salut a tous" + colorer.Reset())
	fmt.Println("Je suis le meilleur " + colorer.BrightMagenta() + "Salut a tous" + colorer.Reset())
	fmt.Println("Je suis le meilleur " + colorer.White() + "Salut a tous" + colorer.Reset())
	fmt.Println("Je suis le meilleur " + colorer.BrightWhite() + "Salut a tous" + colorer.Reset())
	fmt.Println("Je suis le meilleur " + colorer.Cyan() + "Salut a tous" + colorer.Reset())
	fmt.Println("Je suis le meilleur " + colorer.BrightCyan() + "Salut a tous" + colorer.Reset())

	// Styles
	fmt.Println("Je suis le meilleur " + colorer.Underline() + "Salut a tous" + colorer.Reset())
	fmt.Println("Je suis le meilleur " + colorer.Italic() + "Salut a tous" + colorer.Reset())
}
