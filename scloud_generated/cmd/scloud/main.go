package main

import (
	"flag"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd"
)

func main() {
	flag.Parse()

	cmd.Execute()
}
