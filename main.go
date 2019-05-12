package main

import (
    "github.com/jessevdk/go-flags"
    "github.com/johanmcquillan/prompts/prompts"
)

const (
    envMachine     = "MACHINE"
    tsyncEnvVar    = "TERMINALSYNCSPACE"
    columnsEnvVar  = "COLUMNS"
    defaultColumns = 40
)

func main() {
    opts := prompts.Opts{}
    _, err := flags.Parse(&opts)
    if err != nil {
        panic(err)
    }

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
            prompts.
                MakeGitRelativeDirComponent().
                WithFormatter(prompts.MakeAnsiColour(69).SetBold())).
        WithComponent(
            prompts.
                MakeGitBranchComponent().
                WithFormatter(prompts.MakeAnsiColour(90).SetBold())).
        WithEnder(
            prompts.
                MakeBinaryEnder(
                    "$", nil,
                    "$", prompts.MakeAnsiColour(prompts.RED))).
        Print(opts.ExitCode)
}
