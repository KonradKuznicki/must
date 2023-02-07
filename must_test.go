package must_test

import (
	"errors"
	"fmt"
	"github.com/KonradKuznicki/must"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMust_doNothingWhenNoError(t *testing.T) {

	must.Must0(nil)
}

func TestMust_FatalWhenError(t *testing.T) {
	fatalCalled := false
	must.SetFatalHandler(func(format string, v ...any) {
		fatalCalled = true
	})

	must.Must0(errors.New("test"))

	assert.True(t, fatalCalled)
}

func TestMustF_doNothingWhenNoError(t *testing.T) {
	must.Must0f(nil, "")
	must.Must0f(nil, "", "asdf", "asdf")
}

func TestMustF_FatalIsFormatted(t *testing.T) {
	fatalCalled := ""
	must.SetFatalHandler(func(format string, v ...any) {
		fatalCalled = fmt.Sprintf(format, v...)
	})

	must.Must0f(errors.New("error"), "message: %v, %v, %v", "test1", 2)

	assert.Equal(t, "message: error, test1, 2", fatalCalled)
}

func TestReturn_PassesValueWhenNoError(t *testing.T) {
	r := must.Must(1, nil)
	assert.Equal(t, 1, r)

	r2 := must.Must("test", nil)
	assert.Equal(t, "test", r2)
}

func TestReturn_FatalWhenError(t *testing.T) {
	fatalCalled := false
	must.SetFatalHandler(func(format string, v ...any) {
		fatalCalled = true
	})

	_ = must.Must(0, errors.New("test"))

	assert.True(t, fatalCalled)
}

func TestReturnF_PassesValueWhenNoError(t *testing.T) {
	r := must.Mustf(1, nil)("")
	assert.Equal(t, 1, r)

	r2 := must.Mustf("test", nil)("")
	assert.Equal(t, "test", r2)
}

func TestReturnF_FatalIsFormatted(t *testing.T) {
	fatalCalled := ""
	must.SetFatalHandler(func(format string, v ...any) {
		fatalCalled = fmt.Sprintf(format, v...)
	})

	_ = must.Mustf(0, errors.New("error"))("message: %v, %v, %v", "test1", 2)

	assert.Equal(t, "message: error, test1, 2", fatalCalled)
}
