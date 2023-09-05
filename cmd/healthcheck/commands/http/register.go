package http

import (
	"fmt"
	"time"

	"github.com/go-zoox/cli"
	"github.com/go-zoox/healthcheck"
)

func Register(app *cli.MultipleProgram) {
	app.Register("http", &cli.Command{
		Name:  "http",
		Usage: "http is a command that checks http health.",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:  "timeout",
				Usage: "timeout is a flag to set timeout, unit is miliseconds, default is 10 seconds.",
				Value: 10 * 1000,
			},
		},
		Action: func(ctx *cli.Context) error {
			url := ctx.Args().First()
			if url == "" {
				return fmt.Errorf("url is required")
			}

			ok, err := healthcheck.HTTP(url, &healthcheck.HTTPOptions{
				Timeout: time.Duration(ctx.Int64("timeout")) * time.Millisecond,
			})
			if err != nil {
				return err
			} else if !ok {
				return fmt.Errorf("not ok")
			}

			fmt.Println("ok")
			return nil
		},
	})
}
