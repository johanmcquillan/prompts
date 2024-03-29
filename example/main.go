package main

import (
	"github.com/johanmcquillan/prompts"
	"github.com/johanmcquillan/prompts/colors"
	"github.com/johanmcquillan/prompts/env"
	"github.com/johanmcquillan/prompts/git"
	"github.com/johanmcquillan/prompts/k8s"
)

func main() {
	prompts.
		MakePrompt().
		WithComponent(
			env.MakeUserComponent().
				WithFormatter(colors.NewShellFormatter(88).SetBold(true))).
		WithComponent(
			git.MakeGitRelativeDirComponent().
				WithFormatter(colors.NewShellFormatter(69).SetBold(true))).
		WithComponent(
			git.MakeGitBranchComponent().
				WithFormatter(colors.NewShellFormatter(90).SetBold(true))).
		WithComponent(
			k8s.MakeK8sContextComponent().
				WithFormatter(colors.NewShellFormatter(22).SetBold(true))).
		WithEnder(
			prompts.MakeBinaryEnder(
				"$", nil,
				"$", colors.NewShellFormatter(colors.Red),
			),
		).
		ParseArgs().
		Print()
}
