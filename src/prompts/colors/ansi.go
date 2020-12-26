package colors

import (
	"fmt"
	"strings"
)

const (
	ansiOpener    = `\e[`
	ansiReset     = `\e[0m`
	ansiSeparator = ";"
)

type ANSIColor uint8

const (
	ANSIBlack ANSIColor = iota
	ANSIRed
	ANSIGreen
	ANSIYellow
	ANSIBlue
	ANSIMagenta
	ANSICyan
	ANSILightGrey
	ANSIDarkGrey
	ANSIWhite = 15
)

func (ANSIColor) isShellColor() {}

func (f *ShellFormatter) ansiFormat(color ANSIColor, text string) string {
	if f == nil {
		return text
	}
	return f.ansiBegin(color) + text + ansiEnd()
}

func (f *ShellFormatter) ansiBegin(color ANSIColor) string {
	var formats []string
	if f.Bold {
		formats = append(formats, "1")
	}
	formats = append(formats, fmt.Sprintf("38;5;%dm", color))
	return ansiOpener + strings.Join(formats, ansiSeparator)
}

func  ansiEnd() string {
	return ansiReset
}
