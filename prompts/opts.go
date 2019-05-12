package prompts

type Opts struct {
	ExitCode    int    `short:"c" long:"exit-code" default:"0" description:"Exit code of previous command"`
	ShowAll     bool   `short:"a" long:"show-all" description:"Show all elements, even if empty"`
	FallBack    string `long:"fallback" description:"Upon a panic, use this prompt instead"`
	NoRecover   bool   `long:"no-recover" description:"Upon a panic, do not recover. Overrides 'FallBack'"`
}
