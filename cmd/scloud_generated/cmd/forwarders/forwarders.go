package forwarders

////go:generate scloudgen gen-cmd --name forwarders --package forwarders --output forwarders-gen.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return forwardersCmd
}

var forwardersCmd = &cobra.Command{
	Use:   "forwarders",
	Short: "Forwarders service",
}
