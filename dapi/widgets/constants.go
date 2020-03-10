package widgets

import "github.com/gdamore/tcell"

const (
	ColorBlack tcell.Color = iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
	ColorBrightBlack
	ColorBrightRed
	ColorBrightGreen
	ColorBrightYellow
	ColorBrightBlue
	ColorBrightMagenta
	ColorBrightCyan
	ColorBrightWhite
)

const (
	BackgroundColor = ColorBrightBlack
	BorderColor     = ColorBrightGreen
	TextColor       = ColorBrightCyan
	DangerColor     = ColorRed
	SuccessColor    = ColorGreen
)
