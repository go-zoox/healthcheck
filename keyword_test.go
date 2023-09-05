package healthcheck

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestKeyword(t *testing.T) {
	ok, err := Keyword("https://github.com/page/404", "</html>")
	testify.Equal(t, true, ok)
	testify.Assert(t, err == nil)
}
