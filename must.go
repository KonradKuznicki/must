package must

import "log"

type FatalHandler func(format string, v ...any)

var fatalf FatalHandler

// SetFatalHandler it is designed for tests primarily
func SetFatalHandler(handler FatalHandler) {
	fatalf = handler
}

func init() {
	SetFatalHandler(log.Fatalf)
}

func ReturnF[K any](r K, err error) func(format string, v ...any) K {
	return func(format string, v ...any) K {
		if err != nil {
			fatalArgs := append([]any{err}, v...)
			fatalf(format, fatalArgs...)
		}
		return r
	}
}

func Return[K any](r K, err error) K {
	return ReturnF[K](r, err)("%v")
}

func MustF(err error, format string, v ...any) {
	_ = ReturnF(0, err)(format, v...)
}

func Must(err error) {
	_ = Return(0, err)
}
