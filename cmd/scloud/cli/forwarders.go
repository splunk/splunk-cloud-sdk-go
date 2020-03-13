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

	"io/ioutil"

	"github.com/splunk/splunk-cloud-sdk-go/services/forwarders"
)

const (
	ForwardersServiceVersion = "v2beta1"
)

type ForwardersCommand struct {
	forwardersService *forwarders.Service
}

func newForwardersCommand(client *sdk.Client) *ForwardersCommand {
	return &ForwardersCommand{
		forwardersService: client.ForwardersService,
	}
}

func (cmd *ForwardersCommand) parse(args []string) []string {
	flags := flag.NewFlagSet("forwarders command", flag.ExitOnError)
	flags.Parse(args) //nolint:errcheck
	return flags.Args()
}

func (cmd *ForwardersCommand) Dispatch(args []string) (result interface{}, err error) {
	arg, args := head(args)
	args = cmd.parse(args)
	switch arg {
	case "":
		eusage("too few arguments")
	case "create-certificate":
		result, err = cmd.createCertificate(args)
	case "delete-certificate":
		err = cmd.deleteCertificate(args)
	case "delete-certificates":
		err = cmd.deleteCertificates(args)
	case "list-certificates":
		result, err = cmd.getCertificates(args)
	case "get-spec-json":
		result, err = cmd.getSpecJSON(args)
	case "get-spec-yaml":
		result, err = cmd.getSpecYaml(args)
	case "help":
		result, err := getHelp("forwarders.txt")
		if err == nil {
			fmt.Println(result)
		}
	default:
		fatal("unknown command: '%s'", arg)
	}
	return
}

func (cmd *ForwardersCommand) createCertificate(args []string) (interface{}, error) {
	pemFile := head1(args)
	fileBytes, err := ioutil.ReadFile(pemFile)
	if err != nil {
		return nil, err
	}
	str := string(fileBytes)

	return cmd.forwardersService.AddCertificate(forwarders.Certificate{Pem: str})
}

func (cmd *ForwardersCommand) deleteCertificate(args []string) error {
	slot := head1(args)
	return cmd.forwardersService.DeleteCertificate(slot)
}

func (cmd *ForwardersCommand) deleteCertificates(args []string) error {
	checkEmpty(args)
	return cmd.forwardersService.DeleteCertificates()
}

func (cmd *ForwardersCommand) getCertificates(args []string) (interface{}, error) {
	return cmd.forwardersService.ListCertificates()
}

func (cmd *ForwardersCommand) getSpecJSON(args []string) (interface{}, error) {
	checkEmpty(args)
	return GetSpecJSON("api", ForwardersServiceVersion, "forwarders", cmd.forwardersService.Client)
}

func (cmd *ForwardersCommand) getSpecYaml(args []string) (interface{}, error) {
	checkEmpty(args)
	return GetSpecYaml("api", ForwardersServiceVersion, "forwarders", cmd.forwardersService.Client)
}
