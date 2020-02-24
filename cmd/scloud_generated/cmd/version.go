package cmd

import (
	"github.com/splunk/splunk-cloud-sdk-go/util"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/scloud/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of Splunk Cloud Services CLI",
	RunE:  execVersionCmd,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func execVersionCmd(cmd *cobra.Command, args []string) error {

	util.Info("scloud version %s-%s\n", version.Version, version.Commit)
	return nil
}
