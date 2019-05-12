package prompts

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

				actualElement := cmp.MakeElement()
				expectedElement := Element{env[envVar], len(env[envVar])}

				assert.Equal(t, expectedElement, actualElement)
			})
		}
	})
}

