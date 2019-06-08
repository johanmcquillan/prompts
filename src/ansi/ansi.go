package ansi

import (
	"fmt"
	"strings"
)

const (
	ansiOpener    = `\033[`
	ansiReset     = `\033[m`
	ansiSeparator = ";"
)

type ANSIColour uint8

const (
	Black ANSIColour = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGrey
	DarkGrey
	White = 15
)

type ANSIFormat struct {
	Colour *ANSIColour
	Bold   bool
	Italic bool
}

func MakeANSIColour(colour ANSIColour) *ANSIFormat {
	c := colour
	return &ANSIFormat{
		Colour: &c,
	}
}

func (a *ANSIFormat) SetBold() *ANSIFormat {
	a.Bold = true
	return a
}

func (a *ANSIFormat) UnsetBold() *ANSIFormat {
	a.Bold = false
	return a
}

func (a *ANSIFormat) Format(text string) string {
	if a == nil {
		return text
	}
	return a.begin() + text + a.end()
}

func (a *ANSIFormat) begin() string {
	var formats []string
	if a.Bold {
		formats = append(formats, "1")
	}
	if a.Colour != nil {
		formats = append(formats, fmt.Sprintf("38;5;%dm", *a.Colour))
	}
	return ansiOpener + strings.Join(formats, ansiSeparator)
}

func (a *ANSIFormat) end() string {
	return ansiReset
}
