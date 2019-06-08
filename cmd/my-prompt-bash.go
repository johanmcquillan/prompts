package main

import (
	"github.com/johanmcquillan/prompts/src/ansi"
	"github.com/johanmcquillan/prompts/src/git"
	"github.com/johanmcquillan/prompts/src/k8s"
	"github.com/johanmcquillan/prompts/src/prompts"
)

const (
	envMachine     = "MACHINE"
	tsyncEnvVar    = "TERMINALSYNCSPACE"
	columnsEnvVar  = "COLUMNS"
	defaultColumns = 40
)

func main() {
	prompts.
		MakePrompt().
		WithComponent(
			prompts.
				MakeUserComponent().
				WithFormatter(ansi.MakeANSIColour(88).SetBold())).
		WithComponent(
			prompts.
				MakeEnvComponent(envMachine).
				WithFormatter(ansi.MakeANSIColour(214).SetBold())).
		WithComponent(
			git.
				MakeGitRelativeDirComponent().
				WithFormatter(ansi.MakeANSIColour(69).SetBold())).
		WithComponent(
			git.
				MakeGitBranchComponent().
				WithFormatter(ansi.MakeANSIColour(90).SetBold())).
		WithComponent(
			k8s.
				MakeK8sContextComponent().
				WithFormatter(ansi.MakeANSIColour(22).SetBold())).
		WithEnder(
			prompts.
				MakeBinaryEnder(
					"$", nil,
					"$", ansi.MakeANSIColour(ansi.RED))).
		ParseArgs().
		Print()
}
