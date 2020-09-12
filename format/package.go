/*
Package format generates strings for logging.

      import (
         "github.com/gregoryv/fox"
         . "github.com/gregoryv/fox/format"
      )

      Log := fox.NewSyncLog(os.Stdout).Log
      Log(Debugf("failed: %v", err))

Debug and Debugf include the name of file and line number of the call
location similar but not the same as using the builtin log package
with Lshortfile

      log := New(os.Stdout, "", Lshortfile)
      log.Printf("failed: %v", err)

*/
package format

import (
	"fmt"
)

// Info wraps DefaultConfig.Info
func Info(v ...interface{}) string { return DefaultConfig.Info(v...) }

// Infof wraps DefaultConfig.Infof
func Infof(f string, v ...interface{}) string {
	return DefaultConfig.Infof(f, v...)
}

// Debug wraps DefaultConfig.Debug
func Debug(v ...interface{}) string {
	return DefaultConfig.debug(2, fmt.Sprint(v...))
}

// Debugf wraps DefaultConfig.Debugf
func Debugf(f string, v ...interface{}) string {
	return DefaultConfig.debug(2, fmt.Sprintf(f, v...))
}
