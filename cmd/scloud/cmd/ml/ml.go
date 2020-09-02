package ml

////go:generate scloudgen gen-cmd --name ml --package ml --output ml-gen.go | gofmt

import (
	"github.com/spf13/cobra"
	usageUtil "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/util"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return mlCmd
}

var mlCmd = &cobra.Command{
	Use:   "ml",
	Short: "Machine Learning service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	mlCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	mlCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
