package prompts

type Opts struct {
	Args `group:"Application Arguments"`
	Config `group:"Configuration Options"`
}

type Args struct {
	ExitCode     int `short:"c" long:"exit-code" default:"0" description:"Exit code of previous command"`
	//TargetLength int `short:"l" `
}

type Config struct {
	FallBack  string `long:"fallback" description:"Upon a panic, use this prompt instead"`
	NoRecover bool   `long:"no-recover" description:"Upon a panic, do not recover. Overrides 'FallBack'"`
}
