package main

import (
    "fmt"
    "prompts"
)

const (
    tsyncEnvVar = "TERMINALSYNCSPACE"
    columnsEnvVar = "COLUMNS"
    defaultColumns = 40
)

func main() {
    prompt := prompts.
        MakePrompt().
        WithUser()

    fmt.Print(prompt)
}
