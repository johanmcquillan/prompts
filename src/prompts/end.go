package prompts

type StaticEnder struct {
	Formatter
	Symbol string
}

type FunctionalEnder struct {
	Function func(int) (string, Formatter)
}

func MakeStaticEnder(symbol string) *StaticEnder {
	return &StaticEnder{Symbol: symbol}
}

func (e *StaticEnder) WithFormatter(f Formatter) *StaticEnder {
	e.Formatter = f
	return e
}

func MakeFunctionalEnder(f func(int) (string, Formatter)) *FunctionalEnder {
	return &FunctionalEnder{Function: f}
}

func (e *StaticEnder) End(int) Element {
	if e.Formatter == nil {
		return Element{
			Output: e.Symbol,
			Length: len(e.Symbol),
		}
	}
	return Element{
		Output: e.Format(e.Symbol),
		Length: len(e.Symbol),
	}
}

func (e *FunctionalEnder) End(exitCode int) Element {
	s, f := e.Function(exitCode)

	if f == nil {
		return Element{
			Output: s,
			Length: len(s),
		}
	}

	return Element{
		Output: f.Format(s),
		Length: len(s),
	}
}

func MakeBinaryEnder(successSymbol string, successFormat Formatter, failSymbol string, failFormat Formatter) *FunctionalEnder {
	return &FunctionalEnder{
		Function: func(exitCode int) (string, Formatter) {
			if exitCode == 0 {
				return successSymbol, successFormat
			}
			return failSymbol, failFormat
		},
	}
}
