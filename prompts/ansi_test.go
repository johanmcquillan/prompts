package prompts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColourise(t *testing.T) {
	cmp := &StaticComponent{
		Value: "Hello World",
	}

	expected := "Hello World"
	actualString, actualLength := cmp.StringAndLength()

	assert.Equal(t, expected, actualString)
	assert.Equal(t, len(expected), actualLength)
}
