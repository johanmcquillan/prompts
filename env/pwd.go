package env

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/johanmcquillan/prompts"
)

const (
	homeSymbol = "~"
	upOne      = ".."

	PathSeparator = string(os.PathSeparator)
)

func MakeFullWDComponent() *EnvComponent {
	return MakeEnvComponent(EnvPWD)
}

func MakeRelativeWDComponent() *prompts.FunctionalComponent {
	return &prompts.FunctionalComponent{
		Function: func() string {
			s, _ := RelativeToHome()
			return s
		},
	}
}

//func MakeRelativeWDDynamicComponent() *DynamicComponent {
//	return &DynamicComponent{
//		Function: func(p *PromptState) string {
//
//		},
//	}
//}

func RelativeToHome() (string, bool) {
	return SubstitutePathPrefix(os.Getenv(EnvHome), os.Getenv(EnvPWD), homeSymbol)
}

func SubstitutePathPrefix(prefixPath, fullPath, substitution string) (string, bool) {
	relativePath, err := filepath.Rel(prefixPath, fullPath)
	if err != nil || strings.Contains(relativePath, upOne) {
		return fullPath, false
	}

	return filepath.Join(substitution, relativePath), true
}
