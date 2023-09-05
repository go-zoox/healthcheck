package healthcheck

import (
	"testing"

	"github.com/go-zoox/fetch"
	"github.com/go-zoox/testify"
)

func TestHTTP(t *testing.T) {
	ok, err := HTTP("https://github.com")
	testify.Equal(t, true, ok)
	testify.Assert(t, err == nil)
}

func TestHTTPWithOkStatusCodes(t *testing.T) {
	ok, err := HTTP("https://github.com/api/404", &HTTPOptions{
		OkStatusCodes: []int{502, 404},
	})
	testify.Assert(t, err == nil)
	testify.Equal(t, true, ok)
}

func TestHTTPWithOkFunc(t *testing.T) {
	ok, err := HTTP("https://github.com", &HTTPOptions{
		Ok: func(response *fetch.Response) bool {
			// return response.StatusCode() == 200
			return response.String() != ""
		},
	})
	testify.Assert(t, err == nil)
	testify.Equal(t, true, ok)
}
