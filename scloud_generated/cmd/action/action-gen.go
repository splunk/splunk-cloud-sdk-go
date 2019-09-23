// Package action -- generated by scloudgen
// !! DO NOT EDIT !!
//
package action

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/pkg/action"
)


// createAction --
var createActionCmd = &cobra.Command{
	Use:   "create-action",
	Short: "",
	RunE:  impl.CreateAction,
}

// deleteAction --
var deleteActionCmd = &cobra.Command{
	Use:   "delete-action",
	Short: "",
	RunE:  impl.DeleteAction,
}

// getAction --
var getActionCmd = &cobra.Command{
	Use:   "get-action",
	Short: "",
	RunE:  impl.GetAction,
}

// getActionStatus --
var getActionStatusCmd = &cobra.Command{
	Use:   "get-action-status",
	Short: "",
	RunE:  impl.GetActionStatus,
}

// getActionStatusDetails --
var getActionStatusDetailsCmd = &cobra.Command{
	Use:   "get-action-status-details",
	Short: "",
	RunE:  impl.GetActionStatusDetails,
}

// getPublicWebhookKeys --
var getPublicWebhookKeysCmd = &cobra.Command{
	Use:   "get-public-webhook-keys",
	Short: "",
	RunE:  impl.GetPublicWebhookKeys,
}

// listActions --
var listActionsCmd = &cobra.Command{
	Use:   "list-actions",
	Short: "",
	RunE:  impl.ListActions,
}

// triggerAction --
var triggerActionCmd = &cobra.Command{
	Use:   "trigger-action",
	Short: "",
	RunE:  impl.TriggerAction,
}

// updateAction --
var updateActionCmd = &cobra.Command{
	Use:   "update-action",
	Short: "",
	RunE:  impl.UpdateAction,
}


func init() {
	actionCmd.AddCommand(createActionCmd)


	actionCmd.AddCommand(deleteActionCmd)
	var deleteActionActionName string
	deleteActionCmd.Flags().StringVar(&deleteActionActionName, "action-name", "", "The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.")
	deleteActionCmd.MarkFlagRequired("action-name")


	actionCmd.AddCommand(getActionCmd)
	var getActionActionName string
	getActionCmd.Flags().StringVar(&getActionActionName, "action-name", "", "The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.")
	getActionCmd.MarkFlagRequired("action-name")


	actionCmd.AddCommand(getActionStatusCmd)
	var getActionStatusActionName string
	getActionStatusCmd.Flags().StringVar(&getActionStatusActionName, "action-name", "", "The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.")
	getActionStatusCmd.MarkFlagRequired("action-name")
	var getActionStatusStatusId string
	getActionStatusCmd.Flags().StringVar(&getActionStatusStatusId, "status-id", "", "The ID of the action status.")
	getActionStatusCmd.MarkFlagRequired("status-id")


	actionCmd.AddCommand(getActionStatusDetailsCmd)
	var getActionStatusDetailsActionName string
	getActionStatusDetailsCmd.Flags().StringVar(&getActionStatusDetailsActionName, "action-name", "", "The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.")
	getActionStatusDetailsCmd.MarkFlagRequired("action-name")
	var getActionStatusDetailsStatusId string
	getActionStatusDetailsCmd.Flags().StringVar(&getActionStatusDetailsStatusId, "status-id", "", "The ID of the action status.")
	getActionStatusDetailsCmd.MarkFlagRequired("status-id")


	actionCmd.AddCommand(getPublicWebhookKeysCmd)


	actionCmd.AddCommand(listActionsCmd)


	actionCmd.AddCommand(triggerActionCmd)
	var triggerActionActionName string
	triggerActionCmd.Flags().StringVar(&triggerActionActionName, "action-name", "", "The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.")
	triggerActionCmd.MarkFlagRequired("action-name")

	var triggerActionActionMetadata string
	triggerActionCmd.Flags().StringVar(&triggerActionActionMetadata, "action-metadata", "", "")
	var triggerActionCreatedAt string
	triggerActionCmd.Flags().StringVar(&triggerActionCreatedAt, "created-at", "", "string-ified ISO-8601 date/time with zone.")
	var triggerActionCreatedBy string
	triggerActionCmd.Flags().StringVar(&triggerActionCreatedBy, "created-by", "", "The principal that generated the trigger event.")
	var triggerActionId string
	triggerActionCmd.Flags().StringVar(&triggerActionId, "id", "", "A unique identifier for this trigger event. Generated from a hash of all recursively-sorted event field values.")
	var triggerActionKind string
	triggerActionCmd.Flags().StringVar(&triggerActionKind, "kind", "", "")
	var triggerActionPayload string
	triggerActionCmd.Flags().StringVar(&triggerActionPayload, "payload", "", "")
	var triggerActionTenant string
	triggerActionCmd.Flags().StringVar(&triggerActionTenant, "tenant", "", "The tenant within which the trigger event was generated.")
	var triggerActionTriggerCondition string
	triggerActionCmd.Flags().StringVar(&triggerActionTriggerCondition, "trigger-condition", "", "A description of the condition that caused the trigger event.")
	var triggerActionTriggerName string
	triggerActionCmd.Flags().StringVar(&triggerActionTriggerName, "trigger-name", "", "The name of the trigger for which this event was created.")
	var triggerActionTtlSeconds string
	triggerActionCmd.Flags().StringVar(&triggerActionTtlSeconds, "ttl-seconds", "", "A time to live (TTL), expressed as seconds after createdAt, after which the trigger event will no longer be acted upon.")


	actionCmd.AddCommand(updateActionCmd)
	var updateActionActionName string
	updateActionCmd.Flags().StringVar(&updateActionActionName, "action-name", "", "The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.")
	updateActionCmd.MarkFlagRequired("action-name")



}
