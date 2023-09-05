package healthcheck

import (
	"fmt"
	"net"
)

// TCP checks if the given ip and port is healthy (reachable) via TCP
func TCP(ip string, port int) (ok bool, err error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return false, err
	}
	defer conn.Close()

	return true, nil
}
