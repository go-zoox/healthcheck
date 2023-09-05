package healthcheck

import "time"

// HealthCheck is the interface for health check
type HealthCheck interface {
	// HTTP checks if the given url is healthy (reachable) via HTTP
	HTTP(url string, opts ...*HTTPOptions) (ok bool, err error)

	// TCP checks if the given ip and port is healthy (reachable) via TCP
	TCP(ip string, port int, timeout ...time.Duration) (ok bool, err error)

	// UDP checks if the given ip and port is healthy (reachable) via UDP
	UDP(ip string, port int, timeout ...time.Duration) (ok bool, err error)

	// Ping checks if the given domain is healthy (reachable) via ICMP
	Ping(domain string, timeout ...time.Duration) (ok bool, err error)

	// Keyword checks if the given url is healthy (reachable) via HTTP and contains the given keyword
	Keyword(url, keyword string, opts ...*KeywordOptions) (bool, error)
}

type healthcheck struct{}

// New returns a new HealthCheck
func New() HealthCheck {
	return &healthcheck{}
}

func (h *healthcheck) HTTP(url string, opts ...*HTTPOptions) (ok bool, err error) {
	return HTTP(url, opts...)
}

func (h *healthcheck) TCP(ip string, port int, timeout ...time.Duration) (ok bool, err error) {
	return TCP(ip, port, timeout...)
}

func (h *healthcheck) UDP(ip string, port int, timeout ...time.Duration) (ok bool, err error) {
	return UDP(ip, port, timeout...)
}

func (h *healthcheck) Ping(domain string, timeout ...time.Duration) (ok bool, err error) {
	return Ping(domain, timeout...)
}

func (h *healthcheck) Keyword(url, keyword string, opts ...*KeywordOptions) (bool, error) {
	return Keyword(url, keyword, opts...)
}
