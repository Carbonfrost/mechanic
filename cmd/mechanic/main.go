package main

import (
	"io"
	"os"

	"github.com/Carbonfrost/joe-cli"
	"github.com/Carbonfrost/joe-cli/extensions/color"
	"github.com/Carbonfrost/mechanic/internal/build"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
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
			return c.FileSet("files").Do(func(f *cli.File, err error) error {
				if err != nil {
					return err
				}
				return renderPage(f.Name)
			})
		},
		Version: build.Version,
		Args: []*cli.Arg{
			{
				Name:    "files",
				Value:   new(cli.FileSet),
				Options: cli.Merge,
			},
		},
	}
}

func renderPage(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	source, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	node := goldmark.DefaultParser().Parse(text.NewReader(source))
	renderer := goldmark.DefaultRenderer()
	renderer.Render(os.Stdout, source, node)
	return nil
}
