package cmd

import (
	"flag"
	"fmt"
	"os"

	usageUtil "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/util"
	"github.com/splunk/splunk-cloud-sdk-go/util"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/action"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/appregistry"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/catalog"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/collect"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/config"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/context"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/forwarders"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/identity"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/kvstore"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/login"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/ml"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/provisioner"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/search"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cmd/streams"
)

var confFile string
var legacyFile string

var (
	env            string
	tenant         string
	authURL        string
	hostURL        string
	insecure       bool
	cacert         string
	testhookDryrun bool
	testhook       bool
	timeout        uint
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scloud",
	Short: "Splunk Cloud Services CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(action.Cmd())
	rootCmd.AddCommand(appregistry.Cmd())
	rootCmd.AddCommand(catalog.Cmd())
	rootCmd.AddCommand(collect.Cmd())
	rootCmd.AddCommand(config.Cmd())
	rootCmd.AddCommand(context.Cmd())
	rootCmd.AddCommand(forwarders.Cmd())
	rootCmd.AddCommand(identity.Cmd())
	rootCmd.AddCommand(ingest.Cmd())
	rootCmd.AddCommand(kvstore.Cmd())
	rootCmd.AddCommand(login.Cmd())
	rootCmd.AddCommand(ml.Cmd())
	rootCmd.AddCommand(provisioner.Cmd())
	rootCmd.AddCommand(search.Cmd())
	rootCmd.AddCommand(streams.Cmd())

	rootCmd.SetUsageTemplate(usageUtil.UsageTemplate)
	rootCmd.SetHelpTemplate(usageUtil.HelpTemplate)

	rootCmd.PersistentFlags().StringVar(&env, "env", "", "Set the target environment")
	rootCmd.PersistentFlags().StringVar(&tenant, "tenant", "", "Set the tenant to use for operations against platform services")
	rootCmd.PersistentFlags().StringVar(&authURL, "auth-url", "", "Set an auth URL to override the public SDC auth URL (https://<host>:<port>)")
	rootCmd.PersistentFlags().StringVar(&hostURL, "host-url", "", "Set a host URL to override the public SDC host (https://<host>:<port>)")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", false, "Specify whether to skip TLS validation. The default is \"false\" to enable TLS certificate validation")
	rootCmd.PersistentFlags().StringVar(&cacert, "ca-cert", "", "Set the public cert file to use for a local host using HTTPS with TLScertificate validation enabled")
	rootCmd.PersistentFlags().UintVar(&timeout, "timeout", 0, "Set HTTPS timeout (in seconds)")
	// add hidden test flags
	rootCmd.PersistentFlags().BoolVar(&testhookDryrun, "testhook-dryrun", false, "a string flag")
	err := rootCmd.PersistentFlags().MarkHidden("testhook-dryrun")
	if err != nil {
		fmt.Println(err)
	}

	rootCmd.PersistentFlags().BoolVar(&testhook, "testhook", false, "a string flag")
	err = rootCmd.PersistentFlags().MarkHidden("testhook")
	if err != nil {
		fmt.Println(err)
	}

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	err = flag.CommandLine.Parse([]string{})
	if err != nil {
		fmt.Println(err)
	}
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		util.Fatal(err.Error())
	}

	legacyFile = fmt.Sprintf("%s/%s", home, config.LegacyCfgFileName)
	confFile = fmt.Sprintf("%s/%s", home, config.CfgFileName)

	// Copy .scloud -> .scloud.toml (if it does not exist)
	if config.FileExists(legacyFile) && !config.FileExists(confFile) {
		config.Migrate(legacyFile, confFile)
	}

	// Create a new empty .scloud.toml
	if !config.FileExists(confFile) {
		config.Initialize()

	} else {
		// Use an existing .scloud.toml
		config.Load(home, confFile)
	}

	if testhookDryrun {
		config.GlobalFlags["testhookdryrun"] = true
		fmt.Println("enable testhook-dryrun")
	}

	if testhook {
		config.GlobalFlags["testhook"] = true
		fmt.Println("enable testhook")
	}

	// If flag 'log_dir' is provided, create log directory at the given location
	logDir, _ := rootCmd.Flags().GetString("log_dir")
	if logDir != "" {
		util.CreateLogDirectory(logDir)
	}

	config.GlobalFlags["env"] = env

	config.GlobalFlags["tenant"] = tenant

	config.GlobalFlags["host-url"] = hostURL

	config.GlobalFlags["auth-url"] = authURL

	config.GlobalFlags["insecure"] = insecure

	config.GlobalFlags["ca-cert"] = cacert

	config.GlobalFlags["timeout"] = timeout

}
