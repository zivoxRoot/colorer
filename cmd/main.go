package main

import (
	"fmt"

	"github.com/zivoxRoot/colorer/internal/colorer"
)

func main() {
	colors := colorer.NewColorer()

	// Colors
	fmt.Println(colors.Black() + "Black" + colors.Reset())
	fmt.Println(colors.BrightBlack() + "Bright black" + colors.Reset())

	fmt.Println(colors.Red() + "Red" + colors.Reset())
	fmt.Println(colors.BrightRed() + "Bright red" + colors.Reset())

	fmt.Println(colors.Green() + "Green" + colors.Reset())
	fmt.Println(colors.BrightGreen() + "BrightGreen" + colors.Reset())

	fmt.Println(colors.Yellow() + "Yellow" + colors.Reset())
	fmt.Println(colors.BrightYellow() + "Bright yellow" + colors.Reset())

	fmt.Println(colors.Blue() + "Blue" + colors.Reset())
	fmt.Println(colors.BrightBlue() + "Bright blue" + colors.Reset())

	fmt.Println(colors.Magenta() + "Magenta" + colors.Reset())
	fmt.Println(colors.BrightMagenta() + "Bright magenta" + colors.Reset())

	fmt.Println(colors.Cyan() + "Cyan" + colors.Reset())
	fmt.Println(colors.BrightCyan() + "Bright cyan" + colors.Reset())

	fmt.Println(colors.White() + "White" + colors.Reset())
	fmt.Println(colors.BrightWhite() + "Bright white" + colors.Reset())

	// Background colors
	fmt.Println(colors.BgBlack() + "Black" + colors.Reset())
	fmt.Println(colors.BrightBgBlack() + "Bright black" + colors.Reset())

	fmt.Println(colors.BgRed() + "Red" + colors.Reset())
	fmt.Println(colors.BrightBgRed() + "Bright red" + colors.Reset())

	fmt.Println(colors.BgGreen() + "Green" + colors.Reset())
	fmt.Println(colors.BrightBgGreen() + "Bright green" + colors.Reset())

	fmt.Println(colors.BgYellow() + "Yellow" + colors.Reset())
	fmt.Println(colors.BrightBgYellow() + "Bright yellow" + colors.Reset())

	fmt.Println(colors.BgBlue() + "Blue" + colors.Reset())
	fmt.Println(colors.BrightBgBlue() + "Bright blue" + colors.Reset())

	fmt.Println(colors.BgMagenta() + "Magenta" + colors.Reset())
	fmt.Println(colors.BrightBgMagenta() + "Bright magenta" + colors.Reset())

	fmt.Println(colors.BgCyan() + "Cyan" + colors.Reset())
	fmt.Println(colors.BrightBgCyan() + "Bright cyan" + colors.Reset())

	fmt.Println(colors.BgWhite() + "White" + colors.Reset())
	fmt.Println(colors.BrightBgWhite() + "Bright white" + colors.Reset())

	// Styles
	fmt.Println(colors.Bold() + "Bold" + colors.Reset())
	fmt.Println(colors.Strikethrough() + "Strikethrough" + colors.Reset())
	fmt.Println(colors.Dim() + "Dim" + colors.Reset())
	fmt.Println(colors.Italic() + "Italic" + colors.Reset())
	fmt.Println(colors.Underline() + "Underline" + colors.Reset())
}
