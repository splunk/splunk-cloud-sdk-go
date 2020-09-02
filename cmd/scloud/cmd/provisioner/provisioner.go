package provisioner

////go:generate scloudgen gen-cmd --name provisioner --package provisioner --output provisioner-gen.go | gofmt

import (
	"github.com/spf13/cobra"
	usageUtil "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/util"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return provisionerCmd
}

var provisionerCmd = &cobra.Command{
	Use:   "provisioner",
	Short: "Provisioner service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	provisionerCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	provisionerCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
