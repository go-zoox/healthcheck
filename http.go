package healthcheck

import (
	"time"

	"github.com/go-zoox/fetch"
)

// HTTPOptions is the options for the HTTP health check
type HTTPOptions struct {
	Timeout time.Duration
	//
	Method string
	Header map[string]string
	Query  map[string]string
	Body   any
	//
	OkStatusCodes []int
	//
	Ok func(response *fetch.Response) bool
}

// HTTP checks if the given url is healthy (reachable) via HTTP
func HTTP(url string, opts ...*HTTPOptions) (ok bool, err error) {
	method := "GET"
	timeout := 3 * time.Second
	headers := map[string]string{}
	query := map[string]string{}
	var body any
	if len(opts) > 0 && opts[0] != nil {
		if opts[0].Method != "" {
			method = opts[0].Method
		}

		if opts[0].Header != nil {
			headers = opts[0].Header
		}

		if opts[0].Query != nil {
			query = opts[0].Query
		}

		if opts[0].Body != nil {
			body = opts[0].Body
		}

		if opts[0].Timeout > 0 {
			timeout = time.Duration(opts[0].Timeout) * time.Second
		}
	}

	response, err := fetch.New().
		SetConfig(&fetch.Config{
			Headers: headers,
			Query:   query,
			Body:    body,
			//
			Timeout: timeout,
		}).
		SetMethod(method).
		SetURL(url).
		Execute()
	if err != nil {
		return false, err
	}

	if len(opts) > 0 && opts[0] != nil {
		// Ok function
		if opts[0].Ok != nil {
			return opts[0].Ok(response), nil
		}

		// Ok status codes
		if len(opts[0].OkStatusCodes) > 0 {
			for _, code := range opts[0].OkStatusCodes {
				if response.StatusCode() == code {
					return true, nil
				}
			}
			return false, nil
		}
	}

	// default ok status codes 2xx
	if response.StatusCode() >= 200 && response.StatusCode() < 400 {
		return true, nil
	}

	return false, nil
}
