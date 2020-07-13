package context

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	"github.com/thoas/go-funk"
)

const ExpiresIn = 43200
const DefaultExpiresIn = 3600
const TokenType = "Bearer"
const Scope = "offline_access openid email profile"
const DefaultEnv = "prod"

var ValidKeys = map[string]interface{}{
	"access_token": "",
}

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

var set = &cobra.Command{
	Use:   "set",
	Short: "Set token context details",
	Run: func(cmd *cobra.Command, args []string) {

		// Extract flag values
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")

		// Validate Key
		var isValidKey = funk.Contains(ValidKeys, key)

		if !isValidKey {
			message := fmt.Sprintf("Here are the keys you can set:\n %s\n", ValidKeys)
			jsonx.Pprint(cmd, message)
			return
		}

		expirationToUse := DefaultExpiresIn

		if auth.GetEnvironmentName() != DefaultEnv {
			expirationToUse = ExpiresIn
		}

		clientID, err := auth.GetClientID(cmd)
		if err != nil {
			jsonx.Pprint(cmd, err)
			return
		}

		currentContext := auth.GetCurrentContext(clientID)
		var context map[string]interface{}

		if currentContext == nil {
			context = map[string]interface{}{
				"token_type": TokenType,
				"scope":      Scope,
			}
		} else {
			context = auth.ToMap(currentContext)
		}

		context[key] = value
		context["expires_in"] = expirationToUse

		auth.SetContext(cmd, context)
	},
}

func init() {
	contextCmd.AddCommand(list)
	contextCmd.AddCommand(set)

	set.Flags().StringP("key", "k", "", "The key stored in the context file")
	set.Flags().StringP("value", "p", "", "The value stored in the context file")

	_ = set.MarkFlagRequired("key")
	_ = set.MarkFlagRequired("value")
}
