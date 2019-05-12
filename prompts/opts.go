package prompts

type Opts struct {
	ExitCode  int  `short:"c" long:"exit-code" default:"0" description:"Exit code of previous command"`
	ShowAll   bool `short:"a" long:"show-all" description:"Show all elements, even if empty"`
}
