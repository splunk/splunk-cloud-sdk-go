package login

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/auth"
)

// Login -- impl
func Login(cmd *cobra.Command, args []string) error {
	_, err := auth.Login(cmd)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
