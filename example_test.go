package fox

import (
	"log"
	"os"
)

func ExampleLoggerFunc_Log() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	std := LoggerFunc(log.Println)
	std.Log("hello", "fox")
	// output:
	// hello fox
}

func ExampleSyncLog_Log() {
	Log := NewSyncLog(os.Stdout).Log
	Log("hey")
	// output: hey
}

func ExampleFilterEmpty_Log() {
	Log := NewSyncLog(os.Stdout).FilterEmpty().Log
	Log("hey")
	Log(nil)
	Log("")
	Log("there")
	// output:
	// hey
	// there
}
