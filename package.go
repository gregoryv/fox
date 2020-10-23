/*
Package fox provides loggers implementing the simple fox.Logger interface.
The fox.Logger interface matches that of testing.T.Log method which
makes it very easy to inject during testing.


Example usage

Use the Log method as a first class citizen

  Log := fox.NewSyncLog(os.Stdout).Log
  Log("some", "nice", "message")

and warnings are simplified by filtering out empty values

  warn := fox.NewSyncLog(os.Stdout).FilterEmpty().Log
  warn("") // will not be logged

  // Log errors only if there are any
  warn(nil) // nothing, it's nil
  warn(io.EOF)

Wrap the standard log package and it's default logger

  Log := fox.LoggerFunc(log.Println)
  Log("hello", "standard", "logger")
*/
package fox

import "fmt"

var NoLogger Logger = LoggerFunc(func(v ...interface{}) {})

type Logger interface {
	Log(...interface{})
}

type LoggerFunc func(...interface{})

func (me LoggerFunc) Log(args ...interface{}) {
	me(args...)
}

// Logging implements github.com/gregoryv/ant Setting interface
type Logging struct {
	Logger
}

func (me Logging) Set(v interface{}) error {
	switch v := v.(type) {
	case usesLogger:
		v.SetLogger(me.Logger)
	default:
		return fmt.Errorf("failed to set %T on %T", me, v)
	}
	return nil
}

type usesLogger interface {
	SetLogger(Logger)
}
