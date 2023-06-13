package markdown

import (
	"context"
	"fmt"

	cli "github.com/Carbonfrost/joe-cli"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
)

// ContextServices provides an adapter around the context to
type ContextServices struct {
}

type contextKey string

const (
	servicesKey contextKey = "markdown.services"
)

// WithServices obtains the context services from the specified context.
// If they do not exist, they are added and the context result is returned.
func WithServices(c context.Context) (context.Context, *ContextServices) {
	if o, ok := Services(c); ok {
		return c, o
	}
	res := &ContextServices{}
	return context.WithValue(c, servicesKey, res), res
}

// Services obtains the contextual services used by the package.
func Services(c context.Context) (*ContextServices, bool) {
	o, ok := c.Value(servicesKey).(*ContextServices)
	return o, ok
}

func SetContext() cli.Action {
	return cli.ActionFunc(func(c *cli.Context) error {
		s, _ := WithServices(c.Context())
		return c.SetContext(s)
	})
}

// Must panics if v is an error or equal to false.  Must is typically used
// to assert the result of Services method where it is expected to exist.
func Must(c *ContextServices, v any) *ContextServices {
	if c == nil || v == false {
		v = fmt.Errorf("not found")
	}
	if err, ok := v.(error); ok {
		panic(fmt.Errorf("expected markdown context services: %w", err))
	}
	return c
}

func (c *ContextServices) Markdown() goldmark.Markdown {
	md := goldmark.New()
	return md
}

func (c *ContextServices) Renderer() renderer.Renderer {
	return c.Markdown().Renderer()
}

func (c *ContextServices) Parser() parser.Parser {
	return c.Markdown().Parser()
}
