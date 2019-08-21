package repo

import (
	"io"

	"github.com/spf13/cobra"
)

var repoHelm = `
This command consists of multiple subcommands to interact with chart repositories.
It can be used to init, add, remove, and list chart repositories
`

func NewRepoCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo [FLAGS] init|add|remove|list [ARGS]",
		Short: "Init, add, remove, list and update chart repositories",
		Long:  repoHelm,
	}

	cmd.AddCommand(newRepoInitCmd(out))
	cmd.AddCommand(newRepoAddCmd(out))
	cmd.AddCommand(newRepoListCmd(out))
	cmd.AddCommand(newRepoRemoveCmd(out))

	return cmd
}
