package dependency

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"

	"k8s.io/helm/pkg/chartutil"

	"github.com/flant/werf/cmd/werf/common"
	helm_common "github.com/flant/werf/cmd/werf/helm/common"
)

const dependencyListDesc = `
List all of the dependencies declared in a chart.

This can take chart archives and chart directories as input. It will not alter
the contents of a chart.

This will produce an error if the chart cannot be loaded. It will emit a warning
if it cannot find a requirements.yaml.
`

type dependencyListCmd struct {
	out       io.Writer
	chartpath string
}

func newDependencyListCmd(out io.Writer) *cobra.Command {
	var commonCmdData common.CmdData
	var helmCommonCmdData helm_common.CmdData
	dlc := &dependencyListCmd{out: out}

	cmd := &cobra.Command{
		Use:     "list [flags]",
		Aliases: []string{"ls"},
		Short:   "List the dependencies",
		Long:    dependencyListDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			helm_common.InitHelmSettings(&helmCommonCmdData)

			chartPath, err := getWerfChartPath(commonCmdData)
			if err != nil {
				return err
			}

			dlc.chartpath = chartPath
			return dlc.run()
		},
	}

	common.SetupDir(&commonCmdData, cmd)
	helm_common.SetupHelmSettingsFlags(&helmCommonCmdData, cmd)

	return cmd
}

func (l *dependencyListCmd) run() error {
	c, err := chartutil.Load(l.chartpath)
	if err != nil {
		return err
	}

	r, err := chartutil.LoadRequirements(c)
	if err != nil {
		if err == chartutil.ErrRequirementsNotFound {
			fmt.Fprintf(l.out, "WARNING: no requirements at %s\n", filepath.Join(l.chartpath, "charts"))
			return nil
		}
		return err
	}

	l.printRequirements(r, l.out)
	fmt.Fprintln(l.out)
	l.printMissing(r)
	return nil
}

func (l *dependencyListCmd) dependencyStatus(dep *chartutil.Dependency) string {
	filename := fmt.Sprintf("%s-%s.tgz", dep.Name, "*")
	archives, err := filepath.Glob(filepath.Join(l.chartpath, "charts", filename))
	if err != nil {
		return "bad pattern"
	} else if len(archives) > 1 {
		return "too many matches"
	} else if len(archives) == 1 {
		archive := archives[0]
		if _, err := os.Stat(archive); err == nil {
			c, err := chartutil.Load(archive)
			if err != nil {
				return "corrupt"
			}
			if c.Metadata.Name != dep.Name {
				return "misnamed"
			}

			if c.Metadata.Version != dep.Version {
				constraint, err := semver.NewConstraint(dep.Version)
				if err != nil {
					return "invalid version"
				}

				v, err := semver.NewVersion(c.Metadata.Version)
				if err != nil {
					return "invalid version"
				}

				if constraint.Check(v) {
					return "ok"
				}
				return "wrong version"
			}
			return "ok"
		}
	}

	folder := filepath.Join(l.chartpath, "charts", dep.Name)
	if fi, err := os.Stat(folder); err != nil {
		return "missing"
	} else if !fi.IsDir() {
		return "mispackaged"
	}

	c, err := chartutil.Load(folder)
	if err != nil {
		return "corrupt"
	}

	if c.Metadata.Name != dep.Name {
		return "misnamed"
	}

	if c.Metadata.Version != dep.Version {
		constraint, err := semver.NewConstraint(dep.Version)
		if err != nil {
			return "invalid version"
		}

		v, err := semver.NewVersion(c.Metadata.Version)
		if err != nil {
			return "invalid version"
		}

		if constraint.Check(v) {
			return "unpacked"
		}
		return "wrong version"
	}

	return "unpacked"
}

// printRequirements prints all of the requirements in the yaml file.
func (l *dependencyListCmd) printRequirements(reqs *chartutil.Requirements, out io.Writer) {
	table := uitable.New()
	table.MaxColWidth = 80
	table.AddRow("NAME", "VERSION", "REPOSITORY", "STATUS")
	for _, row := range reqs.Dependencies {
		table.AddRow(row.Name, row.Version, row.Repository, l.dependencyStatus(row))
	}
	fmt.Fprint(out, table)
}

// printMissing prints warnings about charts that are present on disk, but are not in the requirements.
func (l *dependencyListCmd) printMissing(reqs *chartutil.Requirements) {
	folder := filepath.Join(l.chartpath, "charts/*")
	files, err := filepath.Glob(folder)
	if err != nil {
		fmt.Fprintln(l.out, err)
		return
	}

	for _, f := range files {
		fi, err := os.Stat(f)
		if err != nil {
			fmt.Fprintf(l.out, "Warning: %s\n", err)
		}
		// Skip anything that is not a directory and not a tgz file.
		if !fi.IsDir() && filepath.Ext(f) != ".tgz" {
			continue
		}
		c, err := chartutil.Load(f)
		if err != nil {
			fmt.Fprintf(l.out, "WARNING: %q is not a chart.\n", f)
			continue
		}
		found := false
		for _, d := range reqs.Dependencies {
			if d.Name == c.Metadata.Name {
				found = true
				break
			}
		}
		if !found {
			fmt.Fprintf(l.out, "WARNING: %q is not in requirements.yaml.\n", f)
		}
	}

}
