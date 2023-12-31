package testing

import (
	"testing"

	"github.com/suzuki-shunsuke/ghcp/pkg/logger"
)

func New(t *testing.T) logger.Interface {
	return &testLogger{t}
}

type testLogger struct {
	t *testing.T
}

func (l *testLogger) Errorf(format string, v ...interface{}) {
	l.t.Logf("ERROR: "+format, v...)
}

func (l *testLogger) Warnf(format string, v ...interface{}) {
	l.t.Logf("WARN: "+format, v...)
}

func (l *testLogger) Infof(format string, v ...interface{}) {
	l.t.Logf("INFO: "+format, v...)
}

func (l *testLogger) Debugf(format string, v ...interface{}) {
	l.t.Logf("DEBUG: "+format, v...)
}
