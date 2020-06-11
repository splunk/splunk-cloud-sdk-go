package login

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const refreshFlow = "refresh"

// TODO: Adding password handling
type Options struct {
	verbose  bool
	authKind string
}

// Login -- impl
func Login(cmd *cobra.Command, args []string) error {

	// Step 1: Setup
	err := auth.LoginSetUp()
	if err != nil {
		return fmt.Errorf(`error login setup: ` + err.Error())
	}

	// Step 2: Obtain LoginOptions
	loginOption, err := parseLoginOption(cmd)
	if err != nil {
		return fmt.Errorf(`error failed to parse command line: ` + err.Error())
	}

	// Step 3: Get Authentication Flow function
	authFlow, err := auth.GetFlow(loginOption.authKind)
	if err != nil {
		return fmt.Errorf(`error authentication flow invalid: ` + err.Error())
	}

	// Step 4: Login given authentication
	context, err := auth.Login(cmd, authFlow)
	if err != nil {
		util.Error(err.Error())

		return errors.New(err.Error() + ".  Try again using the --logtostderr flag to show details about the error.")
	}

	// Step 5: Print Context
	if loginOption.verbose {
		jsonx.Pprint(cmd, context)
	}

	return nil
}

func getAuthKindFromProfile() (string, error) {
	profile, err := auth.GetEnvironmentProfile()

	if err != nil {
		return "", errors.New("error failed to obtain environment")
	}

	kind, ok := profile["kind"]
	if !ok {
		return "", errors.New("missing kind")
	}
	return kind, nil
}

func parseLoginOption(cmd *cobra.Command) (*Options, error) {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return nil, errors.New(`error parsing "verbose": ` + err.Error())
	}

	isRefreshFlow, err := cmd.Flags().GetBool("use-refresh-token")
	if err != nil {
		return nil, errors.New(`error parsing "use-refresh-token": ` + err.Error())
	}

	var authKind string

	if isRefreshFlow {
		authKind = refreshFlow
	} else {
		authKind, err = getAuthKindFromProfile()

		if err != nil {
			return nil, errors.New(`error obtaining authentication kind: ` + err.Error())
		}
	}

	return &Options{
		verbose:  verbose,
		authKind: authKind,
	}, nil
}
