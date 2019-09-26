package login

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/auth"
)

// Login -- impl
func Login(cmd *cobra.Command, args []string) error {
	fmt.Printf("called Login\n")

	_, err:=auth.Login(nil)

	if err!=nil{
		fmt.Println(err)
		return err
	}

	auth.GetClient()

	return nil
}
