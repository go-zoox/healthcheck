package healthcheck

import (
	"fmt"
	"testing"

	"github.com/go-zoox/testify"
)

func TestTCP(t *testing.T) {
	ok, err := TCP("20.205.243.166", 443)
	testify.Assert(t, err == nil)
	testify.Equal(t, true, ok)

	fmt.Println("err:", err)
}

func TestTCPWithDomain(t *testing.T) {
	ok, err := TCP("github.com", 443)
	testify.Assert(t, err == nil)
	testify.Equal(t, true, ok)

	fmt.Println("err:", err)
}
