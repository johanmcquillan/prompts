package main

import (
	"github.com/johanmcquillan/prompts/src/prompts"
	"github.com/johanmcquillan/prompts/src/prompts/colors"
	"github.com/johanmcquillan/prompts/src/prompts/env"
	"github.com/johanmcquillan/prompts/src/prompts/git"
	"github.com/johanmcquillan/prompts/src/prompts/k8s"
)

func main() {
	prompts.
		MakePrompt().
		WithComponent(
			env.MakeUserComponent().
				WithFormatter(colors.NewShellFormatter(88).SetBold(true))).
		WithSeparator().
		WithComponent(
			git.MakeGitRelativeDirComponent().
				WithFormatter(colors.NewShellFormatter(69).SetBold(true))).
		WithSeparator().
		WithComponent(
			git.MakeGitBranchComponent().
				WithFormatter(colors.NewShellFormatter(90).SetBold(true))).
		WithSeparator().
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
