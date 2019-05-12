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
			return substitutePathPrefix(os.Getenv(envHome), os.Getenv(envPWD), homeSymbol)
		},
	}
}

func substitutePathPrefix(prefixPath, fullPath, substitution string) string {
	relativePath, err := filepath.Rel(prefixPath, fullPath)
	if err != nil || strings.Contains(relativePath, upOne) {
		return fullPath
	}

	return filepath.Join(substitution, relativePath)
}
