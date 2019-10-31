/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"flag"
	"fmt"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services/provisioner"
)

const (
	ProvisionerServiceVersion = "v1beta1"
)

type ProvisionerCommand struct {
	provisionerService *provisioner.Service
}

func newProvisionerCommand(client *sdk.Client) *ProvisionerCommand {
	return &ProvisionerCommand{
		provisionerService: client.ProvisionerService,
	}
}

func (cmd *ProvisionerCommand) Dispatch(argv []string) (result interface{}, err error) {
	arg, argv := head(argv)
	switch arg {
	case "":
		eusage("too few arguments")
	case "create-provision-job":
		result, err = cmd.createProvisionJob(argv)
	case "get-provision-job":
		result, err = cmd.getProvisionJob(argv)
	case "get-spec-json":
		result, err = cmd.getSpecJSON(argv)
	case "get-spec-yaml":
		result, err = cmd.getSpecYaml(argv)
	case "get-tenant":
		result, err = cmd.getTenant(argv)
	case "get-invite":
		result, err = cmd.getInvite(argv)
	case "help":
		result, err := getHelp("provisioner.txt")
		if err == nil {
			fmt.Println(result)
		}
	case "list-provision-jobs":
		result, err = cmd.listProvisionJobs(argv)
	case "list-tenants":
		result, err = cmd.listTenants(argv)
	case "create-invite":
		result, err = cmd.createInvite(argv)
	case "update-invite":
		result, err = cmd.updateInvite(argv)
	case "delete-invite":
		err = cmd.deleteInvite(argv)
	case "list-invites":
		result, err = cmd.listInvites(argv)
	default:
		fatal("unknown command: '%s'", arg)
	}
	return
}

func (cmd *ProvisionerCommand) createProvisionJob(argv []string) (interface{}, error) {
	var tenant string
	var apps multiFlags

	flags := flag.NewFlagSet("create-provision-job", flag.ExitOnError)
	flags.StringVar(&tenant, "tenant", "", "Tenant name")
	flags.Var(&apps, "app", "One or more apps that the tenant is subscribed to")
	err := flags.Parse(argv)
	if err != nil {
		return nil, err
	}

	return cmd.provisionerService.CreateProvisionJob(provisioner.CreateProvisionJobBody{Tenant: &tenant, Apps: apps})
}

func (cmd *ProvisionerCommand) createInvite(argv []string) (interface{}, error) {
	var email string
	var comment string
	var groups multiFlags

	flags := flag.NewFlagSet("create-invite", flag.ExitOnError)
	flags.StringVar(&email, "email", "", "Email address of the person receiving the invite")
	flags.StringVar(&comment, "comment", "", "Comment")
	flags.Var(&groups, "groups", "Tenant groups that the user will belong to")
	err := flags.Parse(argv)
	if err != nil {
		return nil, err
	}

	return cmd.provisionerService.CreateInvite(provisioner.InviteBody{Email: email, Comment: &comment, Groups: groups})
}

func (cmd *ProvisionerCommand) getProvisionJob(argv []string) (interface{}, error) {
	jobID := head1(argv)
	return cmd.provisionerService.GetProvisionJob(jobID)
}

func (cmd *ProvisionerCommand) getSpecJSON(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecJSON("api", ProvisionerServiceVersion, "provisioner", cmd.provisionerService.Client)
}

func (cmd *ProvisionerCommand) getSpecYaml(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecYaml("api", ProvisionerServiceVersion, "provisioner", cmd.provisionerService.Client)
}

func (cmd *ProvisionerCommand) getTenant(argv []string) (interface{}, error) {
	tenant := head1(argv)
	return cmd.provisionerService.GetTenant(tenant)
}

func (cmd *ProvisionerCommand) listProvisionJobs(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return cmd.provisionerService.ListProvisionJobs()
}

func (cmd *ProvisionerCommand) listTenants(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return cmd.provisionerService.ListTenants()
}

func (cmd *ProvisionerCommand) listInvites(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return cmd.provisionerService.ListInvites()
}

func (cmd *ProvisionerCommand) getInvite(argv []string) (interface{}, error) {
	inviteID := head1(argv)

	if inviteID == "" {
		etoofew()
	}
	return cmd.provisionerService.GetInvite(inviteID)
}

func (cmd *ProvisionerCommand) updateInvite(argv []string) (interface{}, error) {
	// Required args
	inviteID, args := head(argv)

	if inviteID == "" {
		etoofew()
	}

	// Optional flags
	flags := flag.NewFlagSet("update-invite", flag.ExitOnError)
	var updateAction string
	flags.StringVar(&updateAction, "update-action-type", "accept", "Valid values for update invite actions are accept, reject or resend. By default it's accept.")
	err := flags.Parse(args) //nolint:errcheck
	if err != nil {
		return nil, err
	}

	var updateActionType provisioner.UpdateInviteBodyAction

	switch updateAction {
	case "accept":
		updateActionType = provisioner.UpdateInviteBodyActionAccept
	case "reject":
		updateActionType = provisioner.UpdateInviteBodyActionReject
	case "resend":
		updateActionType = provisioner.UpdateInviteBodyActionResend
	default:
		updateActionType = provisioner.UpdateInviteBodyActionAccept
	}

	updateProvisionerInviteBody := provisioner.UpdateInviteBody{Action: updateActionType}

	return cmd.provisionerService.UpdateInvite(inviteID, updateProvisionerInviteBody)
}

func (cmd *ProvisionerCommand) deleteInvite(argv []string) error {
	inviteID := head1(argv)

	if inviteID == "" {
		etoofew()
	}
	return cmd.provisionerService.DeleteInvite(inviteID)
}
