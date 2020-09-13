package format

import (
	"testing"
	"time"

	"github.com/gregoryv/asserter"
)

func TestConfig(t *testing.T) {
	cnf := DefaultConfig
	assert := asserter.New(t)
	assert().Equals(cnf.Info("hello"), "hello")
	assert().Equals(cnf.Infof("%s", "hello"), "hello")
	assert().Equals(cnf.Debug("x"), "format/config_test.go:15: x")
	assert().Equals(cnf.Debugf("msg: %s", "x"), "format/config_test.go:16: msg: x")

	cnf = &Config{now: time.Now, timeFormat: time.RFC3339}
	cnf.SetPrefix("t:")
	assert().Equals(cnf.Info("hello"), "t:hello")
	assert().Equals(cnf.Infof("%s", "hello"), "t:hello")
	assert().Equals(cnf.Debug("x"), "t:format/config_test.go:22: x")
	assert().Equals(cnf.Debugf("msg: %s", "x"), "t:format/config_test.go:23: msg: x")

	var fixed time.Time
	cnf = &Config{
		now:        func() time.Time { return fixed },
		timeFormat: "2006-01-02 15:04:05.999999999 -0700 MST",
	}
	cnf.SetPrefix("PREFIX:")
	cnf.UseTimestamp()
	assert().Equals(cnf.Info("hello"), "0001-01-01 00:00:00 +0000 UTC PREFIX:hello")
	assert().Equals(cnf.Infof("%s", "hello"), "0001-01-01 00:00:00 +0000 UTC PREFIX:hello")
	assert().Equals(cnf.Debug("x"),
		"0001-01-01 00:00:00 +0000 UTC PREFIX:format/config_test.go:34: x")
	assert().Equals(cnf.Debugf("msg: %s", "x"),
		"0001-01-01 00:00:00 +0000 UTC PREFIX:format/config_test.go:36: msg: x")
}
