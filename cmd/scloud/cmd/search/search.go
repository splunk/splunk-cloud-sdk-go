package search

////go:generate scpgen gen-cmd --name search --package search --output search-gen.go

import (
	"github.com/spf13/cobra"
	usageUtil "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/util"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return searchCmd
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	searchCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	searchCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
