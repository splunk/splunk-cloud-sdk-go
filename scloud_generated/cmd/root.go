package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/action"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/appreg"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/catalog"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/collect"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/config"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/forwarders"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/identity"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/kvstore"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/login"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/ml"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/provisioner"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/search"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cmd/streams"
)

var (
	cfgFile string
	//tenant  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scloud",
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scloud.yaml)")
	rootCmd.PersistentFlags().String("tenant", "", "tenant identifier")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(action.Cmd())
	rootCmd.AddCommand(appreg.Cmd())
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
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".scloud" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".scloud")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
