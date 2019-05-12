package prompts

type Opts struct {
	ExitCode int `short:"e" long:"exit-code" default:"0" description:"Exit code of previous command"`
}
