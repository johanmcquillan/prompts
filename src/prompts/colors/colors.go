package colors

type ShellColor interface {
	isShellColor()
}

type ShellFormatter struct {
	Color ShellColor
	Bold  bool
}

func MakeShellFormatter(color ShellColor) *ShellFormatter {
	return &ShellFormatter{
		Color: color,
	}
}

func (f *ShellFormatter) SetBold(b bool) *ShellFormatter {
	f.Bold = b
	return f
}

func (f *ShellFormatter) Format(text string) string {
	if f == nil {
		return text
	}

	switch color := f.Color.(type) {
	case ANSIColor:
		return f.ansiFormat(color, text)
	case ZSHColor:
		return f.zshFormat(color, text)
	default:
		// Unknown color type.
		// We could panic here, but safer to just return the plain string for now.
		return text
	}
}
