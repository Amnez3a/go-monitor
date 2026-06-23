package checker

import (
	"testing"

	"go-monitor/server"
)

type MockCheckher struct {
	Result string
}

func (m MockCheckher) Check(s server.Server) string {
	return m.Result
}

func TestOnlineServer(t *testing.T) {
	mock := MockCheckher{Result: "[✓] Alpine online"}
	result := mock.Check(server.Server{Name: "Alpine"})

	if result != "[✓] Alpine online" {
		t.Errorf("got %s, want [✓] Alpine online", result)
	}
}

func TestOffline(t *testing.T) {
	mock := MockCheckher{Result: "[✗] Alpine offline"}
	result := mock.Check(server.Server{Name: "Alpine"})

	if result != "[✗] Alpine offline" {
		t.Errorf("got %s, want [✗] Alpine offline", result)
	}
}
