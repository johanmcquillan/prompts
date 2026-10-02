package colors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johanmcquillan/prompts"
)

func TestANSI(t *testing.T) {
	rawString := "Hello World"
	expectedLength := len(rawString)

	t.Run("NoFormatting", func(t *testing.T) {
		cmp := &prompts.StaticComponent{
			Value: rawString,
		}

		actualElement := cmp.MakeElement()
		expectedElement := prompts.Element{
			Output: rawString,
			Length: expectedLength,
		}

		assert.Equal(t, expectedElement, actualElement)
	})

	t.Run("Color", func(t *testing.T) {
		color := Color(1)

		cmp := &prompts.StaticComponent{
			Value: rawString,
			Formatter: &ShellFormatter{
				Type:  BASH,
				Color: color,
			},
		}

		expectedString := fmt.Sprintf(`\[\e[38;5;%dm\]%s\[\e[m\]`, color, rawString)
		expectedElement := prompts.Element{
			Output: expectedString,
			Length: expectedLength,
		}
		actualElement := cmp.MakeElement()

		assert.Equal(t, expectedElement, actualElement)
	})
	t.Run("RawANSI", func(t *testing.T) {
		color := Color(1)

		cmp := &prompts.StaticComponent{
			Value: rawString,
			Formatter: &ShellFormatter{
				Type:  RawANSI,
				Color: color,
				Bold:  true,
			},
		}

		expectedString := fmt.Sprintf("\x1b[1;38;5;%dm%s\x1b[m", color, rawString)
		expectedElement := prompts.Element{
			Output: expectedString,
			Length: expectedLength,
		}
		actualElement := cmp.MakeElement()

		assert.Equal(t, expectedElement, actualElement)
	})
}
