package healthcheck

import (
	"fmt"
	"net"
	"time"
)

// UDP checks if the given ip and port is healthy (reachable) via UDP
func UDP(ip string, port int, timeout ...time.Duration) (ok bool, err error) {
	timeoutx := 3 * time.Second
	if len(timeout) > 0 && timeout[0] > 0 {
		timeoutx = timeout[0]
	}

	d := net.Dialer{Timeout: timeoutx}
	conn, err := d.Dial("udp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return false, err
	}
	defer conn.Close()

	return true, nil
}
