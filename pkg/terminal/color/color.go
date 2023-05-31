package color

type Color string

var Reset Color = "\033[0m"
var Red Color = "\033[31m"
var Green Color = "\033[32m"

func Colorize(color Color, text string) string {
	return string(color) + text + string(Reset)
}
