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
)

func (f *ShellFormatter) ansiFormat( text string) string {
	if f == nil {
		return text
	}
	return f.ansiBegin() + text + ansiEnd()
}

func (f *ShellFormatter) ansiBegin() string {
	var formats []string
	if f.Bold {
		formats = append(formats, "1")
	}
	formats = append(formats, fmt.Sprintf("38;5;%dm", f.Color))
	return ansiOpener + strings.Join(formats, ansiSeparator) + ansiCloser
}

func ansiEnd() string {
	return ansiReset
}
