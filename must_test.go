package must_test

import (
	"errors"
	"fmt"
	"github.com/KonradKuznicki/must"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMust_doNothingWhenNoError(t *testing.T) {

	must.Must(nil)
}

func TestMust_FatalWhenError(t *testing.T) {
	fatalCalled := false
	must.SetFatalHandler(func(format string, v ...any) {
		fatalCalled = true
	})

	must.Must(errors.New("test"))

	assert.True(t, fatalCalled)
}

func TestMustF_doNothingWhenNoError(t *testing.T) {
	must.MustF(nil, "")
	must.MustF(nil, "", "asdf", "asdf")
}

func TestMustF_FatalIsFormatted(t *testing.T) {
	fatalCalled := ""
	must.SetFatalHandler(func(format string, v ...any) {
		fatalCalled = fmt.Sprintf(format, v...)
	})

	must.MustF(errors.New("error"), "message: %v, %v, %v", "test1", 2)

	assert.Equal(t, "message: error, test1, 2", fatalCalled)
}

func TestReturn_PassesValueWhenNoError(t *testing.T) {
	r := must.Return(1, nil)
	assert.Equal(t, 1, r)

	r2 := must.Return("test", nil)
	assert.Equal(t, "test", r2)
}

func TestReturn_FatalWhenError(t *testing.T) {
	fatalCalled := false
	must.SetFatalHandler(func(format string, v ...any) {
		fatalCalled = true
	})

	_ = must.Return(0, errors.New("test"))

	assert.True(t, fatalCalled)
}

func TestReturnF_PassesValueWhenNoError(t *testing.T) {
	r := must.ReturnF(1, nil)("")
	assert.Equal(t, 1, r)

	r2 := must.ReturnF("test", nil)("")
	assert.Equal(t, "test", r2)
}

func TestReturnF_FatalIsFormatted(t *testing.T) {
	fatalCalled := ""
	must.SetFatalHandler(func(format string, v ...any) {
		fatalCalled = fmt.Sprintf(format, v...)
	})

	_ = must.ReturnF(0, errors.New("error"))("message: %v, %v, %v", "test1", 2)

	assert.Equal(t, "message: error, test1, 2", fatalCalled)
}
