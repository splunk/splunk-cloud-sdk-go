package provisioner

////go:generate scloudgen gen-cmd --name provisioner --package provisioner --output provisioner-gen.go | gofmt

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return provisionerCmd
}

var provisionerCmd = &cobra.Command{
	Use:   "provisioner",
	Short: "Provisioner service",
}
