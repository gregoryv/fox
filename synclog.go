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

import (
	"fmt"
	"io"
	"sync"
)

func NewSyncLog(w io.Writer) *SyncLog {
	return &SyncLog{w: w}
}

type Logger interface {
	Log(...interface{})
}

type LoggerFunc func(...interface{})

func (me LoggerFunc) Log(args ...interface{}) {
	me(args...)
}

type SyncLog struct {
	mu sync.Mutex
	w  io.Writer
}

// Log synchronizes calls to the underlying writer and makes sure
// each message ends with one new line
func (l *SyncLog) Log(v ...interface{}) {
	out := fmt.Sprint(v...)
	l.mu.Lock()
	fmt.Fprint(l.w, out)
	if len(out) == 0 || out[len(out)-1] != newline {
		l.w.Write([]byte{newline})
	}
	l.mu.Unlock()
}

var newline byte = '\n'

// FilterEmpty returns a wrapper filtering out empty and nil values
func (l *SyncLog) FilterEmpty() *FilterEmpty {
	return NewFilterEmpty(l)
}

func (l *SyncLog) SetOutput(w io.Writer) { l.w = w }
