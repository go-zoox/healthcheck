package healthcheck

import (
	"github.com/go-zoox/core-utils/strings"

	"github.com/go-zoox/fetch"
)

type KeywordOptions struct {
	HTTPOptions
}

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
