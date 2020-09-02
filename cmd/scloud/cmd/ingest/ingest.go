package ingest

////go:generate scloudgen gen-cmd --name ingest --package ingest --output ingest-gen.go

import (
	"github.com/spf13/cobra"
	usageUtil "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/util"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return ingestCmd
}

var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Ingest service",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func init() {
	ingestCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	ingestCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
