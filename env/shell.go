package env

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/johanmcquillan/prompts"
)

type ShellComponent struct {
	envFetcher
	prompts.Formatter
	envVar string
}

func (c *ShellComponent) GetShell() string {
	return GetShell()
}

func MakeShellComponent() *ShellComponent {
	return &ShellComponent{
		envFetcher: Env{},
	}
}

func (c *ShellComponent) WithFormatter(formatter prompts.Formatter) *ShellComponent {
	c.Formatter = formatter
	return c
}

func (c *ShellComponent) MakeElement() prompts.Element {
	rawValue := c.GetShell()

	if c.Formatter == nil {
		return prompts.Element{
			Output: rawValue,
			Length: len(rawValue),
		}
	}

	return prompts.Element{
		Output: c.Format(rawValue),
		Length: len(rawValue),
	}
}

func GetShell() string {

	output, err := exec.Command("ps", "-o", "command=", "-p", fmt.Sprintf("%d", os.Getppid())).Output()
	if err != nil {
		return ""
	}
	return filepath.Base(strings.Trim(strings.TrimSpace(string(output)), "-"))
}
