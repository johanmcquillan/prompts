package prompts

import (
	"fmt"
	"strings"
)

const (
	ansiOpener = `\[\e[`
	ansiCloser = `\]`
	ansiReset = `\[\e[m\]`
	ansiSeparator = ";"
)

type AnsiColour int

const (
	Red AnsiColour = 1
)

type AnsiFormat struct {
	Colour AnsiColour
	Bold bool
	Italic bool
}

func (a *AnsiFormat) Colourise(text string) string {
	if a == nil {
		return text
	}
	return a.begin() + text + a.end()
}

func (a *AnsiFormat) begin() string {
	var formats []string
	if a.Bold {
		formats = append(formats, "1")
	}
	if a.Colour > 0 {
		formats = append(formats, fmt.Sprintf("38;5;%dm", a.Colour))
	}
	return ansiOpener + strings.Join(formats, ansiSeparator) + ansiCloser
}

func (a *AnsiFormat) end() string {
	return ansiReset
}
