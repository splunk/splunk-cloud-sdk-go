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

package integration

import (
	"fmt"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/v2/services/collect"
	testutils "github.com/splunk/splunk-cloud-sdk-go/v2/test/utils"
	"github.com/stretchr/testify/require"
)

func TestCRUDJob(t *testing.T) {
	client := getSdkClient(t)
	job := collect.Job{
		ConnectorID: "aws-cloudwatch-metrics",
		Name:        fmt.Sprintf("gointegCollectJob%d", testutils.TimeSec),
		Schedule:    "16 * * * *",
		Parameters:  map[string]interface{}{"namespaces": "AWS/EC2"},
		ScalePolicy: collect.ScalePolicy{Static: collect.StaticScale{Workers: 2}},
	}

	createdJob, err := client.CollectService.CreateJob(job)
	require.Nil(t, err)
	require.NotNil(t, createdJob.Data)
	defer client.CollectService.DeleteJob(*createdJob.Data.Id)

	//get job
	getJob, err := client.CollectService.GetJob(*createdJob.Data.Id)
	require.Nil(t, err)
	require.Equal(t, *createdJob.Data.Id, *getJob.Data.Id)

	//List jobs
	listedJob, err := client.CollectService.ListJobs(nil)
	require.Nil(t, err)
	require.NotZero(t, len(listedJob.Data))

	//List jobs with query
	query := collect.ListJobsQueryParams{}.SetConnectorId(createdJob.Data.ConnectorID)
	listedJob1, err := client.CollectService.ListJobs(&query)
	require.Nil(t, err)
	require.NotZero(t, len(listedJob1.Data))
	require.Equal(t, createdJob.Data.ConnectorID, listedJob1.Data[0].ConnectorID)

	//Delete job
	err = client.CollectService.DeleteJob(*createdJob.Data.Id)
	require.Nil(t, err)
}

func TestPatchJob(t *testing.T) {
	client := getSdkClient(t)
	job := collect.Job{
		ConnectorID: "aws-cloudwatch-metrics",
		Name:        fmt.Sprintf("gointegCollectPatchJob%d", testutils.TimeSec),
		Schedule:    "16 * * * *",
		Parameters:  map[string]interface{}{"namespaces": "AWS/EC2"},
		ScalePolicy: collect.ScalePolicy{Static: collect.StaticScale{Workers: 2}},
	}

	createdJob, err := client.CollectService.CreateJob(job)
	require.Nil(t, err)
	defer client.CollectService.DeleteJob(*createdJob.Data.Id)

	//Patch job
	new_name := createdJob.Data.Name + "_new"
	patchJob := collect.JobPatch{
		Name: &new_name,
	}

	newJob, err := client.CollectService.PatchJob(*createdJob.Data.Id, patchJob)
	require.Nil(t, err)
	require.Equal(t, new_name, newJob.Data.Name)
}

func TestPatchJobs(t *testing.T) {
	client := getSdkClient(t)

	job := collect.Job{
		ConnectorID: "aws-cloudwatch-metrics",
		Name:        fmt.Sprintf("gointegCollectPatchJobs%d", testutils.TimeSec),
		Schedule:    "16 * * * *",
		Parameters:  map[string]interface{}{"namespaces": "AWS/EC2"},
		ScalePolicy: collect.ScalePolicy{Static: collect.StaticScale{Workers: 2}},
	}

	createdJob, err := client.CollectService.CreateJob(job)
	require.Nil(t, err)
	defer client.CollectService.DeleteJob(*createdJob.Data.Id)

	scale := collect.ScalePolicy{Static: collect.StaticScale{Workers: 1}}
	jobsPatchs := collect.JobsPatch{
		ScalePolicy: &scale,
	}

	query := collect.PatchJobsQueryParams{}.SetConnectorId(createdJob.Data.ConnectorID)
	newJobs, err := client.CollectService.PatchJobs(jobsPatchs, &query)
	require.Nil(t, err)
	require.True(t, newJobs.Data[0].Updated)
	jobId := newJobs.Data[0].Id

	job1, err := client.CollectService.GetJob(jobId)
	require.Nil(t, err)
	require.NotNil(t, job1.Data.Name)
	static := job1.Data.ScalePolicy.Static
	require.NotNil(t, static)

	work := static.Workers
	require.NotNil(t, work)
	require.Equal(t, int32(1), work)
}
