package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Usage: fmt.Sprintf("%s [name]", os.Args[0]),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "greeting",
				Usage: "",
				Value: "Hello",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return cli.ShowAppHelp(c)
			}

			fmt.Printf("NArg: %d\n", c.NArg())
			fmt.Printf("%s, %s\n", c.String("greeting"), c.Args().Get(0))

			return nil
		},
	}

	app.Run(os.Args)
}
