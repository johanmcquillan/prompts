package colors

import (
	"strconv"
	"strings"
)

func (f *ShellFormatter) zshFormat(text string) string {
	sb := &strings.Builder{}
	sb.WriteString("%{")
	if f.Bold {
		sb.WriteString("%B")
	}

	sb.WriteString("%F{")
	sb.WriteString(strconv.Itoa(int(f.Color)))
	sb.WriteString("}%}")

	sb.WriteString(text)

	sb.WriteString("%{%f")
	if f.Bold {
		sb.WriteString("%b")
	}
	sb.WriteString("%}")

	return sb.String()
}
