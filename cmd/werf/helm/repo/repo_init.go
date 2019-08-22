package repo

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"k8s.io/helm/cmd/helm/installer"
	"k8s.io/helm/pkg/helm/helmpath"

	"github.com/flant/werf/cmd/werf/helm/common"
)

var (
	stableRepositoryURL = "https://kubernetes-charts.storage.googleapis.com"
	// This is the IPv4 loopback, not localhost, because we have to force IPv4
	// for Dockerized Helm: https://github.com/kubernetes/helm/issues/1410
	localRepositoryURL = "http://127.0.0.1:8879/charts"
)

type initCmd struct {
	skipRefresh bool
	out         io.Writer
	home        helmpath.Home
}

func newRepoInitCmd(out io.Writer) *cobra.Command {
	var commonCmdData common.CmdData
	i := &initCmd{out: out}

	cmd := &cobra.Command{
		Use:   "init [flags]",
		Short: "Init chart repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			common.InitHelmSettings(&commonCmdData)

			i.home = common.HelmSettings.Home

			if err := installer.Initialize(i.home, out, i.skipRefresh, *common.HelmSettings, stableRepositoryURL, localRepositoryURL); err != nil {
				return fmt.Errorf("error initializing: %s", err)
			}
			fmt.Fprintf(out, "%s has been configured\n", i.home)

			return nil
		},
	}

	f := cmd.Flags()
	f.BoolVar(&i.skipRefresh, "skip-refresh", false, "do not refresh (download) the local repository cache")
	f.StringVar(&stableRepositoryURL, "stable-repo-url", stableRepositoryURL, "URL for stable repository")
	f.StringVar(&localRepositoryURL, "local-repo-url", localRepositoryURL, "URL for local repository")

	common.SetupHelmSettingsFlags(&commonCmdData, cmd)

	return cmd
}
