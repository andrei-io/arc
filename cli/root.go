package cli

import (
	"github.com/spf13/cobra"
)

var (
	cli = &cobra.Command{
		Use:   "arctic",
		Short: "Smart project generator",
		Long: `Arctic is a project generator.
This application is a tool to generate the needed files
to quickly create an application.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return cli.Execute()
}

func init() {
	cli.AddCommand(versionController)
	cli.AddCommand(createController)
}