package k8s

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/johanmcquillan/prompts/src/prompts"
)

func MakeK8sContextComponent() *prompts.FunctionalComponent {
	return &prompts.FunctionalComponent{
		Function: k8sContext,
	}
}

func k8sContext() string {
	var out bytes.Buffer
	cmd := exec.Command("kubectl", "config", "current-context")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return ""
	}
	return strings.Replace(out.String(), "\n", "", -1)
}
