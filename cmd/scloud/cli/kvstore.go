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
	"encoding/json"
	"flag"
	"strconv"
	"strings"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"

	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/splunk/splunk-cloud-sdk-go/services/kvstore"
)

const (
	KvStoreServiceVersion = "v1beta1"
)

type KVStoreCommand struct {
	kvstoreService *kvstore.Service
}

func newKVStoreCommand(client *sdk.Client) *KVStoreCommand {
	return &KVStoreCommand{
		kvstoreService: client.KVStoreService,
	}
}

func (cmd *KVStoreCommand) Dispatch(args []string) (result interface{}, err error) {
	arg, args := head(args)
	switch arg {
	case "":
		eusage("too few arguments")
	case "create-index":
		result, err = cmd.createIndex(args)
	//There is an existing issue assigned to Gateway for DeleteIndex being omitted in the current specs
	//Delete index should be supported back once this issue is resolved
	//Tracked by SCP-10206
	//case "delete-index":
	//	result, err = cmd.deleteIndex(args)
	case "delete-record":
		result, err = cmd.deleteRecord(args)
	case "delete-records":
		result, err = cmd.deleteRecords(args)
	case "get-record":
		result, err = cmd.getRecordByKey(args)
	case "get-health-status":
		result, err = cmd.getServiceHealthStatus(args)
	case "get-spec-json":
		result, err = cmd.getSpecJSON(args)
	case "get-spec-yaml":
		result, err = cmd.getSpecYaml(args)
	case "help":
		result, err := getHelp("kvstore.txt")
		if err == nil {
			fmt.Println(result)
		}
	case "insert-batch-records":
		result, err = cmd.insertBatchRecords(args)
	case "insert-record":
		result, err = cmd.insertRecord(args)
	case "list-indexes":
		result, err = cmd.listIndexes(args)
	case "list-records":
		result, err = cmd.listRecords(args)
	case "query":
		result, err = cmd.queryRecords(args)
	default:
		fatal("unknown command: '%s'", arg)
	}
	return
}

func (cmd *KVStoreCommand) createIndex(args []string) (interface{}, error) {
	collectionName, opts := head(args)
	var indexName string
	var indexDef = kvstore.IndexDefinition{}
	var fields multiFlags
	flags := flag.NewFlagSet("create-index", flag.ExitOnError)
	flags.StringVar(&indexName, "index", "", "index name")
	flags.Var(&fields, "field", "field definition field:direction")
	flags.Parse(opts) //nolint:errcheck

	fsArgs := flags.Args()
	if len(fsArgs) > 0 {
		fatal("unexpected argument(s): %s", fsArgs)
	}

	indexDef.Name = indexName
	for _, field := range fields {
		fieldDef := strings.Split(field, ":")
		direction, err := strconv.ParseInt(fieldDef[1], 10, 32)
		if err != nil {
			fatal(err.Error())
		}
		indexDef.Fields = append(indexDef.Fields, kvstore.IndexFieldDefinition{Field: fieldDef[0], Direction: int32(direction)})
	}
	return cmd.kvstoreService.CreateIndex(collectionName, indexDef)
}

//func (cmd *KVStoreCommand) deleteIndex(args []string) (interface{}, error) {
//	collectionName, indexName := head2(args)
//	return nil, cmd.kvstoreService.DeleteIndex(collectionName, indexName)
//}

func (cmd *KVStoreCommand) deleteRecord(args []string) (interface{}, error) {
	collectionName, key := head2(args)
	return nil, cmd.kvstoreService.DeleteRecordByKey(collectionName, key)
}

func (cmd *KVStoreCommand) deleteRecords(args []string) (interface{}, error) {
	collectionName, opts := head(args)
	var query string
	flags := flag.NewFlagSet("delete-records", flag.ExitOnError)
	flags.StringVar(&query, "query", "", "JSON string query")
	flags.Parse(opts) //nolint:errcheck

	fsArgs := flags.Args()
	if len(fsArgs) > 0 {
		fatal("unexpected argument(s): %s", fsArgs)
	}
	queryParams := kvstore.DeleteRecordsQueryParams{Query: query}
	return nil, cmd.kvstoreService.DeleteRecords(collectionName, &queryParams)
}

