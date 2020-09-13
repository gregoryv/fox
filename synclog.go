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
