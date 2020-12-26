package colors

import (
	"fmt"
)

func (f *ShellFormatter) zshFormat(text string) string {
	output := "%{"
	if f.Bold {
		output += "%B"
	}

	output += "%F{" + fmt.Sprintf("%d", f.Color) + "}"

	output += text

	output += "%f"

	if f.Bold {
		output += "%b"
	}

	output += "%}"

	return output
}
