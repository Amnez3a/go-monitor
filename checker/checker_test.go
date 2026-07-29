package checker

import (
	"testing"
	"time"

	"go-monitor/server"
)

type MockCheckher struct {
	Result string
}

func (m MockCheckher) Check(s server.Server) string {
	return m.Result
}

func TestFormatResultOnline(t *testing.T) {
	result := formatResult("Alpine", "10.0.0.1:22", true, 15*time.Millisecond)
	want := "[✓] Alpine               10.0.0.1:22  online  15.00ms"
	if result != want {
		t.Errorf("got %q, want %q", result, want)
	}
}
