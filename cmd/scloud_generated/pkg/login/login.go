package login

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/jsonx"
)

// Login -- impl
func Login(cmd *cobra.Command, args []string) error {

	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return fmt.Errorf(`error parsing "verbose": ` + err.Error())
	}

	context, err := auth.Login(cmd)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if verbose {
		jsonx.Pprint(cmd, context)
	}

	return nil
}
