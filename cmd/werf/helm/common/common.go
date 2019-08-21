package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	helm_env "k8s.io/helm/pkg/helm/environment"
	"k8s.io/helm/pkg/helm/helmpath"

	"github.com/flant/werf/cmd/werf/common"
	"github.com/flant/werf/pkg/tag_strategy"
)

func GetImagesRepoOrStub(imagesRepoOption string) string {
	if imagesRepoOption == "" {
		return "IMAGES_REPO"
	}
	return imagesRepoOption
}

func GetEnvironmentOrStub(environmentOption string) string {
	if environmentOption == "" {
		return "ENV"
	}
	return environmentOption
}

func GetTagOrStub(commonCmdData *common.CmdData) (string, tag_strategy.TagStrategy, error) {
	tag, tagStrategy, err := common.GetDeployTag(commonCmdData, common.TagOptionsGetterOptions{Optional: true})
	if err != nil {
		return "", "", err
	}

	if tag == "" {
		tag, tagStrategy = "TAG", tag_strategy.Custom
	}

	return tag, tagStrategy, nil
}

var (
	HelmSettings *helm_env.EnvSettings
)

type CmdData struct {
	helmSettingsHome  *string
	helmSettingsDebug *bool
}

func SetupHelmSettingsFlags(cmdData *CmdData, cmd *cobra.Command) {
	cmdData.helmSettingsHome = new(string)
	cmdData.helmSettingsDebug = new(bool)

	helmSettingsHomeDefaultValue := os.Getenv("HELM_HOME")
	if helmSettingsHomeDefaultValue == "" {
		helmSettingsHomeDefaultValue = helm_env.DefaultHelmHome
	}

	cmd.Flags().StringVarP(cmdData.helmSettingsHome, "helm-home", "", helmSettingsHomeDefaultValue, "location of your Helm config. Defaults to $HELM_HOME")
	cmd.Flags().BoolVarP(cmdData.helmSettingsDebug, "debug", "", common.GetBoolEnvironment("HELM_DEBUG"), "enable verbose output. Defaults to $HELM_DEBUG")
}

func InitHelmSettings(cmdData *CmdData) {
	HelmSettings = new(helm_env.EnvSettings)
	HelmSettings.Home = helmpath.Home(*cmdData.helmSettingsHome)
	HelmSettings.Debug = *cmdData.helmSettingsDebug
}

// DefaultKeyring returns the expanded path to the default keyring.
func DefaultKeyring() string {
	return os.ExpandEnv("$HOME/.gnupg/pubring.gpg")
}

func CheckArgsLength(argsReceived int, requiredArgs ...string) error {
	expectedNum := len(requiredArgs)
	if argsReceived != expectedNum {
		arg := "arguments"
		if expectedNum == 1 {
			arg = "argument"
		}
		return fmt.Errorf("this command needs %v %s: %s", expectedNum, arg, strings.Join(requiredArgs, ", "))
	}
	return nil
}

var CouldNotLoadRepositoriesFileErrorFormat = "could not load repositories file (%s): you might need to run `werf helm repo init`"

func IsCouldNotLoadRepositoriesFileError(err error) bool {
	return strings.HasPrefix(err.Error(), "Couldn't load repositories file")
}
