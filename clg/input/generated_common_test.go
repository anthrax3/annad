package input

// This file is generated by the CLG generator. Don't edit it manually. The CLG
// generator is invoked by go generate. For more information about the usage of
// the CLG generator check https://github.com/xh3b4sd/clggen or have a look at
// the clg package. There is the go generate statement placed to invoke clggen.

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_CLG_filterError_noValues(t *testing.T) {
	out, err := filterError(nil)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(out) != 0 {
		t.Fatal("expected", 2, "got", len(out))
	}
}

func Test_CLG_filterError_noError(t *testing.T) {
	in := []reflect.Value{
		reflect.ValueOf("one"),
		reflect.ValueOf("two"),
		reflect.ValueOf("three"),
	}

	out, err := filterError(in)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(out) != 2 {
		t.Fatal("expected", 2, "got", len(out))
	}
}

func Test_CLG_filterError_nilError(t *testing.T) {
	var err error
	in := []reflect.Value{
		reflect.ValueOf("one"),
		reflect.ValueOf("two"),
		reflect.ValueOf(err),
	}

	out, err := filterError(in)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(out) != 2 {
		t.Fatal("expected", 2, "got", len(out))
	}
}

func Test_CLG_filterError_errgoError(t *testing.T) {
	in := []reflect.Value{
		reflect.ValueOf("one"),
		reflect.ValueOf("two"),
		reflect.ValueOf(invalidConfigError),
	}

	_, err := filterError(in)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_CLG_filterError_fmtError(t *testing.T) {
	fmtErr := fmt.Errorf("test error")
	in := []reflect.Value{
		reflect.ValueOf("one"),
		reflect.ValueOf("two"),
		reflect.ValueOf(fmtErr),
	}

	_, err := filterError(in)
	if err == nil {
		t.Fatal("expected", fmtErr, "got", nil)
	}
	msg := err.Error()
	if msg != "test error" {
		t.Fatal("expected", "test error", "got", msg)
	}
}

func Test_CLG_filterError_CLG(t *testing.T) {
	in := []reflect.Value{
		reflect.ValueOf("one"),
		reflect.ValueOf("two"),
		reflect.ValueOf(MustNew()),
	}

	_, err := filterError(in)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}
