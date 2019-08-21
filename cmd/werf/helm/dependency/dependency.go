package dependency

import (
	"io"

	"github.com/spf13/cobra"
)

func NewDependencyCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "dependency update|build|list",
		Aliases: []string{"dep", "dependencies"},
		Short:   "Manage a chart's dependencies",
	}

	cmd.AddCommand(newDependencyListCmd(out))
	cmd.AddCommand(newDependencyUpdateCmd(out))
	cmd.AddCommand(newDependencyBuildCmd(out))

	return cmd
}
