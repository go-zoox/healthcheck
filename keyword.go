package healthcheck

import (
	"github.com/go-zoox/core-utils/strings"

	"github.com/go-zoox/fetch"
)

// KeywordOptions is the options for keyword health check
type KeywordOptions struct {
	HTTPOptions
}

// Keyword checks if the given url is healthy (reachable) via HTTP and contains the given keyword
func Keyword(url, keyword string, opts ...*KeywordOptions) (bool, error) {
	optx := HTTPOptions{}
	if len(opts) > 0 && opts[0] != nil {
		optx = opts[0].HTTPOptions
	}

	return HTTP(url, &HTTPOptions{
		Timeout: optx.Timeout,
		Method:  optx.Method,
		Header:  optx.Header,
		Query:   optx.Query,
		Body:    optx.Body,
		//
		Ok: func(response *fetch.Response) bool {
			return strings.Contains(response.String(), keyword)
		},
	})
}
