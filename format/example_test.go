package format_test

import (
	"fmt"

	. "github.com/gregoryv/fox/format"
)

func ExampleDebug() {
	fmt.Println(Debug("hello"))
	// output:
	// format/example_test.go:10: hello
}

func ExampleDebugf() {
	fmt.Println(Debugf("Hello, %s", "world!"))
	// output:
	// format/example_test.go:16: Hello, world!
}
