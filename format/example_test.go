package format_test

import (
	"os"

	"github.com/gregoryv/fox"
	. "github.com/gregoryv/fox/format"
)

func ExampleDebug() {
	log := fox.NewSyncLog(os.Stdout).Log
	log(Debug("hello"))
	// output:
	// format/example_test.go:12: hello
}

func ExampleDebugf() {
	log := fox.NewSyncLog(os.Stdout).Log
	log(Debugf("Hello, %s", "world!"))
	// output:
	// format/example_test.go:19: Hello, world!
}
