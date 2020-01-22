package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/action"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/appregistry"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/catalog"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/collect"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/config"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/forwarders"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/identity"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/kvstore"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/login"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/ml"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/provisioner"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/search"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd/streams"
)

var confFile string
var legacyFile string

var (
	env            string
	tenant         string
	userName       string
	authURL        string
	hostURL        string
	insecure       bool
	cacert         string
	testhookDryrun bool
	testhook       bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scloud_gen",
	Short: "Splunk Cloud Platform CLI",
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
	rootCmd.AddCommand(forwarders.Cmd())
	rootCmd.AddCommand(identity.Cmd())
	rootCmd.AddCommand(ingest.Cmd())
	rootCmd.AddCommand(kvstore.Cmd())
	rootCmd.AddCommand(login.Cmd())
	rootCmd.AddCommand(ml.Cmd())
	rootCmd.AddCommand(provisioner.Cmd())
	rootCmd.AddCommand(search.Cmd())
	rootCmd.AddCommand(streams.Cmd())

	rootCmd.PersistentFlags().StringVar(&env, "env", "", "target environment")
	rootCmd.PersistentFlags().StringVar(&tenant, "tenant", "", "tenant identifier")
	rootCmd.PersistentFlags().StringVar(&userName, "username", "", "email address")
	rootCmd.PersistentFlags().StringVar(&authURL, "auth-url", "", "scheme://host<:port>")
	rootCmd.PersistentFlags().StringVar(&hostURL, "host-url", "", "scheme://host<:port>")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", false, "disable TLS cert validation")
	rootCmd.PersistentFlags().StringVar(&cacert, "cacert", "", "cacert file")

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
		fmt.Println(err)
		os.Exit(1)
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
}
