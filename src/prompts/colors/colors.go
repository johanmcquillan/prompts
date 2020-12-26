package colors

type Color uint8

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGrey
	DarkGrey
	White = 15
)

type ShellFormatter struct {
	Type  ShellType
	Color Color
	Bold  bool
}

func NewShellFormatter(color Color) *ShellFormatter {
	// This is imperfect as it oly fetches the default shell, not the current shell.
	return &ShellFormatter{
		Type:  getShellType(),
		Color: color,
	}
}

func (f *ShellFormatter) SetBold(b bool) *ShellFormatter {
	f.Bold = b
	return f
}

func (f *ShellFormatter) SetShellType(shellType ShellType) *ShellFormatter {
	f.Type = shellType
	return f
}

func (f *ShellFormatter) Format(text string) string {
	if f == nil {
		return text
	}

	switch f.Type {
	case BASH:
		return f.ansiFormat(text)
	case ZSH:
		return f.zshFormat(text)
	default:
		// Unknown color type.
		// We could panic here, but safer to just return the plain string for now.
		return text
	}
}
