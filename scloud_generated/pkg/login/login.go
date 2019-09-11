package login

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Login -- impl
func Login(cmd *cobra.Command, args []string) error {
	fmt.Printf("called Login\n")
	return nil
}
