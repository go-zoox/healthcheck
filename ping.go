package healthcheck

import probing "github.com/prometheus-community/pro-bing"

// Ping checks if a domain is reachable.
func Ping(domain string) (ok bool, err error) {
	pinger, err := probing.NewPinger(domain)
	if err != nil {
		return false, err
	}

	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return false, err
	}

	return true, nil
}
