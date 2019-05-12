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

type AnsiColour uint8

const (
	BLACK   AnsiColour = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	LIGHT_GREY
	DARK_GREY
	WHITE           = 15
)

type AnsiFormat struct {
	Colour *AnsiColour
	Bold   bool
	Italic bool
}

func MakeAnsiColour(colour AnsiColour) *AnsiFormat {
	c := colour
	return &AnsiFormat{
		Colour: &c,
	}
}

func (a *AnsiFormat) SetBold() *AnsiFormat {
	a.Bold = true
	return a
}

func (a *AnsiFormat) UnsetBold() *AnsiFormat {
	a.Bold = false
	return a
}

func (a *AnsiFormat) Format(text string) string {
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
	if a.Colour != nil {
		formats = append(formats, fmt.Sprintf("38;5;%dm", *a.Colour))
	}
	return ansiOpener + strings.Join(formats, ansiSeparator) + ansiCloser
}

func (a *AnsiFormat) end() string {
	return ansiReset
}
