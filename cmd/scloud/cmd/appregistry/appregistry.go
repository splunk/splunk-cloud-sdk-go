package appregistry

////go:generate scloudgen gen-cmd --name app-registry --package appreg --output appreg-gen.go

import (
	"github.com/spf13/cobra"
	usageUtil "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/util"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return appregistryCmd
}

// appregistryCmd represents the App Registry command
var appregistryCmd = &cobra.Command{
	Use:   "appreg",
	Short: "App Registry service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	appregistryCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	appregistryCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
