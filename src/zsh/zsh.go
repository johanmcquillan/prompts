package zsh

import (
	"fmt"
)

type ZSHColour uint8

const (
	BLACK ZSHColour = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

type ZSHFormat struct {
	Colour *ZSHColour
	Bold   bool
}

func MakeZSHColour(colour ZSHColour) *ZSHFormat {
	c := colour
	return &ZSHFormat{
		Colour: &c,
	}
}

func (f *ZSHFormat) SetBold() *ZSHFormat {
	f.Bold = true
	return f
}

func (f *ZSHFormat) UnsetBold() *ZSHFormat {
	f.Bold = false
	return f
}

func (f *ZSHFormat) Format(text string) string {
	if f == nil {
		return text
	}

	output := "%{"
	if f.Bold {
		output += "%B"
	}

	if f.Colour != nil {
		output += "%F{" + fmt.Sprintf("%d", *f.Colour) + "}"
	}

	output += text

	if f.Colour != nil {
		output += "%f"
	}

	if f.Bold {
		output += "%b"
	}

	output += "%}"

	return output
}
