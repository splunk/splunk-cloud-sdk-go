package appregistry

////go:generate scloudgen gen-cmd --name app-registry --package appreg --output appreg-gen.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return appregistryCmd
}

// catalogCmd represents the catalog command
var appregistryCmd = &cobra.Command{
	Use:   "appreg",
	Short: "App Registry service",
}
