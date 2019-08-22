package repo

import (
	"errors"
	"fmt"
	"io"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"

	"k8s.io/helm/pkg/helm/helmpath"
	"k8s.io/helm/pkg/repo"

	"github.com/flant/werf/cmd/werf/helm/common"
)

type repoListCmd struct {
	out  io.Writer
	home helmpath.Home
}

func newRepoListCmd(out io.Writer) *cobra.Command {
	var commonCmdData common.CmdData
	list := &repoListCmd{out: out}

	cmd := &cobra.Command{
		Use:   "list [flags]",
		Short: "List chart repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			common.InitHelmSettings(&commonCmdData)

			list.home = common.HelmSettings.Home
			return list.run()
		},
	}

	common.SetupHelmSettingsFlags(&commonCmdData, cmd)

	return cmd
}

func (a *repoListCmd) run() error {
	f, err := repo.LoadRepositoriesFile(a.home.RepositoryFile())
	if err != nil {
		if common.IsCouldNotLoadRepositoriesFileError(err) {
			return fmt.Errorf(common.CouldNotLoadRepositoriesFileErrorFormat, a.home.RepositoryFile())
		}

		return err
	}
	if len(f.Repositories) == 0 {
		return errors.New("no repositories to show")
	}
	table := uitable.New()
	table.AddRow("NAME", "URL")
	for _, re := range f.Repositories {
		table.AddRow(re.Name, re.URL)
	}
	fmt.Fprintln(a.out, table)
	return nil
}
