package mechanic

import (
	"io"
	"os"

	cli "github.com/Carbonfrost/joe-cli"
	"github.com/Carbonfrost/joe-cli/extensions/color"
	"github.com/Carbonfrost/mechanic/internal/build"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func Run(args []string) {
	NewApp().Run(args)
}

func NewApp() *cli.App {
	return &cli.App{
		Name:     "mechanic",
		HelpText: "",
		Comment:  "Markdown processing and rendering",
		Uses: cli.Pipeline(
			&color.Options{},
		),
		Action:  processExpression,
		Version: build.Version,
		Args: []*cli.Arg{
			{
				Name:    "files",
				Value:   new(cli.FileSet),
				Options: cli.Merge,
				NArg:    cli.TakeUntilNextFlag,
			},
			{
				Name: "expression",
				Value: &cli.Expression{
					Exprs: Exprs(),
				},
			},
		},
	}
}

func processExpression(c *cli.Context) error {
	return c.FileSet("files").Do(func(f *cli.File, err error) error {
		if err != nil {
			return err
		}

		source, node, err := parseDocument(f.Name)
		if err != nil {
			return err
		}

		exp := ensurePrinter(c.Expression("expression"), source)
		return exp.Evaluate(c, newSet(node))
	})
}

func parseDocument(file string) ([]byte, ast.Node, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}

	source, err := io.ReadAll(f)
	if err != nil {
		return nil, nil, err
	}

	node := goldmark.DefaultParser().Parse(text.NewReader(source))
	return source, node, nil
}
