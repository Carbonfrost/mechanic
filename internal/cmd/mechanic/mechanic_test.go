package mechanic_test

import (
	"bytes"
	"context"

	cli "github.com/Carbonfrost/joe-cli"
	"github.com/Carbonfrost/mechanic/internal/cmd/mechanic"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

var _ = Describe("NewApp", func() {

	Describe("integration", func() {

		DescribeTable("examples", func(args string, expected types.GomegaMatcher) {
			var output bytes.Buffer
			app := mechanic.NewApp()
			app.Stdout = &output
			arguments, _ := cli.Split(args)

			err := app.RunContext(context.Background(), arguments)
			Expect(err).NotTo(HaveOccurred())
			Expect(output.String()).To(expected)
		},
			Entry("-headers",
				"app testdata/sample_1.md -header",
				Equal("<h1>Heading</h1>\n"+
					"<h2>Sub-heading</h2>\n"+
					"<h1>Alternative heading</h1>\n"+
					"<h2>Alternative sub-heading</h2>\n"),
			),
		)
	})

})
