package mechanic

import (
	"fmt"

	cli "github.com/Carbonfrost/joe-cli"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
)

const (
	filterCategory = "filter on attributes"
)

type nodeSet struct {
	nodes []ast.Node
}

func newSet(n ast.Node) *nodeSet {
	return &nodeSet{
		nodes: []ast.Node{n},
	}
}

func Exprs() []*cli.Expr {
	return []*cli.Expr{
		{
			Name:     "header", // -header
			Evaluate: nodesOfType(ast.KindHeading),
			Category: filterCategory,
		},
	}
}

func Predicate(filter func(*cli.Context, ast.Node) bool) cli.EvaluatorFunc {
	return func(c *cli.Context, v any, yield func(any) error) error {
		switch val := v.(type) {
		case ast.Node:
			if filter(c, val) {
				return yield(val)
			}
			return nil
		case *nodeSet:
			err := val.Walk(func(a ast.Node, entering bool) (ast.WalkStatus, error) {
				if entering && filter(c, a) {
					if err := yield(a); err != nil {
						return ast.WalkStop, err
					}

				}
				return ast.WalkContinue, nil
			})
			return err
		}
		panic(fmt.Sprintf("unreachable! %T", v))
	}
}

func Render(source []byte) cli.Evaluator {
	return Predicate(func(ctx *cli.Context, node ast.Node) bool {
		renderer := goldmark.DefaultRenderer()
		err := renderer.Render(ctx.Stdout, source, node)

		return alwaysTrue(node, err)
	})
}

func alwaysTrue(a ast.Node, err error) bool {
	return true
}

func ensurePrinter(e *cli.Expression, source []byte) *cli.Expression {
	e.Append(cli.NewExprBinding(Render(source)))
	return e
}

func (s *nodeSet) Walk(fn func(ast.Node, bool) (ast.WalkStatus, error)) error {
	for _, node := range s.nodes {
		err := ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
			return fn(n, entering)
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func nodesOfType(kind ast.NodeKind) cli.Evaluator {
	return Predicate(func(c *cli.Context, n ast.Node) bool {
		return n.Kind() == kind
	})
}
