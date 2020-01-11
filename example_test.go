package fox

import "os"

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
