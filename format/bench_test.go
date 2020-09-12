package format

import (
	"fmt"
	"testing"
)

func BenchmarkConfig_Info(b *testing.B) {
	cnf := NewConfig()
	cnf.Stamp()
	cnf.SetPrefix("BENCH:")
	err := fmt.Errorf("bench err message")

	b.Run("Info", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cnf.Info("info", err)
		}
	})

	b.Run("Infof", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cnf.Infof("%s %v", "infof", err)
		}
	})

	b.Run("Debug", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cnf.Debug("debug", err)
		}
	})

	b.Run("Debugf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cnf.Debugf("debugf %v", err)
		}
	})

}
