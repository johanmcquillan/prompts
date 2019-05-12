package main

import (
    "fmt"
    "prompts"
)

const (
    envMachine     = "MACHINE"
    tsyncEnvVar    = "TERMINALSYNCSPACE"
    columnsEnvVar  = "COLUMNS"
    defaultColumns = 40
)

func main() {
    prompt := prompts.
        MakePrompt().
        WithComponent(prompts.MakeUserComponent()).
        WithEnvVar(envMachine).
        WithComponent(prompts.MakeRelativeWDComponent()).
        WithComponent(prompts.MakeGitBranchComponent())

    fmt.Print(prompt)
}
