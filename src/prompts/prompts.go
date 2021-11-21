package prompts

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

const (
	defaultEnder     = "$"
	defaultSeparator = ":"
	defaultPanicMsg  = "<Prompt panicked!>"

	// Upon a panic, this prompt will be used instead
	defaultFallBack = defaultPanicMsg + defaultEnder + " "
)

// Prompt contains all the components and configuration for a prompt line.
type Prompt struct {
	PromptState
	Ender
	Components []Component
	Separator  string
}

type PromptState struct {
	Opts
	CurrentLength int
}

// MakePrompt initialises a standard Prompt
func MakePrompt() *Prompt {
	return &Prompt{
		Ender:     MakeStaticEnder(defaultEnder),
		Separator: defaultSeparator,
	}
}

// Print prints the prompt string to the terminal, using a default exit code of 0
func (p *Prompt) Print() string {
	s := p.String(p.ExitCode)
	fmt.Print(s)
	return s
}

// Print prints the prompt string to the terminal
func (p *Prompt) PrintWithExitCode(exitCode int) string {
	s := p.String(exitCode)
	fmt.Print(s)
	return s
}

// String returns the prompt string, including formatting escape characters
func (p *Prompt) String(exitCode int) (output string) {
	defer func() {
		if p.NoRecover {
			return
		}

		// Recover and return a fall back prompt
		if recover() != nil {
			output = p.getFallBack()
		}
	}()

	elementStrings := make([]string, 0, len(p.Components))
	for _, component := range p.Components {
		element := component.MakeElement()
		if element.Length == 0 {
			// Skip 0 length elements
			continue
		}
		elementStrings = append(elementStrings, element.Output)
		p.CurrentLength += element.Length
	}

	output = strings.Join(elementStrings, p.Separator)
	output += p.End(exitCode).Output + " "
	return output
}

func (p *Prompt) WithComponent(component Component) *Prompt {
	p.Components = append(p.Components, component)
	return p
}

func (p *Prompt) SetSeparator(separator string) *Prompt {
	p.Separator = separator
	return p
}

func (p *Prompt) WithEnder(e Ender) *Prompt {
	p.Ender = e
	return p
}

func (p *Prompt) ParseArgs() *Prompt {
	defer p.recoverFromHelpError()

	p.Opts = Opts{}
	_, err := flags.Parse(&p.Opts)
	if err != nil {
		panic(err)
	}

	return p
}

func (p *Prompt) WithArgs(args []string) *Prompt {
	defer p.recoverFromHelpError()

	p.Opts = Opts{}
	_, err := flags.ParseArgs(&p.Opts, args)
	if err != nil {
		panic(err)
	}

	return p
}

func (p *Prompt) WithOpts(opts Opts) *Prompt {
	p.Opts = opts
	return p
}

func (p *Prompt) getFallBack() string {
	if p.FallBack == "" {
		return defaultFallBack
	}
	return p.FallBack
}

func (p *Prompt) recoverFromHelpError() {
	if p.NoRecover {
		return
	}

	if r := recover(); r != nil {
		if err, ok := r.(*flags.Error); ok && err.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			panic(r)
		}
	}
}
