package main

import (
	"github.com/johanmcquillan/prompts/src/git"
	"github.com/johanmcquillan/prompts/src/k8s"
	"github.com/johanmcquillan/prompts/src/prompts"
	"github.com/johanmcquillan/prompts/src/zsh"
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
				WithFormatter(zsh.MakeZSHColour(88).SetBold())).
		WithComponent(
			prompts.
				MakeEnvComponent(envMachine).
				WithFormatter(zsh.MakeZSHColour(214).SetBold())).
		WithComponent(
			git.
				MakeGitRelativeDirComponent().
				WithFormatter(zsh.MakeZSHColour(69).SetBold())).
		WithComponent(
			git.
				MakeGitBranchComponent().
				WithFormatter(zsh.MakeZSHColour(90).SetBold())).
		WithComponent(
			k8s.
				MakeK8sContextComponent().
				WithFormatter(zsh.MakeZSHColour(22).SetBold())).
		WithEnder(
			prompts.
				MakeBinaryEnder(
					`$`, nil,
					`$`, zsh.MakeZSHColour(zsh.Red))).
		ParseArgs().
		Print()
}
