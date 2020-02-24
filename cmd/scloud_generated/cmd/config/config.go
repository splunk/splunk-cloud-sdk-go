package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/splunk/splunk-cloud-sdk-go/util"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const LegacyCfgFileName = ".scloud"
const CfgFileName = ".scloud.toml"

var GlobalFlags = map[string]interface{}{
	"env":            "",
	"tenant":         "",
	"username":       "",
	"auth-url":       "",
	"host-url":       "",
	"ca-cert":        "",
	"insecure":       false,
	"testhookdryrun": false,
	"testhook":       false,
}

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return configCmd
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Save settings in a local configuration file",
}

var get = &cobra.Command{
	Use:   "get",
	Short: "Retrieve the value of a given setting (key)",
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		util.Info(viper.GetString(key))
	},
}

// Note: delete this, or read the file directly?
var list = &cobra.Command{
	Use:   "list",
	Short: "Retrieve all configuration settings",
	Run: func(cmd *cobra.Command, args []string) {
		for k, v := range viper.AllSettings() {
			util.Info("%s = %v\n", k, v)
		}
	},
}

var set = &cobra.Command{
	Use:   "set",
	Short: "Save a value for a given setting (key)",
	Run: func(cmd *cobra.Command, args []string) {
		// Note: need to check again because it could have been deleted
		home, err := homedir.Dir()
		if err != nil {
			util.Fatal(err.Error())
		}
		confFile := fmt.Sprintf("%s"+string(os.PathSeparator)+"%s", home, CfgFileName)
		if !FileExists(confFile) {
			Initialize()
		}
		// Note: this written this way so we don't override the global vars
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")

		// prevent non-supported keys from being written
		if isValidProperty(key) {
			viper.Set(key, value)
			viper.SetConfigType("toml")
			viper.SetConfigFile(confFile)
			err = viper.ReadInConfig()
			if err != nil {
				util.Error(err.Error())
			}
			err := viper.WriteConfig()
			if err != nil {
				util.Error(err.Error())
			}
		} else {
			util.Info("Here are the settings you can save:\n %s\n", GlobalFlags)
		}
	},
}

var reset = &cobra.Command{
	Use:   "reset",
	Short: "Delete the saved settings from the local configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			util.Error(err.Error())
		}
		confFile := fmt.Sprintf("%s"+string(os.PathSeparator)+"%s", home, CfgFileName)
		if FileExists(confFile) {
			util.Info("Deleting configuration file: %s.\n", CfgFileName)
			err = os.Remove(confFile)
			if err != nil {
				util.Error(err.Error())
			}
		}
	},
}

//   try to use .scloud (without) .toml
func Initialize() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		util.Error(err.Error())
	}
	confFile := fmt.Sprintf("%s"+string(os.PathSeparator)+"%s", home, CfgFileName)
	err = ioutil.WriteFile(confFile, []byte{}, 0755)
	if err != nil {
		util.Error("Unable to write a new configuration file: %v", err)
	}
	// Search config in home directory with name ".scloud" (without extension).
	Load(home, confFile)
}

func Load(home string, confFile string) {
	viper.AddConfigPath(home)
	viper.SetConfigName(CfgFileName)
	viper.SetConfigFile(confFile)
	// Read in the new config
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		util.Error(err.Error())
	}
}

func Migrate(source string, target string) {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		util.Error(err.Error())
		return
	}
	err = ioutil.WriteFile(target, input, 0644)
	if err != nil {
		util.Error("Error creating", target)
		util.Error(err.Error())
		return
	}
}

func isValidProperty(key string) bool {
	for prop := range GlobalFlags {
		if key == prop {
			return true
		}
	}
	return false
}

func FileExists(filename string) bool {
	fileinfo, e := os.Stat(filename)
	if e == nil {
		return !fileinfo.IsDir()
	}
	if os.IsNotExist(e) {
		return false
	}
	return true
}

func init() {
	configCmd.AddCommand(get)
	configCmd.AddCommand(list)
	configCmd.AddCommand(set)
	configCmd.AddCommand(reset)

	get.Flags().StringP("key", "k", "", "The setting name.")
	_ = get.MarkFlagRequired("key")

	set.Flags().StringP("key", "k", "", "The key stored in the settings.")
	set.Flags().StringP("value", "p", "", "The value stored in the settings.")

	_ = set.MarkFlagRequired("key")
	_ = set.MarkFlagRequired("value")
}
