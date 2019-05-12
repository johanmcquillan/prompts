package prompts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnsi(t *testing.T) {
	rawString := "Hello World"
	expectedLength := len(rawString)

	t.Run("NoFormatting", func(t *testing.T) {
		cmp := &StaticComponent{
			Value: rawString,
		}

		actualElement := cmp.MakeElement()
		expectedElement := Element{rawString, expectedLength}

		assert.Equal(t, expectedElement, actualElement)
	})

	t.Run("Colour", func(t *testing.T) {
		colour := 1

		cmp := &StaticComponent{
			Value: rawString,
			Formatter: &AnsiFormat{
				Colour: AnsiColour(colour),
			},
		}

		expectedString := fmt.Sprintf(`\[\e[38;5;%dm\]%s\[\e[m\]`, colour, rawString)
		expectedElement := Element{expectedString, expectedLength}
		actualElement := cmp.MakeElement()

		assert.Equal(t, expectedElement, actualElement)
	})
}
