package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "casper",
	Short: "Casper",
	Long: `Configuration Automation for Safe and Painless Environment Releases

     .-----.
   .' -   - '.
  /  .-. .-.  \
  |  | | | |  |
   \ \o/ \o/ /
  _/    ^    \_
 | \  '---'  / |
 / /'--. .--'\ \
/ /'---' '---'\ \
'.__.       .__.'
    '|     |'
     |     \
     \      '--.
      '.        '\
        ''---.   |
           ,__) /
            '..'
	`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.AddConfigPath(".")
	viper.AutomaticEnv() // read in environment variables that match

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error parsing config file:", err)
	}
}
