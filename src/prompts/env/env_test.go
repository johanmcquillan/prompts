package env

import (
	"testing"

	"github.com/johanmcquillan/prompts/src/prompts"
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
		EnvUser: user,
		EnvHome: home,
		EnvPWD:  fullDir,
	}

	t.Run("Vars", func(t *testing.T) {
		for envVar := range env {
			t.Run("$"+envVar, func(t *testing.T) {
				cmp := &EnvComponent{
					envFetcher: env,
					envVar:     envVar,
				}

				actualElement := cmp.MakeElement()
				expectedElement := prompts.Element{
					Output: env[envVar],
					Length: len(env[envVar]),
				}

				assert.Equal(t, expectedElement, actualElement)
			})
		}
	})
}
