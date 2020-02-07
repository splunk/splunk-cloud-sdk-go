package login

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/pkg/login"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return loginCmd
}

// loginCmd represents the catalog command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to Splunk Cloud Services",
	RunE:  impl.Login,
}

func init() {

	loginCmd.Flags().StringP("uid", "u", "", "Your username")
	loginCmd.Flags().StringP("pwd", "p", "", "Your password")
	loginCmd.Flags().BoolP("verbose", "", false, "Whether to display your access token")
}
