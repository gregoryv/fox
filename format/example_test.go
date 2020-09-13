package format_test

import (
	"os"

	"github.com/gregoryv/fox"
	"github.com/gregoryv/fox/format"
	. "github.com/gregoryv/fox/format"
)

func ExampleDebug() {
	log := fox.NewSyncLog(os.Stdout).Log
	log(Debug("hello"))
	// output:
	// format/example_test.go:13: hello
}

func ExampleDebugf() {
	log := fox.NewSyncLog(os.Stdout).Log
	log(Debugf("Hello, %s", "world!"))
	// output:
	// format/example_test.go:20: Hello, world!
}

func ExampleInfo() {
	type Car struct {
		fox.Logger
		*format.Config
	}
	me := Car{
		Logger: fox.NewSyncLog(os.Stdout),
		Config: format.NewConfig(),
	}
	me.Config.UseTimestamp()
	me.Config.SetPrefix("[Car x]")

	me.Log(me.Info("got a time"))
}
