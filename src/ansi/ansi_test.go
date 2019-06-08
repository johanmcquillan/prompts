package ansi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johanmcquillan/prompts/src/prompts"
)

func TestAnsi(t *testing.T) {
	rawString := "Hello World"
	expectedLength := len(rawString)

	t.Run("NoFormatting", func(t *testing.T) {
		cmp := &prompts.StaticComponent{
			Value: rawString,
		}

		actualElement := cmp.MakeElement()
		expectedElement := prompts.Element{rawString, expectedLength}

		assert.Equal(t, expectedElement, actualElement)
	})

	t.Run("Colour", func(t *testing.T) {
		colour := ANSIColour(1)

		cmp := &prompts.StaticComponent{
			Value: rawString,
			Formatter: &ANSIFormat{
				Colour: &colour,
			},
		}

		expectedString := fmt.Sprintf(`\[\e[38;5;%dm\]%s\[\e[m\]`, colour, rawString)
		expectedElement := prompts.Element{expectedString, expectedLength}
		actualElement := cmp.MakeElement()

		assert.Equal(t, expectedElement, actualElement)
	})
}
