package checker

import (
	"fmt"
	"go-monitor/server"
	"net"
	"time"
)

type TCPChecker struct {
	Timeout time.Duration
}

func (t TCPChecker) Check(s server.Server) string {
	addr := net.JoinHostPort(s.IP, fmt.Sprintf("%d", s.Port))
	start := time.Now()
	conn, err := net.DialTimeout(
		"tcp", addr, t.Timeout,
	)
	if err != nil {
		return fmt.Sprintf("[✗] %-20s %s  offline", s.Name, addr)
	}
	conn.Close()
	latency := time.Since(start)
	return fmt.Sprintf("[✓] %-20s %s  online  %.2fms", s.Name, addr, float64(latency.Microseconds()/1000))
}
