package kvstore

////go:generate scloudgen gen-cmd --name action --package action --output action-gen.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return kvstoreCmd
}

// catalogCmd represents the catalog command
var kvstoreCmd = &cobra.Command{
	Use:   "kvstore",
	Short: "KV Store service",
}
