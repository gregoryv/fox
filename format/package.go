/*
   Package format generates strings for logging

      import (
         "github.com/gregoryv/fox"
         . "github.com/gregoryv/fox/format"
      )

      Log := fox.NewSyncLog(os.Stdout).Log
      Log(Debugf("failed: %v", err))

   Debug and Debugf include the name of file and lineno of where they where called
   similar but not the same as using the builtin log package with Lshortfile

      log := New(os.Stdout, "", Lshortfile)
      log.Printf("failed: %v", err)

*/
package format

import (
	"fmt"
	"path"
	"runtime"
)

// Info returns same as fmt.Sprint, here for consistency
func Info(v ...interface{}) string { return fmt.Sprint(v...) }

// Info returns same as fmt.Sprintf
func Infof(f string, v ...interface{}) string { return fmt.Sprintf(f, v...) }

// Debug prefixes the message with file information.
// In the form parent/file:lineno: message
func Debug(v ...interface{}) string {
	return debug(2, fmt.Sprint(v...))
}

// Debugf formats and prefixes message same as Debug
func Debugf(f string, v ...interface{}) string {
	return debug(2, fmt.Sprintf(f, v...))
}

var Caller func(int) (uintptr, string, int, bool) = runtime.Caller

func debug(calldepth int, v string) string {
	_, file, line, ok := Caller(calldepth)
	if !ok {
		return v
	}
	dir := path.Base(path.Dir(file))
	filename := path.Base(file)
	return fmt.Sprintf("%s/%s:%v: %s", dir, filename, line, v)
}
