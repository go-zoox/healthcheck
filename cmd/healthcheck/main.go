package main

import (
	"github.com/go-zoox/healthcheck/cmd/healthcheck/commands/http"
	"github.com/go-zoox/healthcheck/cmd/healthcheck/commands/ping"
	"github.com/go-zoox/healthcheck/cmd/healthcheck/commands/tcp"
	"github.com/go-zoox/healthcheck/cmd/healthcheck/commands/udp"

	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewMultipleProgram(&cli.MultipleProgramConfig{
		Name:  "healthcheck",
		Usage: "healthcheck is a command that checks health.",
	})

	http.Register(app)
	tcp.Register(app)
	udp.Register(app)
	ping.Register(app)

	app.Run()
}
