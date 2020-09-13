package fox_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregoryv/fox"
	"github.com/gregoryv/fox/format"
)

func BenchmarkConfig_Info(b *testing.B) {
	cnf := format.NewConfig()
	cnf.Stamp()
	cnf.SetPrefix("BENCH:")
	err := fmt.Errorf("bench err message")
	var buf bytes.Buffer
	log := fox.NewSyncLog(&buf).FilterEmpty().Log

	b.Run("Info", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			log(cnf.Info("info", err))
		}
	})

	b.Run("Infof", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			log(cnf.Infof("%s %v", "infof", err))
		}
	})

	b.Run("Debug", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			log(cnf.Debug("debug", err))
		}
	})

	b.Run("Debugf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			log(cnf.Debugf("debugf %v", err))
		}
	})

}
