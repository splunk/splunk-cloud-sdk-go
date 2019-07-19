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

	"time"

	"strings"

	"github.com/golang/glog"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/argx"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
)

type SearchCommand struct {
	searchService *search.Service
}

const (
	SearchServiceVersion = "v2beta1"
)

func newSearchCommand(client *sdk.Client) *SearchCommand {
	return &SearchCommand{
		searchService: client.SearchService,
	}
}

func (cmd *SearchCommand) Dispatch(argv []string) (result interface{}, err error) {
	arg, argv := head(argv)
	switch arg {
	case "":
		eusage("too few arguments")
	case "cancel":
		result, err = cmd.cancel(argv)
	case "finalize":
		result, err = cmd.finalize(argv)
	case "list-results":
		result, err = cmd.listResults(argv)
	case "get-spec-json":
		result, err = cmd.getSpecJSON(argv)
	case "get-spec-yaml":
		result, err = cmd.getSpecYaml(argv)
	case "help":
		err = help("search.txt")
	case "list-jobs":
		result, err = cmd.listJobs(argv)
	case "wait":
		result, err = cmd.wait(argv)
	default:
		// otherwise, interpret the command as a literal query
		argv = push(arg, argv) // unget
		result, err = cmd.createJob(argv)
	}
	return
}

func (cmd *SearchCommand) cancel(argv []string) (interface{}, error) {
	jobID := head1(argv)
	return cmd.searchService.UpdateJob(jobID, search.UpdateJob{Status: search.UpdateJobStatusCanceled})
}

func (cmd *SearchCommand) finalize(argv []string) (interface{}, error) {
	jobid := head1(argv)
	return cmd.searchService.UpdateJob(jobid, search.UpdateJob{Status: search.UpdateJobStatusFinalized})
}

func (cmd *SearchCommand) listResults(argv []string) (interface{}, error) {
	var count float64
	var offset float64
	var fields multiFlags
	jobid, argv := head(argv)
	flags := flag.NewFlagSet("list-results", flag.ExitOnError)
	flags.Float64("count", count, "maximum number of records to return")
	flags.Float64("offset", offset, "number of records to skip from the start")
	flags.Var(&fields, "fields", "field to return for the result set, you can specify multiple fields of comma-separated values if multiple fields are required")
	flags.Parse(argv) //nolint:errcheck
	var count32 = float32(count)
	var offset32 = float32(offset)
	return cmd.searchService.ListResults(jobid, &search.ListResultsQueryParams{Count: &count32, Offset: &offset32, Field: strings.Join(fields, ",")})
}

func (cmd *SearchCommand) listJobs(argv []string) (interface{}, error) {
	var count float64
	var status string
	var searchStatus search.SearchStatus

	flags := flag.NewFlagSet("list-jobs", flag.ExitOnError)
	flags.Float64("count", count, "maximum number of records to return")
	flags.StringVar(&status, "status", "done", "current status of the search job, valid status values are 'running', 'done', 'canceled', and 'failed'")
	err := flags.Parse(argv)
	if err != nil {
		return nil, err
	}

	if status == "running" {
		searchStatus = search.SearchStatusRunning
	}
	if status == "done" {
		searchStatus = search.SearchStatusDone
	}
	if status == "cancel" {
		searchStatus = search.SearchStatusCanceled
	}
	if status == "failed" {
		searchStatus = search.SearchStatusCanceled
	}

	var count32 = float32(count)
	return cmd.searchService.ListJobs(&search.ListJobsQueryParams{Count: &count32, Status: &searchStatus})
}

// todo: add -poll-interval
func (cmd *SearchCommand) wait(argv []string) (interface{}, error) {
	jobid := head1(argv)
	return waitJob(cmd, &jobid)
}

// todo: "exponential" backoff on sleep duration
func waitJob(sc *SearchCommand, jobid *string) (interface{}, error) {
	for {
		job, err := sc.searchService.GetJob(*jobid)
		if err != nil {
			return nil, err
		}
		// return if we have reached a terminal state
		switch *job.Status {
		case search.SearchStatusDone, search.SearchStatusFailed:
			return job.Status, nil
		}
		glog.Infof("Wait jobid=%s state=%s", *jobid, *job.Status)
		time.Sleep(1 * time.Second)
	}
}

type jobArgs struct {
	Query        string `arg:"0"`
	EarliestTime string `arg:"earliest"`
	LatestTime   string `arg:"latest"`
	Wait         bool   `arg:"wait"`
	Sync         bool   `arg:"sync"`
}

func (cmd *SearchCommand) createJob(argv []string) (interface{}, error) {
	var (
		count  float32
		offset float32
	)
	args := jobArgs{Sync: true}
	argv, err := argx.Parse(argv, &args)
	if err != nil {
		return nil, err
	}
	checkEmpty(argv)
	params := search.SearchJob{
		Query: args.Query,
		QueryParameters: &search.QueryParameters{
			Earliest: &args.EarliestTime,
			Latest:   &args.LatestTime}}
	result, err := cmd.searchService.CreateJob(params)
	if err != nil {
		return nil, err
	}
	if args.Sync {
		args.Wait = true
	}
	if args.Wait {
		_, err = waitJob(cmd, result.Sid)
	}
	if args.Sync {
		return cmd.searchService.ListResults(*result.Sid, &search.ListResultsQueryParams{Count: &count, Offset: &offset, Field: ""})
	}
	return result, err
}

func (cmd *SearchCommand) getSpecJSON(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecJSON("api", SearchServiceVersion, "search", cmd.searchService.Client)
}

func (cmd *SearchCommand) getSpecYaml(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecYaml("api", SearchServiceVersion, "search", cmd.searchService.Client)
}
