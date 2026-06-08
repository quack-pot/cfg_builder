package tests

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func Assert(
	t *testing.T,
	condition bool,
	err_format string,
	err_args ...any,
) {
	t.Helper()

	if !condition {
		err_msg := fmt.Sprintf(err_format, err_args...)
		t.Fatalf("%s\n\nStack Trace:\n%s", err_msg, debug.Stack())
	}
}

func AssertEqual[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()

	Assert(t, actual == expected, "Expected %v, but got %v.", expected, actual)
}
