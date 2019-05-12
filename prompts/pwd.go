package prompts

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	homeSymbol = "~"
	upOne      = ".."

	pathSeparator = string(os.PathSeparator)
)

func MakeFullWDComponent() *EnvComponent {
	return MakeEnvComponent(envPWD)
}

func MakeRelativeWDComponent() *FunctionalComponent {
	return &FunctionalComponent{
		function: func() string {
			s, _ := relativeToHome()
			return s
		},
	}
}

func relativeToHome() (string, bool) {
	return substitutePathPrefix(os.Getenv(envHome), os.Getenv(envPWD), homeSymbol)
}

func substitutePathPrefix(prefixPath, fullPath, substitution string) (string, bool) {
	relativePath, err := filepath.Rel(prefixPath, fullPath)
	if err != nil || strings.Contains(relativePath, upOne) {
		return fullPath, false
	}

	return filepath.Join(substitution, relativePath), true
}
