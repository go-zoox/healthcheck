package ping

import (
	"fmt"
	"time"

	"github.com/go-zoox/cli"
	"github.com/go-zoox/healthcheck"
)

func Register(app *cli.MultipleProgram) {
	app.Register("ping", &cli.Command{
		Name:  "ping",
		Usage: "ping is a command that checks icmp health.",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:  "timeout",
				Usage: "timeout is a flag to set timeout, unit is miliseconds, default is 10 seconds.",
				Value: 10 * 1000,
			},
		},
		Action: func(ctx *cli.Context) error {
			domain := ctx.Args().First()
			if domain == "" {
				return fmt.Errorf("domain is required")
			}

			if ok, err := healthcheck.Ping(domain, time.Duration(ctx.Int64("timeout"))*time.Millisecond); err != nil {
				return err
			} else if !ok {
				return fmt.Errorf("not ok")
			}

			fmt.Println("ok")
			return nil
		},
	})
}
