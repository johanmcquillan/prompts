package ansi

import (
	"fmt"
	"strings"
)

const (
	ansiOpener    = `\[\e[`
	ansiCloser    = `\]`
	ansiReset     = `\[\e[m\]`
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

func (f *ANSIFormat) SetBold() *ANSIFormat {
	f.Bold = true
	return f
}

func (f *ANSIFormat) UnsetBold() *ANSIFormat {
	f.Bold = false
	return f
}

func (f *ANSIFormat) Format(text string) string {
	if f == nil {
		return text
	}
	return f.begin() + text + f.end()
}

func (f *ANSIFormat) begin() string {
	var formats []string
	if f.Bold {
		formats = append(formats, "1")
	}
	if f.Colour != nil {
		formats = append(formats, fmt.Sprintf("38;5;%dm", *f.Colour))
	}
	return ansiOpener + strings.Join(formats, ansiSeparator) + ansiCloser
}

func (f *ANSIFormat) end() string {
	return ansiReset
}
