package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/scloud/version"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/jsonx"
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

	err := fmt.Sprintf("scloud version %s-%s\n", version.Version, version.Commit)
	jsonx.Pprint(cmd, err)
	return nil
}
