package fox

import "fmt"

// NewFilterEmpty wraps the given logger with a filter for empty values.
func NewFilterEmpty(l Logger) *FilterEmpty {
	return &FilterEmpty{l}
}

type FilterEmpty struct {
	sync Logger
}

// Log calls the underlying logger if v is non empty or len(v) > 1
func (l *FilterEmpty) Log(v ...interface{}) {
	switch len(v) {
	case 0:
		return
	case 1:
		if v[0] == nil {
			return
		}
		if fmt.Sprintf("%v", v[0]) == "" {
			return
		}
	}
	// if there are more values, always log all
	l.sync.Log(v...)
}
