package context

import (
	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
)

func Cmd() *cobra.Command {
	return contextCmd
}

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Token Context Information",
	Long:  "Context represents an authentication context, which is the result of a successful OAuth authentication",
}

var list = &cobra.Command{
	Use:   "list",
	Short: "Display the token context details, including the access token and expiration",
	Run: func(cmd *cobra.Command, args []string) {
		context := auth.GetContext(cmd)
		jsonx.Pprint(cmd, context)
	},
}

func init() {
	contextCmd.AddCommand(list)
}
