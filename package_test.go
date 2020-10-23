package fox

import "testing"

func TestLogging(t *testing.T) {
	log := Logging{t}
	if err := log.Set(1); err == nil {
		t.Error("should fail")
	}

	if err := log.Set(&x{}); err != nil {
		t.Error(err)
	}
}

type x struct{}

// Method
func (*x) SetLogger(Logger) {}
