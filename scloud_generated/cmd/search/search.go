package search

////go:generate scpgen gen-cmd --name search --package search --output search-gen.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return searchCmd
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search service",
}
