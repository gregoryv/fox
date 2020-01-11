package format

import (
	"fmt"
	"testing"
)

func Test_level_funcs(t *testing.T) {
	ok := func(got, exp string) {
		t.Helper()
		if got != exp {
			t.Error("\n", got, " not equal to\n", exp)
		}
	}
	ok(Debug("x"), "format/unit_test.go:15: x")
	ok(Debugf("debug x"), "format/unit_test.go:16: debug x")

	ok(Info("x"), "x")
	ok(Infof("info %s", "x"), "info x")
}

func ExampleDebug() {
	fmt.Println(Debug("hello"))
	// output:
	// format/unit_test.go:23: hello
}

func ExampleDebugf() {
	fmt.Println(Debugf("Hello, %s", "world!"))
	// output:
	// format/unit_test.go:29: Hello, world!
}
