package checker

import (
	"fmt"
	"net"
	"time"

	"go-monitor/server"
)

type TCPChecker struct {
	Timeout time.Duration
}

func formatResult(name string, addr string, online bool, latency time.Duration) string {
	if !online {
		return fmt.Sprintf("[✗] %-20s %s  offline", name, addr)
	}
	return fmt.Sprintf("[✓] %-20s %s  online  %.2fms", name, addr, float64(latency.Microseconds())/1000)
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
	return formatResult(s.Name, addr, true, latency)
}
