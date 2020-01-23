package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const LegacyCfgFileName = ".scloud"
const CfgFileName = ".scloud.toml"

var GlobalFlags = []string{"env", "tenant", "username", "auth-url", "host-url", "insecure"}

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return configCmd
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Store global flags in your local settings configuration.",
}

var get = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a setting for the given flag.",
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		fmt.Println(viper.GetString(key))
	},
}

// Note: delete this, or read the file directly?
var list = &cobra.Command{
	Use:   "list",
	Short: "Retrieve all properties from the config.",
	Run: func(cmd *cobra.Command, args []string) {
		for k, v := range viper.AllSettings() {
			fmt.Printf("%s = %v\n", k, v)
		}
	},
}

var set = &cobra.Command{
	Use:   "set",
	Short: "Cache a value for the given flag in the config.",
	Run: func(cmd *cobra.Command, args []string) {
		// Note: need to check again because it could have been deleted
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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
				fmt.Println(err)
			}
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("These are the valid keys that can be set:\n %s\n", GlobalFlags)
		}
	},
}

var reset = &cobra.Command{
	Use:   "reset",
	Short: "Restore config file, fall back on default flag settings, or legacy $HOME/.scloud file",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
		}
		confFile := fmt.Sprintf("%s"+string(os.PathSeparator)+"%s", home, CfgFileName)
		if FileExists(confFile) {
			fmt.Printf("Deleting config file: %s.\n", CfgFileName)
			err = os.Remove(confFile)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

//   try to use .scloud (without) .toml
func Initialize() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
	}
	confFile := fmt.Sprintf("%s"+string(os.PathSeparator)+"%s", home, CfgFileName)
	err = ioutil.WriteFile(confFile, []byte{}, 0755)
	if err != nil {
		fmt.Printf("Unable to write a new config file: %v", err)
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
		fmt.Println(err)
	}
}

func Migrate(source string, target string) {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(target, input, 0644)
	if err != nil {
		fmt.Println("Error creating", target)
		fmt.Println(err)
		return
	}
}

func isValidProperty(key string) bool {
	for _, prop := range GlobalFlags {
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

	get.Flags().StringP("key", "k", "", "The key stored in the settings.")
	_ = get.MarkFlagRequired("key")

	set.Flags().StringP("key", "k", "", "The key stored in the settings.")
	set.Flags().StringP("value", "v", "", "The value stored in the settings.")
	_ = set.MarkFlagRequired("key")
	_ = set.MarkFlagRequired("value")
}
