package tcp

import (
	"fmt"
	"time"

	"github.com/go-zoox/cli"
	"github.com/go-zoox/core-utils/cast"
	"github.com/go-zoox/healthcheck"
)

func Register(app *cli.MultipleProgram) {
	app.Register("tcp", &cli.Command{
		Name:  "tcp",
		Usage: "tcp is a command that checks tcp port health.",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:  "timeout",
				Usage: "timeout is a flag to set timeout, unit is miliseconds, default is 10 seconds.",
				Value: 10 * 1000,
			},
		},
		Action: func(ctx *cli.Context) error {
			ip := ctx.Args().Get(0)
			port := cast.ToInt(ctx.Args().Get(1))
			if ip == "" || port == 0 {
				return fmt.Errorf("ip and port are required")
			}

			if ok, err := healthcheck.TCP(ip, port, time.Duration(ctx.Int64("timeout"))*time.Millisecond); err != nil {
				return err
			} else if !ok {
				return fmt.Errorf("not ok")
			}

			fmt.Println("ok")
			return nil
		},
	})
}
