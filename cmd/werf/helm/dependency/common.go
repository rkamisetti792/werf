package dependency

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/flant/werf/cmd/werf/common"
	"github.com/flant/werf/pkg/util"
)

func isNoRepositoryDefinitionError(err error) bool {
	return strings.HasPrefix(err.Error(), "no repository definition for")
}

func processNoRepositoryDefinitionError(err error) error {
	return fmt.Errorf(strings.Replace(err.Error(), "helm repo add", "werf helm repo add", -1))
}

func getWerfChartPath(commonCmdData common.CmdData) (string, error) {
	var dirOrPWD string

	dirOrPWD, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if *commonCmdData.Dir != "" {
		if path.IsAbs(*commonCmdData.Dir) {
			dirOrPWD = *commonCmdData.Dir
		} else {
			dirOrPWD = path.Clean(path.Join(dirOrPWD, *commonCmdData.Dir))
		}
	}

	exist, err := util.DirExists(path.Join(dirOrPWD, ".helm"))
	if err != nil {
		return "", err
	}

	if exist {
		return path.Join(dirOrPWD, ".helm"), nil
	} else {
		return dirOrPWD, nil
	}
}
