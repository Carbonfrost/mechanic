package main

import (
	"os"

	"github.com/Carbonfrost/joe-cli"
	"github.com/Carbonfrost/joe-cli/extensions/color"
	"github.com/Carbonfrost/mechanic/internal/build"
)

func main() {
	createApp().Run(os.Args)
}

func createApp() *cli.App {
	return &cli.App{
		Name:     "mechanic",
		HelpText: "",
		Comment:  "Markdown processing and rendering",
		Uses: cli.Pipeline(
			&color.Options{},
		),
		Action: func(c *cli.Context) error {
			c.Stdout.WriteString("Hello, world!")
			return nil
		},
		Version: build.Version,
		Args:    []*cli.Arg{},
	}
}
