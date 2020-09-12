package format

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

var DefaultConfig = NewConfig()

// NewConfig returns a new message format configuration.
func NewConfig() *Config {
	return &Config{
		now:        time.Now,
		timeFormat: time.RFC3339,
	}
}

type Config struct {
	prefix     string
	now        func() time.Time
	stamp      bool
	timeFormat string
}

// SetHeader
func (me *Config) SetPrefix(v string) { me.prefix = v }

// Stamp turns timestamps on
func (me *Config) Stamp() { me.stamp = true }

// Info returns same as fmt.Sprint, here for consistency
func (me *Config) Info(v ...interface{}) string {
	var sb strings.Builder
	sb.WriteString(me.header(me.now()))
	sb.WriteString(fmt.Sprint(v...))
	return sb.String()
}

// Info returns same as fmt.Sprintf
func (me *Config) Infof(f string, v ...interface{}) string {
	var sb strings.Builder
	sb.WriteString(me.header(me.now()))
	sb.WriteString(fmt.Sprintf(f, v...))
	return sb.String()
}

// Debug prefixes the message with file information.
// In the form parent/file:lineno: message.
// Returns only message if file information is not available.
func (me *Config) Debug(v ...interface{}) string {
	return me.debug(2, fmt.Sprint(v...))
}

// Debugf formats and prefixes message same as Debug
func (me *Config) Debugf(f string, v ...interface{}) string {
	return me.debug(2, fmt.Sprintf(f, v...))
}

// Here so we can test behaviour of debug formats
var caller func(int) (uintptr, string, int, bool) = runtime.Caller

func (me *Config) debug(calldepth int, v string) string {
	now := me.now() // early
	_, file, line, ok := caller(calldepth)
	if !ok {
		return v
	}
	dir := path.Base(path.Dir(file))
	filename := path.Base(file)
	return fmt.Sprintf("%s%s/%s:%v: %s", me.header(now), dir, filename, line, v)
}

// header returns a header based on the configuration fields
func (me *Config) header(now time.Time) string {
	var sb strings.Builder
	if me.stamp {
		t := now.Format(me.timeFormat)
		sb.WriteString(t)
	}
	if me.prefix != "" {
		if me.stamp {
			sb.WriteString(" ")
		}
		sb.WriteString(me.prefix)
	}
	return sb.String()
}
