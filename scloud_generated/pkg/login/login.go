package login

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/utils"
)

// Login -- impl
func Login(cmd *cobra.Command, args []string) error {
	fmt.Printf("called Login\n")

	utils.Login(nil)

	return nil
}
