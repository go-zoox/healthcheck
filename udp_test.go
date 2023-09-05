package healthcheck

import (
	"fmt"
	"testing"

	"github.com/go-zoox/testify"
)

func TestUDP(t *testing.T) {
	ok, err := UDP("8.8.8.8", 53)
	testify.Assert(t, err == nil)
	testify.Equal(t, true, ok)

	fmt.Println("err:", err)
}
