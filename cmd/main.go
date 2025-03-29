package main

import (
	"fmt"

	"github.com/zivoxRoot/colorer/internal/colorer"
)

func main() {
	// Colors
	fmt.Println(colorer.Magenta() + "Magenta" + colorer.Reset())
	fmt.Println(colorer.BrightMagenta() + "Bright magenta" + colorer.Reset())
	fmt.Println(colorer.Black() + "Black" + colorer.Reset())
	fmt.Println(colorer.BrightBlack() + "Bright black" + colorer.Reset())
	fmt.Println(colorer.White() + "White" + colorer.Reset())
	fmt.Println(colorer.BrightWhite() + "Bright white" + colorer.Reset())
	fmt.Println(colorer.Cyan() + "Cyan" + colorer.Reset())
	fmt.Println(colorer.BrightCyan() + "Bright cyan" + colorer.Reset())

	// Styles
	fmt.Println(colorer.Underline() + "Underline" + colorer.Reset())
	fmt.Println(colorer.Italic() + "Italic" + colorer.Reset())
}
