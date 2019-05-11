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

		actualString, actualLength := cmp.StringAndLength()

		assert.Equal(t, rawString, actualString)
		assert.Equal(t, expectedLength, actualLength)
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
		actualString, actualLength := cmp.StringAndLength()
		assert.Equal(t, expectedString, actualString)
		assert.Equal(t, expectedLength, actualLength)
	})
}
