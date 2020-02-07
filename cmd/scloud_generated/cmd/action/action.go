package action

////go:generate scloudgen gen-cmd --name action --package action --output action-gen.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return actionCmd
}

// catalogCmd represents the catalog command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Action service",
}
