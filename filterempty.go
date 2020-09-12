package fox

import "fmt"

// NewFilterEmpty wraps the given logger with a filter for empty values.
func NewFilterEmpty(l Logger) *FilterEmpty {
	return &FilterEmpty{l}
}

type FilterEmpty struct {
	sync Logger
}

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
	l.sync.Log(out)
}
