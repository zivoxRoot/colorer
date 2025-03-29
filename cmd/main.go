package main

import (
	"fmt"

	"github.com/zivoxRoot/colorer/internal/colorer"
)

func main() {
	// Colors
	fmt.Println(colorer.Black() + "Black" + colorer.Reset())
	fmt.Println(colorer.BrightBlack() + "Bright black" + colorer.Reset())

	fmt.Println(colorer.Red() + "Red" + colorer.Reset())
	fmt.Println(colorer.BrightRed() + "Bright red" + colorer.Reset())

	fmt.Println(colorer.Green() + "Green" + colorer.Reset())
	fmt.Println(colorer.BrightGreen() + "BrightGreen" + colorer.Reset())

	fmt.Println(colorer.Yellow() + "Yellow" + colorer.Reset())
	fmt.Println(colorer.BrightYellow() + "Bright yellow" + colorer.Reset())

	fmt.Println(colorer.Blue() + "Blue" + colorer.Reset())
	fmt.Println(colorer.BrightBlue() + "Bright blue" + colorer.Reset())

	fmt.Println(colorer.Magenta() + "Magenta" + colorer.Reset())
	fmt.Println(colorer.BrightMagenta() + "Bright magenta" + colorer.Reset())

	fmt.Println(colorer.Cyan() + "Cyan" + colorer.Reset())
	fmt.Println(colorer.BrightCyan() + "Bright cyan" + colorer.Reset())

	fmt.Println(colorer.White() + "White" + colorer.Reset())
	fmt.Println(colorer.BrightWhite() + "Bright white" + colorer.Reset())

	// Styles
	fmt.Println(colorer.Bold() + "Bold" + colorer.Reset())
	fmt.Println(colorer.Dim() + "Dim" + colorer.Reset())
	fmt.Println(colorer.Italic() + "Italic" + colorer.Reset())
	fmt.Println(colorer.Underline() + "Underline" + colorer.Reset())
}
