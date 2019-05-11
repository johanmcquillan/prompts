package prompts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testEnvFetcher map[string]string

func (e testEnvFetcher) GetEnv(envVar string) string {
	return e[envVar]
}

func TestEnv(t *testing.T) {
	user := "johan"
	home := "/home/johan"
	fullDir := "/home/johan/test/directory"
	//relDir := "~/test/directory"

	env := testEnvFetcher{
		envUser: user,
		envHome: home,
		envPWD:  fullDir,
	}

	t.Run("Vars", func(t *testing.T) {
		for envVar := range env {
			t.Run("$" + envVar, func(t *testing.T) {
				cmp := &EnvComponent{
					envFetcher: env,
					envVar:     envVar,
				}

				s, l := cmp.StringAndLength()
				actualString, ok := env[envVar]

				if assert.True(t, ok) {
					assert.Equal(t, actualString, s)
					assert.Equal(t, len(actualString), l)
				}
			})
		}
	})
}

