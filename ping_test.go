package healthcheck

import (
	"fmt"
	"testing"

	"github.com/go-zoox/testify"
)

func TestPing(t *testing.T) {
	ok, err := Ping("github.com")
	testify.Assert(t, err == nil)
	testify.Equal(t, true, ok)

	fmt.Println("err:", err)
}
