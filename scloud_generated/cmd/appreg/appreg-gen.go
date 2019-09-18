// Package appreg -- generated by scloudgen
// !! DO NOT EDIT !! 
// 
package appreg

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/pkg/appreg"
)


// createApp -- Creates an app.
var createAppCmd = &cobra.Command{
	Use:   "create-app",
	Short: "Creates an app.",
	RunE:  impl.CreateApp,
}

// createSubscription -- Creates a subscription.
var createSubscriptionCmd = &cobra.Command{
	Use:   "create-subscription",
	Short: "Creates a subscription.",
	RunE:  impl.CreateSubscription,
}

// deleteApp -- Removes an app.
var deleteAppCmd = &cobra.Command{
	Use:   "delete-app",
	Short: "Removes an app.",
	RunE:  impl.DeleteApp,
}

// deleteSubscription -- Removes a subscription.
var deleteSubscriptionCmd = &cobra.Command{
	Use:   "delete-subscription",
	Short: "Removes a subscription.",
	RunE:  impl.DeleteSubscription,
}

// getApp -- Returns the metadata of an app.
var getAppCmd = &cobra.Command{
	Use:   "get-app",
	Short: "Returns the metadata of an app.",
	RunE:  impl.GetApp,
}

// getKeys -- Returns a list of the public keys used for verifying signed webhook requests.
var getKeysCmd = &cobra.Command{
	Use:   "get-keys",
	Short: "Returns a list of the public keys used for verifying signed webhook requests.",
	RunE:  impl.GetKeys,
}

// getSubscription -- Returns or validates a subscription.
var getSubscriptionCmd = &cobra.Command{
	Use:   "get-subscription",
	Short: "Returns or validates a subscription.",
	RunE:  impl.GetSubscription,
}

// listAppSubscriptions -- Returns the collection of subscriptions to an app.
var listAppSubscriptionsCmd = &cobra.Command{
	Use:   "list-app-subscriptions",
	Short: "Returns the collection of subscriptions to an app.",
	RunE:  impl.ListAppSubscriptions,
}

// listApps -- Returns a list of apps.
var listAppsCmd = &cobra.Command{
	Use:   "list-apps",
	Short: "Returns a list of apps.",
	RunE:  impl.ListApps,
}

// listSubscriptions -- Returns the tenant subscriptions.
var listSubscriptionsCmd = &cobra.Command{
	Use:   "list-subscriptions",
	Short: "Returns the tenant subscriptions.",
	RunE:  impl.ListSubscriptions,
}

// rotateSecret -- Rotates the client secret for an app.
var rotateSecretCmd = &cobra.Command{
	Use:   "rotate-secret",
	Short: "Rotates the client secret for an app.",
	RunE:  impl.RotateSecret,
}

// updateApp -- Updates an app.
var updateAppCmd = &cobra.Command{
	Use:   "update-app",
	Short: "Updates an app.",
	RunE:  impl.UpdateApp,
}


func init() {

    appregCmd.AddCommand(createAppCmd)
    appregCmd.AddCommand(createSubscriptionCmd)
    appregCmd.AddCommand(deleteAppCmd)
    appregCmd.AddCommand(deleteSubscriptionCmd)
    appregCmd.AddCommand(getAppCmd)
    appregCmd.AddCommand(getKeysCmd)
    appregCmd.AddCommand(getSubscriptionCmd)
    appregCmd.AddCommand(listAppSubscriptionsCmd)
    appregCmd.AddCommand(listAppsCmd)
    appregCmd.AddCommand(listSubscriptionsCmd)
    appregCmd.AddCommand(rotateSecretCmd)
    appregCmd.AddCommand(updateAppCmd)
    

	// subTest1Cmd.Flags().StringP("id", "i", "", "resource identifier")
	// subTest2Cmd.Flags().StringP("id", "i", "", "resource identifier")
}