package identity

////go:generate scloudgen gen-cmd --name identity --package identity --output identity-gen.go

import (
	"github.com/spf13/cobra"
	usageUtil "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/util"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return identityCmd
}

var identityCmd = &cobra.Command{
	Use:   "identity",
	Short: "Identity service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	identityCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	identityCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
