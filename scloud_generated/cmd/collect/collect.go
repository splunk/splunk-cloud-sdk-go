package collect

//go:generate scloudgen gen-cmd --name collect --package collect --output collect-gen.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return collectCmd
}

// collectCmd represents the catalog command
var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "collect service",
}
