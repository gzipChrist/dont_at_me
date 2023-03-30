package style

import "strings"

const MaxCharWidth = 16
const Reset = "\033[0m"

type Color int

const (
	Cyan Color = iota + 1
	Green
	Red
	Dim
)

var ColorAnsiCodes = map[Color]string{
	Cyan:  "\033[36m",
	Green: "\033[32m",
	Red:   "\033[31m",
	Dim:   "\033[2m",
}

func (c Color) Ansi() string {
	return ColorAnsiCodes[c]
}

func (c Color) Colorize(s string) string {
	b := strings.Builder{}
	b.WriteString(c.Ansi())
	b.WriteString(s)
	b.WriteString(Reset)

	return b.String()
}
