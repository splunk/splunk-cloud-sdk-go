package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version info",
	RunE:  execVersionCmd,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func execVersionCmd(cmd *cobra.Command, args []string) error {
	versionString := semVersion
	if versionString == "" {
		versionString = "0.0.0"
	}
	commitString := commitHash
	if commitHash == "" {
		commitString = "develop"
	}

	fmt.Fprintf(os.Stdout, "%s - v%s#%s.%s.%s.%s\n",
		appName,
		versionString,
		commitString,
		runtime.GOOS,
		runtime.GOARCH,
		buildDate,
	)
	return nil
}
