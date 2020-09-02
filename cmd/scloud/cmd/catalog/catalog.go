package catalog

////go:generate scloudgen gen-cmd --name catalog --package catalog --output catalog-gen.go

import (
	"github.com/spf13/cobra"
	usageUtil "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/util"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return catalogCmd
}

// catalogCmd represents the catalog command
var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Catalog service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	catalogCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	catalogCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
