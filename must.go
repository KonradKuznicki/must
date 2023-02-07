package must

import (
	"log"
)

type FatalHandler func(format string, v ...any)

var fatalf FatalHandler

// SetFatalHandler it is designed for tests primarily
func SetFatalHandler(handler FatalHandler) {
	fatalf = handler
}

func init() {
	SetFatalHandler(log.Fatalf)
}

func Must0f(err error, format string, v ...any) {
	if err != nil {
		fatalArgs := append([]any{err}, v...)
		fatalf(format, fatalArgs...)
	}
}

func Must0(err error) {
	Must0f(err, "%v")
}

func Mustf[K any](r K, err error) func(format string, v ...any) K {
	return func(format string, v ...any) K {
		Must0f(err, format, v...)
		return r
	}
}

func Must[K any](r K, err error) K {

	return Mustf[K](r, err)("%v")
}
