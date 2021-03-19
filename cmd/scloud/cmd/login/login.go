package login

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/pkg/login"
	usageUtil "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/util"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return loginCmd
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to Splunk Cloud Services",
	RunE:  impl.Login,
}

func init() {

	loginCmd.Flags().StringP("uid", "u", "", "Your username")
	loginCmd.Flags().StringP("pwd", "p", "", "Your password")
	loginCmd.Flags().BoolP("verbose", "", false, "Whether to display your access token")

	// Auth flow flags
	loginCmd.Flags().BoolP("use-refresh-token", "", false, "Whether to use refresh token authentication flow")
	loginCmd.Flags().BoolP("use-pkce", "", false, "use PKCE authentication flow")
	loginCmd.Flags().BoolP("use-device", "", false, "use device authentication flow")

	loginCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	loginCmd.SetHelpTemplate(usageUtil.HelpTemplate)
}
