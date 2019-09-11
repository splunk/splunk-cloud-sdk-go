package ingest

//go:generate scloudgen gen-cmd --name ingest --package ingest --output ingest-gen.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return ingestCmd
}

var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "ingest service",
}