func (cmd *KVStoreCommand) getRecordByKey(args []string) (interface{}, error) {
	collectionName, key := head2(args)
	return cmd.kvstoreService.GetRecordByKey(collectionName, key)
}

func (cmd *KVStoreCommand) getServiceHealthStatus(args []string) (interface{}, error) {
	return cmd.kvstoreService.Ping()
}

func (cmd *KVStoreCommand) insertBatchRecords(args []string) (interface{}, error) {
	collectionName := head1(args)
	r := bufio.NewReader(os.Stdin)
	for {
		batch, err := readBatch(r)
		records := fmt.Sprintf("[%v]", strings.Join(batch, ","))
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		var data []map[string]interface{}
		err = json.Unmarshal([]byte(records), &data)
		if err != nil {
			fatal(err.Error())
		}
		if result, err := cmd.kvstoreService.InsertRecords(collectionName, data); err != nil {
			return result, err
		}
	}
	return nil, nil
}

func (cmd *KVStoreCommand) insertRecord(args []string) (interface{}, error) {
	collectionName, record := head2(args)
	var data map[string]interface{}
	err := json.Unmarshal([]byte(record), &data)
	if err != nil {
		fatal(err.Error())
	}
	return cmd.kvstoreService.InsertRecord(collectionName, data)
}

func (cmd *KVStoreCommand) listIndexes(args []string) (interface{}, error) {
	collectionName := head1(args)
	return cmd.kvstoreService.ListIndexes(collectionName)
}

func (cmd *KVStoreCommand) listRecords(args []string) (interface{}, error) {
	collectionName, opts := head(args)
	var count int64
	var offset int64
	var fields multiFlags
	var orderBy multiFlags
	var filter string

	flags := flag.NewFlagSet("list-records", flag.ExitOnError)
	flags.Var(&fields, "field", "field definition field")
	flags.Int64Var(&count, "count", 10, "maximum number of records to return")
	flags.Int64Var(&offset, "offset", 1, "number of records to skip from the start")
	flags.Var(&orderBy, "order-by", "keys to sort by")
	flags.StringVar(&filter, "filter", "", "A filter to apply to the records")
	flags.Parse(opts) //nolint:errcheck

	fsArgs := flags.Args()
	if len(fsArgs) > 0 {
		fatal("unexpected argument(s): %s", fsArgs)
	}

	var count32 = int32(count)
	var offset32 = int32(offset)
	listRecordsQueryParams := kvstore.ListRecordsQueryParams{Fields: fields, Count: &count32, Offset: &offset32, Orderby: orderBy}
	return cmd.kvstoreService.ListRecords(collectionName, &listRecordsQueryParams)
}

func (cmd *KVStoreCommand) queryRecords(args []string) (interface{}, error) {
	collectionName, opts := head(args)
	var count int64
	var offset int64
	var fields multiFlags
	var orderBy multiFlags
	var query string

	flags := flag.NewFlagSet("query-records", flag.ExitOnError)
	flags.Var(&fields, "field", "field definition field")
	flags.Int64Var(&count, "count", 10, "maximum number of records to return")
	flags.Int64Var(&offset, "offset", 1, "number of records to skip from the start")
	flags.Var(&orderBy, "order-by", "keys to sort by")
	flags.StringVar(&query, "json", "", "JSON string query")
	flags.Parse(opts) //nolint:errcheck

	fsArgs := flags.Args()
	if len(fsArgs) > 0 {
		fatal("unexpected argument(s): %s", fsArgs)
	}
	var count32 = int32(count)
	var offset32 = int32(offset)
	queryRecordsQueryParams := kvstore.QueryRecordsQueryParams{Query: query, Fields: fields, Count: &count32, Offset: &offset32, Orderby: orderBy}
	return cmd.kvstoreService.QueryRecords(collectionName, &queryRecordsQueryParams)
}

func (cmd *KVStoreCommand) getSpecJSON(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecJSON("api", KvStoreServiceVersion, "kvstore", cmd.kvstoreService.Client)
}

func (cmd *KVStoreCommand) getSpecYaml(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecYaml("api", KvStoreServiceVersion, "kvstore", cmd.kvstoreService.Client)
}
