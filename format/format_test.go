package format

import "testing"

func Test_level_funcs(t *testing.T) {
	ok := func(got, exp string) {
		t.Helper()
		if got != exp {
			t.Error("\n", got, " not equal to\n", exp)
		}
	}
	ok(Debug("x"), "format/format_test.go:12: x")
	ok(Debugf("debug x"), "format/format_test.go:13: debug x")

	ok(Info("x"), "x")
	ok(Infof("info %s", "x"), "info x")

}
