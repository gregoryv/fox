package fox

import (
	"fmt"
	"io"
	"sync"
)

func NewSyncLog(w io.Writer) *SyncLog {
	return &SyncLog{w: w}
}

type SyncLog struct {
	mu sync.Mutex
	w  io.Writer
}

// Log synchronizes calls to the underlying writer and makes sure
// each message ends with one new line
func (l *SyncLog) Log(v ...interface{}) {
	l.mu.Lock()
	fmt.Fprintln(l.w, v...)
	l.mu.Unlock()
}

var newline byte = '\n'

// FilterEmpty returns a wrapper filtering out empty and nil values
func (l *SyncLog) FilterEmpty() *FilterEmpty {
	return NewFilterEmpty(l)
}

func (l *SyncLog) SetOutput(w io.Writer) { l.w = w }
