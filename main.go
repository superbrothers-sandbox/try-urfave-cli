package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.UsageText = fmt.Sprintf("%s [name]", os.Args[0])
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "greeting",
			Usage: "",
			Value: "Hello",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() != 1 {
			return cli.ShowAppHelp(c)
		}

		fmt.Printf("NArg: %d\n", c.NArg())
		fmt.Printf("%s, %s", c.String("greeting"), c.Args()[0])

		return nil
	}

	app.Run(os.Args)
}
