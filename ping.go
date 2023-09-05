package healthcheck

import (
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

// Ping checks if a domain is reachable.
func Ping(domain string, timeout ...time.Duration) (ok bool, err error) {
	timeoutx := 3 * time.Second
	if len(timeout) > 0 && timeout[0] > 0 {
		timeoutx = timeout[0]
	}

	pinger, err := probing.NewPinger(domain)
	if err != nil {
		return false, err
	}

	pinger.Timeout = timeoutx
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return false, err
	}

	return true, nil
}
