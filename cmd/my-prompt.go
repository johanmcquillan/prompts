package main

import (
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
                WithFormatter(prompts.MakeAnsiColour(88).SetBold())).
        WithComponent(
            prompts.
                MakeEnvComponent(envMachine).
                WithFormatter(prompts.MakeAnsiColour(214).SetBold())).
        WithComponent(
            git.
                MakeGitRelativeDirComponent().
                WithFormatter(prompts.MakeAnsiColour(69).SetBold())).
        WithComponent(
            git.
                MakeGitBranchComponent().
                WithFormatter(prompts.MakeAnsiColour(90).SetBold())).
        WithComponent(
            k8s.
                MakeK8sContextComponent().
                WithFormatter(prompts.MakeAnsiColour(22).SetBold())).
        WithEnder(
            prompts.
                MakeBinaryEnder(
                    "$", nil,
                    "$", prompts.MakeAnsiColour(prompts.RED))).
        ParseArgs().
        Print()
}
