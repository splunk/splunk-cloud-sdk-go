package config

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return configCmd
}

// configCmd represents the catalog command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config service",
}
