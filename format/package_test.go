package format

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_level_funcs(t *testing.T) {
	var (
		assert = asserter.New(t)
		ok, _  = assert().Errors()
	)
	ok(equals(Info("x"), "x"))
	ok(equals(Infof("info %s", "x"), "info x"))
	ok(equals(Debug("x"), "format/package_test.go:18: x"))
	ok(equals(Debugf("debug x"), "format/package_test.go:19: debug x"))

	Caller = BrokenCaller
	ok(equals(Debug("x"), "x"))
	Caller = runtime.Caller // Reset for subsequent tests
}

func equals(a, b interface{}) error {
	if a != b {
		return fmt.Errorf("%v != %v", a, b)
	}
	return nil
}

func BrokenCaller(int) (uintptr, string, int, bool) {
	var x uintptr
	return x, "", 0, false
}
