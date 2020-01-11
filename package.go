/*
Package log provides loggers for writing messages.

    std := NewSyncLog(os.Stdout)
    std.Log("some", "nice", "message")

    filter := std.FilterEmpty()
    filter.Log("") // will not be logged

    // Log errors only if there are any
    var err error
    filter.Log(err) // nothing, it's nil
    err = io.EOF
    filter.Log(err)
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
	return &FilterEmpty{l}
}

func (l *SyncLog) SetOutput(w io.Writer) { l.w = w }

type FilterEmpty struct{ *SyncLog }

// Log calls the underlying logger only if v is non empty
func (l *FilterEmpty) Log(v ...interface{}) {
	switch len(v) {
	case 0:
		return
	case 1:
		if v[0] == nil {
			return
		}
	}
	out := fmt.Sprint(v...)
	if out == "" {
		return
	}
	l.SyncLog.Log(out)
}
