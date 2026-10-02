package colors

import (
	"fmt"
	"strings"
)

const (
	ansiOpener    = `\[\e[`
	ansiCloser    = `\]`
	ansiReset     = `\[\e[m\]`
	ansiSeparator = ";"

	rawANSIOpener = "\x1b["
	rawANSIReset  = "\x1b[m"
)

func (f *ShellFormatter) ansiFormat( text string) string {
	if f == nil {
		return text
	}
	return f.ansiBegin() + text + ansiEnd()
}

func (f *ShellFormatter) rawANSIFormat(text string) string {
	if f == nil {
		return text
	}
	return rawANSIOpener + f.ansiCodes() + text + rawANSIReset
}

func (f *ShellFormatter) ansiBegin() string {
	return ansiOpener + f.ansiCodes() + ansiCloser
}

func (f *ShellFormatter) ansiCodes() string {
	var formats []string
	if f.Bold {
		formats = append(formats, "1")
	}
	formats = append(formats, fmt.Sprintf("38;5;%dm", f.Color))
	return strings.Join(formats, ansiSeparator)
}

func ansiEnd() string {
	return ansiReset
}
