package catalog

////go:generate scloudgen gen-cmd --name catalog --package catalog --output catalog-gen.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return catalogCmd
}

// catalogCmd represents the catalog command
var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "catalog service",
}
