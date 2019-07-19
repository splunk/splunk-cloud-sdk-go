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
	case "help":
		err = help("provisioner.txt")
	case "list-provision-jobs":
		result, err = cmd.listProvisionJobs(argv)
	case "list-tenants":
		result, err = cmd.listTenants(argv)
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
	flags.Var(&apps, "apps", "List of apps that the tenant is subscribed to")
	err := flags.Parse(argv)
	if err != nil {
		return nil, err
	}

	return cmd.provisionerService.CreateProvisionJob(provisioner.CreateProvisionJobBody{Tenant: &tenant, Apps: apps})
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
