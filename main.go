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
        WithUser().
        WithEnvVar(envMachine).
        WithRelativeWorkingDir().
        WithGitBranch()

    fmt.Print(prompt)
}
