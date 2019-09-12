package identity

////go:generate scloudgen gen-cmd --name identity --package identity --output identity-gen.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return identityCmd
}

var identityCmd = &cobra.Command{
	Use:   "identity",
	Short: "identity service",
}
