package log

import (
	"bytes"
	"io"
	"testing"
)

func Test_logging(t *testing.T) {
	var buf bytes.Buffer
	ok := func(exp string) {
		t.Helper()
		got := buf.String()
		if got != exp {
			t.Errorf("got %q, expected %q", got, exp)
		}
		buf.Reset()
	}

	l := NewSyncLog(nil)
	l.SetOutput(&buf)
	l.Log("x")
	ok("x\n")

	l.Log()
	ok("\n")

	f := l.FilterEmpty()
	f.Log()
	ok("")
	f.Log("")
	ok("")
	f.Log("x")
	ok("x\n")
	f.Log(nil)
	ok("")
	f.Log(io.EOF)
	ok("EOF\n")
}
