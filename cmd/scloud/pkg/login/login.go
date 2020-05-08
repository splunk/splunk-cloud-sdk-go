package login

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

// Login -- impl
func Login(cmd *cobra.Command, args []string) error {

	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return fmt.Errorf(`error parsing "verbose": ` + err.Error())
	}

	context, err := auth.Login(cmd)

	if err != nil {
		util.Error(err.Error())

		return errors.New(err.Error() + ".  Try again using the --logtostderr flag to show details about the error.")
	}

	if verbose {
		jsonx.Pprint(cmd, context)
	}

	return nil
}
