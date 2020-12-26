package colors

import (
	"fmt"
)

type ZSHColor uint8

const (
	ZSHBlack ZSHColor = iota
	ZSHRed
	ZSHGreen
	ZSHYellow
	ZSHBlue
	ZSHMagenta
	ZSHCyan
	ZSHWhite
)

func (ZSHColor) isShellColor() {}

func (f *ShellFormatter) zshFormat(color ZSHColor, text string) string {
	output := "%{"
	if f.Bold {
		output += "%B"
	}

	output += "%F{" + fmt.Sprintf("%d", color) + "}"

	output += text

	output += "%f"

	if f.Bold {
		output += "%b"
	}

	output += "%}"

	return output
}
